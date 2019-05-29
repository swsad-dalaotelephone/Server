package userController

import (
	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	nickName := c.PostForm("nickName")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	openId := c.DefaultPostForm("openId", "")
	//check phone is not used
	_, err := userModel.GetUsersByStrKey("Phone", phone)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "phone is registered",
		})
		c.Error(errors.New("phone is registered"))
		return
	}
	// create new user
	user := userModel.User{}
	user.OpenId = openId
	user.NickName = nickName
	// encrypt password with MD5
	password = util.MD5(password)
	user.Password = password
	user.Phone = phone
	res := userModel.AddUser(user)

	if res {
		// successfully register
		userJson, err := util.StructToJson(user)
		if err != nil {
			log.ErrorLog.Println(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"msg":  "successfully register",
			"user": userJson,
		})
		c.Error(errors.New("successfully register"))
	} else {
		log.ErrorLog.Println("fail to register")
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "fail to register",
		})
		c.Error(errors.New("fail to register"))
	}
}
