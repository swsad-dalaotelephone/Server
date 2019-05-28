package userController

import (
	"baobaozhuan/config"
	"baobaozhuan/models/user"
	"baobaozhuan/modules/log"
	"baobaozhuan/modules/util"
	"errors"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp"
)

/*
weapp login
if user exist, login auto and return 200 and user infomation
if user not exist , return 200 and "user is unregistered"
else return 400
*/
func WeAppLogin(c *gin.Context) {
	code := c.Query("code")
	appID := config.WeAppConfig.AppID
	secret := config.WeAppConfig.Secret
	//get openID  and  session-key
	res, err := weapp.Login(appID, secret, code)
	if err != nil {
		log.ErrorLog.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error weapp code",
		})
		c.Error(errors.New("error weapp code"))
		return
	}
	//find user
	users, err := userModel.GetUsersByStrKey("OpenId", res.OpenID)
	// if user is unregistered
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"msg":    "user is unregistered",
			"openId": res.OpenID,
		})
		c.Error(errors.New("user is unregistered"))
		return
	}
	user := users[0]
	session := sessions.Default(c)
	session.Set("user", user)
	err = session.Save()
	if err != nil {
		log.ErrorLog.Println("fail to generate session token")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "fail to generate session token",
		})
		c.Error(errors.New("fail to generate session token"))
	} else {
		userJson, err := util.StructToJson(user)
		if err != nil {
			log.ErrorLog.Println("fail to convert user data to json")
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "successfully login",
			"user": userJson,
		})
		c.Error(errors.New("successfully login"))
	}
}
