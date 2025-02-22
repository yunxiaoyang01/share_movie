package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func main() {
	// 设置管理员账户信息
	username := "admin"
	password := "admin123"
	email := "admin@example.com"
	nickname := "管理员"

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Printf("密码加密失败: %v\n", err)
		return
	}

	// 生成当前时间
	now := time.Now().Format("2006-01-02 15:04:05")

	// 生成SQL插入语句
	sql := fmt.Sprintf(`INSERT INTO users (username, password, email, nickname, role, status, created_at, updated_at) 
	VALUES ('%s', '%s', '%s', '%s', 2, 1, '%s', '%s');`,
		username, string(hashedPassword), email, nickname, now, now)

	fmt.Println("-- 创建管理员账户的SQL语句：")
	fmt.Println(sql)
}