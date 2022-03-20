package controller

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"restapi/service/authService"
	"restapi/util"
)

type auth struct {
	Username string `json:"username" valid:"Required; MaxSize(50)"`
	Password string `json:"password" valid:"Required; MaxSize(50)"`
}

func Authenticate(c *gin.Context) {
	valid := validation.Validation{}

	var cred auth

	if err := c.BindJSON(&cred); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, "Some error occurred while parsing the data")
		return
	}

	a := auth{Username: cred.Username, Password: cred.Password}

	ok, err := valid.Valid(&a)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Username or Password"})
		return
	}

	authenticationService := authService.Auth{Username: cred.Username, Password: cred.Password}
	exist, err := authenticationService.Authenticate()

	if err != nil || !exist {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Username or Password"})
		return
	}

	token, err := util.GenerateToken(cred.Username, cred.Password)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, nil)
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}
