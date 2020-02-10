package implement

import (
	"apiproject/conf"
	"apiproject/database/connection"
	"apiproject/models"
	"bytes"
	"database/sql"
	"errors"
	"fmt"
	"log"
)

func GetUser(u models.User) ([]models.User, error) {
	var userList []models.User
	var query bytes.Buffer
	database, errConnectDb := connection.CreateConnection()
	if errConnectDb != nil {
		panic(errConnectDb)
	}
	defer  database.Close()
	count := countUser(query, u, database)
	if count == 0 {
		return nil, errors.New(conf.IniConf.String("userNotExist"))
	}
	rows := selectUser(query, u, database)
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			panic(err)
		}

		userList = append(userList, user)
	}

	return userList, nil
}

func countUser(query bytes.Buffer, userRequest models.User, database *sql.DB) int {
	var count int
	queryCount := `
		SELECT COUNT(*)
		FROM user_infos
	`
	query.WriteString(queryCount)
	query = queryUtil(query, userRequest)
	log.Println("query:", query.String())
	err := database.QueryRow(query.String()).Scan(&count)
	if err != nil {
		log.Fatal(err)
	}
	return count
}

func selectUser(query bytes.Buffer, userRequest models.User, database *sql.DB) *sql.Rows {
	querySelect := `
		SELECT user_id, name, age
		FROM user_infos
	`
	query.WriteString(querySelect)
	query = queryUtil(query, userRequest)
	log.Println("query:", query.String())
	rows, errSelect := database.Query(query.String())
	if errSelect != nil {
		panic(errSelect)
	}
	return rows
}

func queryUtil(query bytes.Buffer, request models.User) bytes.Buffer {
	query.WriteString("WHERE 1 = 1")
	var queryTemp string
	if request.Id != "" {
		queryTemp := fmt.Sprint(" AND user_id LIKE ", "'%", request.Id, "%'")
		query.WriteString(queryTemp)
	}
	if request.Name != "" {
		queryTemp := fmt.Sprint(" AND name LIKE ", "'%", request.Name, "%'")
		query.WriteString(queryTemp)
	}
	if request.Age != 0 {
		queryTemp = fmt.Sprintf(" AND age = %d", request.Age)
		query.WriteString(queryTemp)
	}
	return query
}