package implement

import (
	"apiproject/conf"
	"apiproject/models"
	"errors"
	"github.com/astaxie/beego/orm"
)

func AddUser(u models.UserInfo) (string, error) {
	//database, errConnectDb := connection.CreateConnection()
	//if errConnectDb != nil {
	//	return "", errConnectDb
	//}
	//defer  database.Close()
	if u.UserId == "" || u.Password == "" {
		return "", errors.New(conf.IniConf.String("missingIdOrPass"))
	}
	o := orm.NewOrm()
	o.Using("default")

	user := models.UserInfo{
		UserId:    u.UserId,
		Name:      u.Name,
		Age:       u.Age,
		Password:  u.Password,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}

	_, err := o.Insert(&user)
	return u.UserId, err

	//if u.Id == "" {
	//	u.Id = "user_" + strconv.FormatInt(time.Now().UnixNano(), 10)
	//}
	//query := `
	//	INSERT INTO
	//	user_infos(user_id, name, age)
	//	VALUES($1, $2, $3)`
	//log.Println(query)
	//_, errQuery := database.Exec(query, u.Id, u.Name, u.Age)
	//if errQuery != nil {
	//	return "", errQuery
	//}
	//return u.Id, nil
}
