package controllers

import (
	"apiproject/conf"
	"apiproject/implement"
	"apiproject/models"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"net/http"
)

// @Param body body models.LoginRequest true "User Id"
// @Success 201 {string} authCd
// @Failure 403 body message error login
// @router /login [post]
func (u *UserController) Login() {
	logs.Start("LOGIN")
	sess := u.StartSession()
	var loginRequest models.LoginRequest

	defer u.deferFunc("LOGIN")

	errJson := json.Unmarshal(u.Ctx.Input.RequestBody, &loginRequest)
	if errJson != nil {
		u.Ctx.Output.SetStatus(http.StatusBadRequest)
		u.Data["json"] = models.Msg{
			StatusCd: 1,
			Message:  errJson.Error(),
		}
		panic(0)
	}
	err := implement.Validate.Struct(loginRequest)
	if err != nil {
		u.Ctx.Output.SetStatus(http.StatusBadRequest)
		u.Data["json"] = models.Msg{
			StatusCd: 1,
			Message:  err.Error(),
		}
		panic(0)
	}

	ok, err := implement.Login(loginRequest.UserId, loginRequest.Password)
	if !ok || err != nil {
		u.Ctx.Output.SetStatus(http.StatusUnauthorized)
		u.Data["json"] = models.Msg{
			StatusCd: 1,
			Message:  conf.IniConf.String("loginErr"),
		}
		panic(0)
	}
	authKey := tokenGenerator()
	sess.Set(authKey, authKey)
	u.Data["json"] = models.LoginResponse{
		AuthCd: fmt.Sprintf("%v", authKey),
		UserId: loginRequest.UserId,
	}
	panic(0)
}

// @Title logout
// @Description Logs user into the system
// @Success 200 {string} logout success
// @Failure 403 user not exist
// @router /logout [get]
func (u *UserController) Logout() {
	authCd := u.Ctx.Input.Header("authCd")
	defer u.deferFunc("LOGOUT")
	sess := u.StartSession()
	authCdSes := sess.Get(authCd)
	if authCdSes != nil {
		sess.Delete(authCd)
		u.Data["json"] = models.Msg{
			StatusCd: 0,
			Message:  conf.IniConf.String("logoutSuccess"),
		}
		panic(0)
	}
	u.Ctx.Output.SetStatus(http.StatusUnauthorized)
	u.Data["json"] = models.Msg{
		StatusCd: 1,
		Message:  conf.IniConf.String("loginErr"),
	}
	panic(0)
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

func (u *UserController) deferFunc(apiName string) {
	if r := recover(); r != nil {
		logs.End(apiName)
		u.ServeJSON()
	}
}
