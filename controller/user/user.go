package user

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"login_register_demo/config"
	"login_register_demo/middleware"
	"login_register_demo/model"
)

// @userLogin
func UserLoginG(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	passWord := c.Request.URL.Query().Get("password")
	group_id := c.Request.URL.Query().Get("group_id")
	//查询列表
	st2 := new(model.User)
	result, err := config.Engine.Where("user_username=?", userName).Get(st2)
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
			c.JSON(200, gin.H{
				"success": false,
				"code":    400,
				"msg":     "密码错误",
			})
		} else {

			// 为用户生成token
			token,code:=middleware.SetToken(userName)
			// 查询此用户是管理员还是普通用户，管理员为1,普通用户为2
			token=group_id+token
			if code!=200{
				c.JSON(201, gin.H{
					"success": true,
					"code":    403,
					"msg":     "token生成失败！",
				})
			}else{
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

// @userRegister
func UserRegisterG(c *gin.Context) {
	userName := c.Request.URL.Query().Get("username")
	passWord := c.Request.URL.Query().Get("password")
	//查询列表
	st2 := new(model.User)
	result, err := config.Engine.Where("user_username=?", userName).Get(st2)
	fmt.Println("查询结果为", result)
	if err != nil {
		fmt.Println(err)
	}
	if userName != st2.UserUsername{
		// 无此用户
		st2.UserUsername = userName
		st2.UserPassword = passWord
		affected, err := config.Engine.Insert(st2)
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
				"username": userName,
				"msg":      "Register success",
			})
		}
	} else {
		// 已存在用户，注册失败
		fmt.Println("Already has one exsit account!")
		c.JSON(200, gin.H{
			"code":    400,
			"success": false,
			"msg":     "用户名已被注册",
		})
	}
}

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
		if passWord != st2.UserPassword{
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