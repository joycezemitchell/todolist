package controllers

import (
	"encoding/json"
	"todolist2/models"

	"fmt"

	"github.com/astaxie/beego"
)

type TodoController struct {
	beego.Controller
}

func (u *TodoController) Post() {
	var todo models.Todo
	json.Unmarshal(u.Ctx.Input.RequestBody, &todo)
	uid := models.AddTodo(todo)
	u.Data["json"] = map[string]string{"uid": uid}
	u.ServeJSON()
}

func (u *TodoController) Get() {
	todo := models.GetAllTodos()
	u.Data["json"] = todo
	u.ServeJSON()
}

func (u *TodoController) Delete() {
	uid := u.GetString("uid")
	fmt.Println(uid)
	models.DeleteTodo(uid)
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

func (u *TodoController) Put() {
	var todo models.Todo
	json.Unmarshal(u.Ctx.Input.RequestBody, &todo)
	models.EditTodo(todo)
	//u.Data["json"] = map[string]string{"uid": uid}
	u.Data["json"] = "edit success!"
	u.ServeJSON()
}
