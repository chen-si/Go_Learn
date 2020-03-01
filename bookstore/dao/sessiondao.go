package dao

import (
	"bookstore/model"
	"bookstore/utils"
	"net/http"
)

//AddSession 向数据库中添加Session
func AddSession(sess *model.Session) error {
	//写sql语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}

//DeleteSession
func DeleteSession(sessID string) error {
	sqlStr := "delete from sessions where session_id = ?"
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}

//GetSession
func GetSession(sessId string) (*model.Session, error) {
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	row := inStmt.QueryRow(sessId)

	sess := &model.Session{}
	err = row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

//IsLogin 判断用户是否已经登录
func IsLogin(r *http.Request) (bool, *model.Session) {
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//根据cookievalue去数据库中查询与之对应的session
		session, _ := GetSession(cookie.Value)
		if session != nil && session.UserID > 0 {
			//已经登录
			return true, session
		}
	}
	//没有登录
	return false, nil
}
