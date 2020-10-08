package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/astaxie/beego/context"
	"github.com/dgrijalva/jwt-go"

	_ "todolist2/routers"

	"github.com/astaxie/beego"
    
    "github.com/joho/godotenv"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func main() {

	var AuthFilter = func(ctx *context.Context) {

		header := strings.Split(ctx.Input.Header("Authorization"), " ")
		if len(header) != 2 || header[0] != "Bearer" {
			ctx.Abort(401, "Not authorized")
		}
        
        godotenv.Load("/var/www/todo.allyapps.com/todo.env")
        
		token, err := jwt.Parse(header[1], func(token *jwt.Token) (interface{}, error) {

            
			aud := os.Getenv("AUTH0AUDIENCE")
			checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
			if !checkAud {
				return token, errors.New("Invalid audience.")
			}

			iss := os.Getenv("AUTH0URL2")
			checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
			if !checkIss {
				return token, errors.New("Invalid issuer.")
			}

			cert, err := GetPemCert(token)
			if err != nil {
				panic(err.Error())
			}

			result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
			return result, nil

		})

		if err != nil {
			fmt.Println(err)
			ctx.Abort(401, "Not authorized")
		}

		if token.Valid {
			fmt.Println("valid")
		}

	}

	beego.InsertFilter("/todo", beego.BeforeRouter, AuthFilter)
	beego.InsertFilter("/user", beego.BeforeRouter, AuthFilter)

	beego.Run()
}

func checktoken(token *jwt.Token) (interface{}, error) {
	// Verify 'aud' claim
	aud := os.Getenv("AUTH0AUDIENCE")
	checkAud := token.Claims.(jwt.MapClaims).VerifyAudience(aud, false)
	if !checkAud {
		return token, errors.New("Invalid audience.")
	}
	// Verify 'iss' claim
	iss := os.Getenv("AUTH0URL2")
	checkIss := token.Claims.(jwt.MapClaims).VerifyIssuer(iss, false)
	if !checkIss {
		return token, errors.New("Invalid issuer.")
	}

	cert, err := GetPemCert(token)
	if err != nil {
		panic(err.Error())
	}

	result, _ := jwt.ParseRSAPublicKeyFromPEM([]byte(cert))
	return result, nil
}

func GetPemCert(token *jwt.Token) (string, error) {
	cert := ""
	resp, err := http.Get("https://dev-5pey12lq.us.auth0.com/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
