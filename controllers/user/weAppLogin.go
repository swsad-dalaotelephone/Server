package userController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/config"
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/medivhzhan/weapp"
)

/*
WeAppLogin : log in wechat app
if user exist, login auto and return 200 and user infomation
if user not exist , return 200 and "user is unregistered"
else return 400
require: code
return: msg, user, open_id
*/
func WeAppLogin(c *gin.Context) {
	code := c.Query("code")
	appID := config.WeAppConfig.AppID
	secret := config.WeAppConfig.Secret
	//get openID  and  session-key
	res, err := weapp.Login(appID, secret, code)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "error weapp code",
		})
		log.ErrorLog.Println(err)
		c.Error(errors.New("error weapp code"))
		return
	}

	//find user
	users, err := userModel.GetUsersByStrKey("open_id", res.OpenID)

	// if user is unregistered
	if len(users) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"msg":     "user is unregistered",
			"open_id": res.OpenID,
		})
		log.ErrorLog.Println("user is unregistered")
		c.Error(errors.New("user is unregistered"))
		return
	}

	user := users[0]
	session := sessions.Default(c)
	session.Set("user", user)
	err = session.Save()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "fail to generate session token",
		})
		log.ErrorLog.Println("fail to generate session token")
		c.Error(errors.New("fail to generate session token"))
	} else {
		userJson, err := util.StructToJsonStr(user)
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
