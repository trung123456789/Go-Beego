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
// @Param user_id body string true "User Id"
// @Param password body string true "password"
// @router /login [post]
func (u *UserController) Login() {
	// Check if user is logged in
	sess := u.StartSession()
	var loginRequest models.LoginRequest
	errJson := json.Unmarshal(u.Ctx.Input.RequestBody, &loginRequest)
	if errJson != nil {
		log.Println("err: ", errJson)
		u.Ctx.Output.SetStatus(http.StatusBadRequest)
		u.Data["json"] = models.Msg{
			StatusCd: 1,
			Message:  errJson.Error(),
		}
		u.ServeJSON()
		return
	}
	ok, err := implement.Login(loginRequest.UserId, loginRequest.Password)
	if !ok || err != nil {
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

// Auto generate token
func tokenGenerator() string {
	b := make([]byte, 32)
	rand.Read(b)
	return fmt.Sprintf("%x", b)
}

func (u *UserController) CheckAuth() bool {
	sess := u.StartSession()
	authCd := u.Ctx.Input.Header("authCd")

	authSes := sess.Get(authCd)
	if authSes == nil {
		u.Ctx.Output.SetStatus(http.StatusUnauthorized)
		u.Data["json"] = models.Msg{
			StatusCd: 0,
			Message:  conf.IniConf.String("loginErr"),
		}
		u.ServeJSON()
	}
	return true
}
