package controllers

import (
	"apiproject/implement"
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

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /add [post]
func (u *UserController) Post() {
	var user models.User
	if u.CheckAuth() {
		_ = json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		uid, err := implement.AddUser(user)
		if err != nil {
			u.Ctx.Output.SetStatus(http.StatusBadRequest)
			u.Data["json"] = models.Msg{
				StatusCd: 0,
				Message:  err.Error(),
			}
			u.ServeJSON()
		}
		u.Data["json"] = map[string]string{"uid": uid}
		u.ServeJSON()
	}
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /get/All [get]
func (u *UserController) GetAll() {
	// Check if user is logged in
	if u.CheckAuth() {
		users := implement.GetAllUsers()
		u.Data["json"] = users
		u.ServeJSON()
	}
}

// @Title Get Some User
// @Description get user by user info
// @Success 200 {object} models.User
// @router /get/One [post]
func (u *UserController) PostOneUser() {
	var user models.User
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
			var user models.User
			_ = json.Unmarshal(u.Ctx.Input.RequestBody, &user)
			uu, err := implement.UpdateUser(uid, &user)
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