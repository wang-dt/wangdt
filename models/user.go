package models

type User struct {
	Id       int
	Username string
	Age      string
	Address  string
}

// 配置操作数据库标名称
func (User) TableName() string {
	return "user"
}
