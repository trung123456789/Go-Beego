package implement

import (
	"apiproject/database/connection"
	"apiproject/models"
)

func GetAllUsers() []models.User {
	var userList []models.User
	database, errConnectDb := connection.CreateConnection()
	if errConnectDb != nil {
		panic(errConnectDb)
	}
	defer database.Close()
	querySelect := `
		SELECT user_id, name, age
		FROM user_infos
	`
	rows, errSelect := database.Query(querySelect)
	if errSelect != nil {
		panic(errSelect)
	}
	defer rows.Close()
	for rows.Next() {
		user := models.User{}
		if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
			panic(err)
		}

		userList = append(userList, user)
	}

	return userList
}