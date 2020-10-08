package controllers

import (
	"encoding/json"
	"todolist2/models"

	"fmt"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) Post() {
	var User models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &User)
	uid := models.AddUser(User)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

func (u *UserController) Get() {
	User := models.GetAllUsers()
	u.Data["json"] = User
	u.ServeJSON()
}

func (u *UserController) Delete() {
	uid := u.GetString("uid")
	fmt.Println(uid)
	models.DeleteUser(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

func (u *UserController) Put() {
	var User models.User
	json.Unmarshal(u.Ctx.Input.RequestBody, &User)
	models.EditUser(User)
	//u.Data["json"] = map[string]string{"uid": uid}
	u.Data["json"] = "edit success!"
	u.ServeJSON()
}
