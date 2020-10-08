package controllers

import (
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/astaxie/beego"
	"github.com/joho/godotenv"
)

type GeneratetokenController struct {
	beego.Controller
}

func (u *GeneratetokenController) Get() {
	godotenv.Load("/var/www/todo.allyapps.com/todo.env")
	url := os.Getenv("AUTH0URL")

	payload := strings.NewReader("{\"client_id\":\"" + os.Getenv("AUTH0CLIENTID") + "\",\"client_secret\":\"" + os.Getenv("AUTH0CLIENTSECRET") + "\",\"audience\":\"" + os.Getenv("AUTH0AUDIENCE") + "\",\"grant_type\":\"client_credentials\"}")
	req, _ := http.NewRequest("POST", url, payload)
	req.Header.Add("content-type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()

	u.Data["json"] = string(body)
	u.ServeJSON()
}
