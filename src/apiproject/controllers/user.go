package controllers

import (
	"apiproject/implement"
	"apiproject/logconf"
	"apiproject/models"
	"encoding/json"
	"github.com/astaxie/beego"
	"log"
	"net/http"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

var (
	logs = logconf.NewLogger("UserController")
)

// @Title CreateUser
// @Description create users
// @param	password	body 	string	true		"body for user content"
// @Success 201 {int} models.UserInfo.Id
// @Failure 403 body is empty
// @router /add [post]
func (u *UserController) Post() {
	logs.Start("ADD USER")
	var user models.UserInfo
	if u.CheckAuth() {
		defer u.deferFunc("ADD USER")
		errJson := json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		if errJson != nil {
			u.Ctx.Output.SetStatus(http.StatusBadRequest)
			u.Data["json"] = models.Msg{
				StatusCd: 0,
				Message:  errJson.Error(),
			}
			panic(0)
		}

		err := implement.Validate.Struct(user)
		if err != nil {
			u.Ctx.Output.SetStatus(http.StatusBadRequest)
			u.Data["json"] = models.Msg{
				StatusCd: 1,
				Message:  err.Error(),
			}
			panic(0)
		}

		uid, err := implement.AddUser(user)
		if err != nil {
			u.Ctx.Output.SetStatus(http.StatusBadRequest)
			u.Data["json"] = models.Msg{
				StatusCd: 0,
				Message:  err.Error(),
			}
			panic(0)
		}
		u.Data["json"] = map[string]string{"uid": uid}
		u.ServeJSON()
		panic(0)
	}
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /get/All [get]
func (u *UserController) GetAll() {
	logs.Start("GET ALL")
	if u.CheckAuth() {
		users, err := implement.GetAllUsers()
		if err != nil {
			u.Ctx.Output.SetStatus(http.StatusInternalServerError)
			u.Data["json"] = models.Msg{
				StatusCd: 1,
				Message:  err.Error(),
			}
			u.ServeJSON()
			return
		}
		u.Data["json"] = users
		u.ServeJSON()
		logs.End("GET ALL")
		return
	}
}

// @Title Get Some User
// @Description get user by user info
// @Param   id     query   int true       "task id"
// @Success 200 {object} models.User
// @router /get/One [post]
func (u *UserController) PostOneUser() {
	var user models.UserInfo
	if u.CheckAuth() {
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		log.Println(user)
		userList, err := implement.GetUser(user)
		if err != nil {
			u.Ctx.Output.SetStatus(http.StatusInternalServerError)
			u.Data["json"] = models.Msg{
				StatusCd: 1,
				Message:  err.Error(),
			}
		} else {
			u.Data["json"] = userList
		}
		u.ServeJSON()
	}
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /update/:uid [put]
func (u *UserController) Put() {
	if u.CheckAuth() {
		uid := u.GetString(":uid")
		if uid != "" {
			var user models.UserInfo
			_ = json.Unmarshal(u.Ctx.Input.RequestBody, &user)
			uu, err := implement.UpdateUser(uid, user)
			if err != nil {
				u.Ctx.Output.SetStatus(http.StatusInternalServerError)
				u.Data["json"] = models.Msg{
					StatusCd: 1,
					Message:  err.Error(),
				}
			} else {
				u.Data["json"] = uu
			}
		}
		u.ServeJSON()
	}
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /delete/:uid [delete]
func (u *UserController) Delete() {
	if u.CheckAuth() {
		uid := u.GetString(":uid")
		data, err := implement.DeleteUser(uid)
		if err != nil {
			u.Ctx.Output.SetStatus(http.StatusInternalServerError)
			u.Data["json"] = models.Msg{
				StatusCd: 1,
				Message:  err.Error(),
			}
		} else {
			u.Data["json"] = data
		}
		u.ServeJSON()
	}
}
