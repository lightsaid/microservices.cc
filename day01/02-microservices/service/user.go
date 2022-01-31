package main

import (
	"errors"
	"strings"
	"fmt"
)

type User struct{
	ID int
	Name string
	Password string
}

var users = []User{
	{
		ID: 1,
		Name: "张三",
		Password: "abc123",
	},
}

func (this *User) Register(user User, resp *User) error{
	user.Name = strings.TrimSpace(user.Name)
	if user.Name == "" || user.Password == "" {
		return errors.New("用户名或者密码不能为空")
	}else{
		if len(users) == 0 {
			user.ID = 1
		}else{
			user.ID = len(users) + 1
		}
	}
	*resp = user
	users = append(users, user)
	return nil
}

func (this *User) Login(user User, resp *User) error{
	fmt.Println("login user:",users)
	for i := 0; i<len(users); i++{
		fmt.Printf("users[%d]=%v", i, users[i])
		if user.Name == users[i].Name && user.Password == users[i].Password {
			*resp = users[i]
			return nil
		}
	}
	return errors.New("用户名和密码不匹配")
}



