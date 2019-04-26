package userController

import (
	"baobaozhuan/models/user"
	"baobaozhuan/modules/log"
	"baobaozhuan/modules/util"
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
	user, err := userModel.GetUserByKey("Phone", phone)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "phone is registered",
		})
		c.Error(errors.New("phone is registered"))
		return
	}
	// encrypt password with MD5
	password = util.MD5(password)
	user, res := userModel.AddUser(openId, nickName, password, phone)
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
