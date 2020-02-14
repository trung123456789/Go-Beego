package implement

import (
	"apiproject/models"
	"github.com/astaxie/beego/orm"
)

func GetAllUsers() (models.UserResponse, error) {
	var userResponse models.UserResponse
	var userList []models.UserInfo
	//database, errConnectDb := connection.CreateConnection()
	//if errConnectDb != nil {
	//	panic(errConnectDb)
	//}
	//defer database.Close()
	//querySelect := `
	//	SELECT user_id, name, age
	//	FROM user_infos
	//`
	//rows, errSelect := database.Query(querySelect)
	//if errSelect != nil {
	//	panic(errSelect)
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	user := models.User{}
	//	if err := rows.Scan(&user.Id, &user.Name, &user.Age); err != nil {
	//		panic(err)
	//	}
	//
	//	userList = append(userList, user)
	//}

	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("userinfo")
	record, err := qs.All(&userList)

	userResponse.RecordNum = record
	userResponse.UserList = userList
	return userResponse, err
}