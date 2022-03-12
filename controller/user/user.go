package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"login_register_demo/config"
	"login_register_demo/middleware"
	"login_register_demo/model"
	"login_register_demo/utils/errmsg"
	"time"
)

// changePassword
func ChangePasswordG(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	passWord := c.Request.URL.Query().Get("password")
	newPassWord := c.Request.URL.Query().Get("newpassword")
	//查询列表
	st2 := new(model.User)
	result, err := config.Engine.Where("username=?", userName).Get(st2)
	fmt.Println("查询结果为", result)
	if err != nil {
		fmt.Println(err)
	}
	if userName != st2.UserUsername {
		// 无此用户
		c.JSON(200, gin.H{
			"success": false,
			"code":    400,
			"msg":     "无此用户",
		})
	} else {
		// 密码是否匹配
		if passWord != st2.UserPassword {
			fmt.Println("password error")
			c.JSON(200, gin.H{
				"success": false,
				"code":    400,
				"msg":     "密码错误",
			})
		} else {
			//修改密码
			mm, err := config.Engine.Exec("update users set password = ? where username = ?", newPassWord, userName)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(mm)
			fmt.Println("change password success")
			c.JSON(200, gin.H{
				"success": true,
				"code":    200,
				"msg":     "修改成功",
			})
		}
	}
}

// deleteUser
func DeleteUsernameG(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	passWord := c.Request.URL.Query().Get("password")
	//查询列表
	st2 := new(model.User)
	result, err := config.Engine.Where("username=?", userName).Get(st2)
	fmt.Println("查询结果为", result)
	if err != nil {
		fmt.Println(err)
	}
	if userName != st2.UserUsername {
		// 无此用户
		c.JSON(200, gin.H{
			"success": false,
			"code":    400,
			"msg":     "无此用户",
		})
	} else {
		// 密码是否匹配
		if passWord != st2.UserPassword {
			fmt.Println("password error")
			c.JSON(200, gin.H{
				"success": false,
				"code":    400,
				"msg":     "密码错误",
			})
		} else {
			//删除账号
			mm, err := config.Engine.Where("username=?", userName).Delete(st2)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(mm)
			fmt.Println("delete account success")
			c.JSON(200, gin.H{
				"success": true,
				"code":    200,
				"msg":     "删除成功",
			})
		}
	}
}

//Login Test
func UserLoginT(c *gin.Context) {
	var user model.User
	//参数绑定
	if err := c.ShouldBind(&user); err != nil {
		log.Println(err.Error())
		code := errmsg.ERROR_JSON_TYPE_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}
	//查询数据库中有无当前用户名
	st2 := new(model.User) //存放查询结果
	userName := user.UserUsername
	passWord := user.UserPassword
	group_id := user.GroupId
	_, err := config.Engine.Where("user_username=?", userName).Get(st2)
	//fmt.Println("查询结果为", result)
	if err != nil {
		//code := errmsg.ERROR_DATABASE_WRONG
		//c.JSON(200, gin.H{
		//	"success": false,
		//	"code":    code,
		//	"msg":     errmsg.GetErrMsg(code),
		//})
		fmt.Println(err)
	}
	if userName != st2.UserUsername {
		// 无此用户
		code := errmsg.ERROR_USER_NOT_EXIST
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
	} else {
		// 密码是否匹配
		if passWord != st2.UserPassword {
			code := errmsg.ERROR_PASSWORD_WRONG
			c.JSON(200, gin.H{
				"success": false,
				"code":    code,
				"msg":     errmsg.GetErrMsg(code),
			})
		} else {
			// 为用户生成token
			token, code := middleware.SetToken(userName, st2.UserId)
			// 查询此用户是管理员还是普通用户，管理员为1,普通用户为2
			token = group_id + token
			if code != 200 {
				c.JSON(201, gin.H{
					"success": true,
					"code":    403,
					"msg":     "token生成失败！",
				})
			} else {
				c.JSON(200, gin.H{
					"success": true,
					"code":    200,
					"msg":     "登录成功",
					"token":   token,
				})
			}
		}
	}

}

// Register Test
func UserRegisterT(c *gin.Context) {
	var user model.User

	//参数绑定
	// 参数绑定与校验
	if err := c.ShouldBind(&user); err != nil {
		log.Println(err.Error())
		code := errmsg.ERROR_JSON_TYPE_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}
	fmt.Println(user)
	stu := new(model.User)
	//验证用户名是否已经存在，存在返回错误
	result, err := config.Engine.Where("user_username=?", user.UserUsername).Get(stu)
	if err != nil {
		log.Println(err.Error())
		code := errmsg.ERROR_JSON_TYPE_WRONG
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
		return
	}
	if result {
		code := errmsg.ERROR_USERNAME_USED
		c.JSON(200, gin.H{
			"success": false,
			"code":    code,
			"msg":     errmsg.GetErrMsg(code),
		})
	} else {
		user.CreateTime = time.Now()
		affected, err := config.Engine.Insert(user)
		if err != nil {
			fmt.Println(err)
		}
		if affected != 1 {
			c.JSON(200, gin.H{
				"success": false,
			})
		} else {
			c.JSON(200, gin.H{
				"success":  true,
				"username": user.UserUsername,
				"msg":      "Register success",
			})
		}
	}
	//将用户信息插入到数据库表中

}
