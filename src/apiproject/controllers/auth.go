package controllers

import (
	"apiproject/conf"
	"apiproject/implement"
	"apiproject/models"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	// Check if user is logged in
	sess := u.StartSession()
	var loginRequest models.LoginRequest
	json.Unmarshal(u.Ctx.Input.RequestBody, &loginRequest)
	log.Println(loginRequest)
	log.Println(loginRequest.UserId)
	log.Println(loginRequest.Password)
	count, err := implement.Login(loginRequest.UserId, loginRequest.Password)
	if err != nil {
		u.Ctx.Output.SetStatus(http.StatusInternalServerError)
		u.Data["json"] = models.Msg{
			StatusCd: 1,
			Message: err.Error(),
		}
		u.ServeJSON()
		return
	}
	if count == 0 {
		u.Ctx.Output.SetStatus(http.StatusUnauthorized)
		u.Data["json"] = models.Msg{
			StatusCd: 1,
			Message:  conf.IniConf.String("loginErr"),
		}
		u.ServeJSON()
		return
	}
	authKey := tokenGenerator()
	sess.Set(authKey, authKey)
	u.Data["json"] = models.LoginResponse{
		AuthCd: fmt.Sprintf("%v", authKey),
		UserId: loginRequest.UserId,
	}
	u.ServeJSON()
	return
}

// @Title logout
// @Description Logs user into the system
// @Success 200 {string} logout success
// @Failure 403 user not exist
// @router /logout [get]
func (u *UserController) Logout() {
	authCd := u.Ctx.Input.Header("authCd")

	sess := u.StartSession()
	authCdSes := sess.Get(authCd)
	if authCdSes != nil {
		sess.Delete(authCd)
		u.Data["json"] = models.Msg{
			StatusCd: 0,
			Message:  conf.IniConf.String("logoutSuccess"),
		}
		u.ServeJSON()
		return
	}
	u.Ctx.Output.SetStatus(http.StatusUnauthorized)
	u.Data["json"] = models.Msg{
		StatusCd: 1,
		Message:  conf.IniConf.String("loginErr"),
	}
	u.ServeJSON()
	return
}

func tokenGenerator() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (u *UserController) CheckAuth() bool {
	sess := u.StartSession()
	authCd := u.Ctx.Input.Header("authCd")

	authSes := sess.Get(authCd)
	log.Println(authCd)
	if authSes == nil {
		u.Ctx.Output.SetStatus(http.StatusUnauthorized)
		u.Data["json"] = models.Msg{
			StatusCd: 0,
			Message:  conf.IniConf.String("loginErr"),
		}
		u.ServeJSON()
		return false
	}
	return true
}