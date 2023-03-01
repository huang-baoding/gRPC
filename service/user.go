//将用户添加到我们的系统，并向登录用户添加服务并返回JWT访问令牌
package service

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username       string //用户名
	HashedPassword string //哈希密码
	Role           string //角色
}

// NewUser returns a new user
func NewUser(username string, password string, role string) (*User, error) {
	//我们不应该在系统中存储明文密码，所以使用bcrypt对密码进行哈希处理
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost) //使用默认成本
	if err != nil {
		return nil, fmt.Errorf("cannot hash password: %w", err)
	}

	user := &User{ //创建一个新的用户对象
		Username:       username,
		HashedPassword: string(hashedPassword),
		Role:           role,
	}

	return user, nil
}

// 检查给定的密码是否正确
func (user *User) IsCorrectPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	return err == nil
}

// 克隆用户，在后面将用户存储在内存中有用
func (user *User) Clone() *User {
	return &User{
		Username:       user.Username,
		HashedPassword: user.HashedPassword,
		Role:           user.Role,
	}
}
