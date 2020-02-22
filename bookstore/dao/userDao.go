package dao

import (
	"bookstore/model"
	"bookstore/utils"
)

//CheckUsernameAndPassword 根据用户名和密码从数据控中查询一条记录
func CheckUsernameAndPassword(username string, password string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username=? and password=?"
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	return user, nil
}

//CheckUsername 根据用户名数据库中查询一条记录
func CheckUsername(username string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username=?"
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)

	return user, nil
}

//SaveUser 向数据库中插入用户信息
func SaveUser(username string, password string, email string) error {
	sqlStr := "insert into users(username,password,email) values(?,?,?)"

	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email)
	if err != nil {
		return err
	}
	return nil
}
