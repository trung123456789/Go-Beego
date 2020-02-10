package implement

import (
	"apiproject/conf"
	"apiproject/database/connection"
	"apiproject/models"
	"errors"
	"fmt"
	"log"
)

func UpdateUser(uid string, uu *models.User) (string, error) {
	var count int
	database, errConnectDb := connection.CreateConnection()
	if errConnectDb != nil {
		panic(errConnectDb)
	}
	defer  database.Close()
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
		UPDATE user_infos
		SET 
			name = $1,
			age = $2
		WHERE user_id = $3`
	log.Println(queryUpdate)
	log.Println(uu)
	_, errQuery := database.Exec(queryUpdate, uu.Name, uu.Age, uid)
	if errQuery != nil {
		panic(errQuery)
	}
	uid = fmt.Sprintf(conf.IniConf.String("updateSuccess"), uid)
	return uid, nil
}