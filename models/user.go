package models

type User struct {
	Id       int
	Username string
	Age      string
	Address  string
}

/**
 * @Author wangdt
 * @Description //TODO ${end}
 * @Date 2022-09-15 19:35:04
 * @return string
 **/
// 配置操作数据库标名称
func (User) TableName() string {
	return "user"
}
