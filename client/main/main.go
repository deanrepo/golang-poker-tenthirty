package main

import (
	"fmt"
	"log"

	"tenthirty/client/process"
)

// User account
var email string

// User password
var userPwd string

func main() {

	/* // reader用于读取字符串
	reader := bufio.NewReader(os.Stdin) */

	// User's choice
	var key int

	// 显示主界面
loop:
	for {
		fmt.Println("-------------------棋牌游戏十点半-------------------")
		fmt.Println("                   1. 登录")
		fmt.Println("                   2. 注册")
		fmt.Println("                   3. 退出")
		fmt.Println("=>请选择(1-3):")

		fmt.Scanf("%d\n", &key)

		switch key {
		case 1:
			fmt.Println("-------------------登录-------------------")
			// TODO: LOGIN BUSINESS LOGIC.
			fmt.Print("请输入用户邮箱：")
			_, err := fmt.Scanf("%s\n", &email)
			if err != nil {
				log.Printf("input user account err: %v\n", err)
				continue
			}

			fmt.Print("请输入用户密码：")
			_, err = fmt.Scanf("%s\n", &userPwd)
			if err != nil {
				log.Printf("input user password err: %v\n", err)
				continue
			}

			err = process.Login(email, userPwd)
			if err != nil {
				log.Printf("Login err: %v.\n", err)
				continue
			}

		case 2:
			fmt.Println("-------------------注册-------------------")
		case 3:
			fmt.Println("-------------------退出-------------------")
			break loop
		default:
			fmt.Println("您输入有误， 请重新输入！")
		}
	}
}
