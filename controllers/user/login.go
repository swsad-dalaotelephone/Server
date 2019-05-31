package userController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

/*
weapp login
if user exist, login auto and return 200 and user infomation
if user not exist , return 200 and "user is unregistered"
if password error , return 401 and "Authentication failed"
else return 400
*/
func Login(c *gin.Context) {
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	//find user
	users, err := userModel.GetUsersByStrKey("phone", phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error(),
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	// if user is unregistered
	if len(users) == 0 {

		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "phone is unregistered",
		})
		log.ErrorLog.Println("phone is unregistered")
		c.Error(errors.New("phone is unregistered"))
		return
	}
	user := users[0]
	// encrypt password with MD5
	password = util.MD5(password)
	// if password error
	if password != user.Password {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "username or password is incorrect",
		})
		log.ErrorLog.Println("username or password is incorrect")
		c.Error(errors.New("username or password is incorrect"))
		return
	}
	session := sessions.Default(c)
	userJson, err := util.StructToJson(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "json conver error",
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}
	session.Set("user", userJson)
	err = session.Save()
	if err != nil {
		log.ErrorLog.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "fail to generate session token",
		})
		log.ErrorLog.Println("fail to generate session token")
		c.Error(errors.New("fail to generate session token"))
	} else {
		userJson, err := util.StructToJson(user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"msg": err.Error(),
			})
			log.ErrorLog.Println(err)
			c.Error(err)
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "successfully login",
			"user": userJson,
		})
		log.InfoLog.Println("successfully login")
	}
}
