package implement

import (
	"apiproject/conf"
	"apiproject/database/connection"
	"errors"
	"fmt"
	"log"
)

func DeleteUser(uid string) (string, error) {
	var count int
	database, errConnectDb := connection.CreateConnection()
	if errConnectDb != nil {
		panic(errConnectDb)
	}
	defer database.Close()

	// Check user exist
	querySelect := `
		SELECT COUNT(*)
		FROM user_infos
		WHERE user_id = $1
	`
	log.Println(querySelect)
	errSelect := database.QueryRow(querySelect, uid).Scan(&count)
	if errSelect != nil {
		panic(errSelect)
	}
	if count == 0 {
		return "", errors.New(conf.IniConf.String("userNotExist"))
	}
	// Update user
	queryUpdate := `
		DELETE
		FROM user_infos 
		WHERE user_id = $1 returning user_id`
	log.Println(queryUpdate)
	_, errQuery := database.Exec(queryUpdate, uid)

	if errQuery != nil {
		return "", errQuery
	}
	uid = fmt.Sprintf(conf.IniConf.String("delSuccess"), uid)
	return uid, nil
}
