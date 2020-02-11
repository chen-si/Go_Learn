package model

import (
	"encoding/json"
	"fmt"
	"ChatSystem/common/message"
	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动时就初始化一个UserDao实例
//把他做成全局变量，在需要使用时直接使用即可
var (
	MyUserDao *UserDao
)

//先定义一个UserDao 结构体 Dao:data access object
//完成对User 结构体的各种操作

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式，创建UserDao的实例
func NewUserDao(pool *redis.Pool) (userdao *UserDao) {
	userdao = &UserDao{
		pool: pool,
	}
	return
}

func (this *UserDao) getUserById(conn redis.Conn, id int) (user *User, err error) {
	//同过给定id去redis查询这个用户
	res, err := redis.String(conn.Do("hget", "users", id))
	if err != nil {
		//表示在users哈希中没有找到该用户
		if err == redis.ErrNil {
			err = ERROR_USER_NOTEXISTS
		}
		return
	}
	//这里我们要把res 反序列化成User实例
	user = &User{}
	err = json.Unmarshal([]byte(res), user)
	if err != nil {
		fmt.Println("json.Unmarshal([]byte(res), user) error:", err)
	}

	return
}

//完成登录的校验
//1、Login 完成用户的验证
//2、如果id和pwd都正确，则返回一个user实例
//3、如果id或pwd错误，返回错误提示信息
func (this *UserDao) Login(UserID int, UserPWD string) (user *User, err error) {
	//先从UserDao的链接池中取出一个链接
	conn := this.pool.Get()
	defer conn.Close()

	user, err = this.getUserById(conn, UserID)
	if err != nil {
		return
	}

	//这时候验证密码的正确性
	if UserPWD != user.UserPWD {
		err = ERROR_USER_PWD
		return
	}
	return
}

func (this *UserDao) Register(user *message.User) (err error) {
	conn := this.pool.Get()
	defer conn.Close()

	_, err = this.getUserById(conn, user.UserID)
	if err == nil {
		err = ERROR_USER_EXISTS
		return
	}

	//说明id在redis中不存在 可以完成注册
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	//可以入库了
	_, err = conn.Do("HSet", "users", user.UserID, string(data))
	if err != nil {
		fmt.Println("conn.Do(\"HSet\",\"users\",user.UserID,string(data)) error:", err)
		return
	}
	return
}
