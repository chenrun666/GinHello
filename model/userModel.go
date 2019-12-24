package model

import (
	"GinHello/initDB"
	"log"
)

type UserModel struct {
	Id            int
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password" binding:"required"`
	PasswordAgain string `form:"password-again"`
}

func (user *UserModel) Save() int64 {

	result, e := initDB.Db.Exec("insert into ginhello.user (email, password) values (?,?);", user.Email, user.Password)
	if e != nil {
		log.Panic("user insert error", e.Error())
	}

	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}

	return id

}

func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where email = ?;", user.Email)
	e := row.Scan(&u.Id, &u.Email, &u.Password)
	if e != nil {
		log.Panic(e)
	}

	return u
}
