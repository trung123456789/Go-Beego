package implement

import (
	"apiproject/database/connection"
	"log"
)

func Login(id, password string) (int, error) {
	var count int
	database, errConnectDb := connection.CreateConnection()
	if errConnectDb != nil {
		return 0, errConnectDb
	}
	defer  database.Close()

	querySelect := `
		SELECT COUNT(*)
		FROM user_infos
		WHERE user_id = $1 AND password = $2
	`
	errSelect := database.QueryRow(querySelect, id, password).Scan(&count)
	log.Println("Err:", errSelect)
	log.Println(querySelect)
	if errSelect != nil {
		return 0, errSelect
	}
	return count, nil
}