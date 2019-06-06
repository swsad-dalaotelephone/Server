package userController

import (
	"errors"
	"net/http"

	"github.com/swsad-dalaotelephone/Server/models/user"
	"github.com/swsad-dalaotelephone/Server/modules/log"
	"github.com/swsad-dalaotelephone/Server/modules/util"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	nickName := c.PostForm("nick_name")
	phone := c.PostForm("phone")
	password := c.PostForm("password")
	openId := c.DefaultPostForm("open_id", "")

	//check phone is not used
	users, err := userModel.GetUsersByStrKey("phone", phone)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": err.Error,
		})
		log.ErrorLog.Println(err)
		c.Error(err)
		return
	}

	if len(users) != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"msg": "phone is registered",
		})
		log.ErrorLog.Println("phone is registered")
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
	user, res := userModel.AddUser(user)

	if res {
		// successfully register
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
			"msg":  "successfully register",
			"user": userJson,
		})
		log.InfoLog.Println("successfully register")
	} else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"msg": "fail to register",
		})
		log.ErrorLog.Println("fail to register")
		c.Error(errors.New("fail to register"))
	}
}
