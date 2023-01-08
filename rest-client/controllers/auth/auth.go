package auth

import (
	"api/jwthelper"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type userTable struct {
	Users []User
}

var table *userTable

func New() *userTable {
	if table == nil {
		table = &userTable{
			Users: []User{},
		}
	}

	return table
}

func Login(ctx *gin.Context) {
	tbl := New()

	userBody := User{}

	if err := ctx.ShouldBindJSON(&userBody); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var isUserValid bool
	httpStatus := http.StatusUnauthorized
	jwtToken := ""

	for _, u := range tbl.Users {
		if u.Username == userBody.Username && u.Password == userBody.Password {

			tokenString, err := jwthelper.GenerateJWT()
			if err != nil {
				fmt.Println(err)
				ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				jwtToken = ""
				return
			}

			jwtToken = tokenString

			isUserValid = true
			httpStatus = http.StatusOK

			break
		}
	}

	ctx.JSON(httpStatus, gin.H{
		"success": isUserValid,
		"token":   jwtToken,
	})

}
