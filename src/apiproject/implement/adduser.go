package implement

import (
	"apiproject/database/connection"
	"apiproject/models"
	"log"
	"strconv"
	"time"
)

func AddUser(u models.User) (string, error) {
	database, errConnectDb := connection.CreateConnection()
	if errConnectDb != nil {
		return "", errConnectDb
	}
	defer  database.Close()
	if u.Id == "" {
		u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	}
	query := `
		INSERT INTO
		user_infos(user_id, name, age)
		VALUES($1, $2, $3)`
	log.Println(query)
	_, errQuery := database.Exec(query, u.Id, u.Name, u.Age)
	if errQuery != nil {
		return "", errQuery
	}
	return u.Id, nil
}
