package models

import (
	"github.com/astaxie/beego/orm"     // 操作mysql的
	_ "github.com/go-sql-driver/mysql" // 驱动
)

var (
	db orm.Ormer
)

// UserInfo 这个结构体要和数据库表的名字（user_info）对应
//（数据库会自动把大写的变成小写并增加下划线）
type UserInfo struct {
	ID       int64
	Username string
	Passwd   string
}

func init() {
	orm.Debug = true // 调试模式下会打印出sql语句
	// 别名连接数据库只能调用一次，再次调用会报错defaulty已经注册不能reuse
	// testmodel是操作的数据库名 mysql是使用的驱动名 default是连接数据库的一个别名
	orm.RegisterDataBase("default", "mysql", "root:@tcp(127.0.0.1:3306)/testmodel?charset=utf8", 30)
	orm.RegisterModel(new(UserInfo)) // 创建一个user_info表

	db = orm.NewOrm()
}

// AddUser .
func AddUser(user *UserInfo) (int64, error) {
	// 插入  想要插入成功，需要事先在数据库里面创建表user_info
	// create table `user_info` (`i_d` int(11) auto_increment, `username` varchar(20) DEFAULT '', `passwd` varchar(11), primary key (i_d));
	return db.Insert(user)
}

// ReadAllUserInfo .
func ReadAllUserInfo(users *[]UserInfo) {
	qb, _ := orm.NewQueryBuilder("mysql")
	qb.Select("*").From("user_info")
	sql := qb.String()
	db.Raw(sql).QueryRows(users)
}

// ReadUserInfo 读取数据库
func ReadUserInfo(user *UserInfo) error {
	return db.Read(user)
}

// UpdateUserInfo .
func UpdateUserInfo(user *UserInfo) (int64, error) {
	return db.Update(user)
}

// DeleteUserInfo .
func DeleteUserInfo(user *UserInfo) (int64, error) {
	return db.Delete(user)
}
