package model

import (
	"GinHello/initDB"
	"log"
)

type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}


func (a *Article) Insert() int {
	result, e := initDB.Db.Exec("insert into article (type, content) values (?, ?);", a.Type, a.Content)
	if e != nil {
		log.Panic("添加文章失败", e.Error())
	}
	i, _ := result.LastInsertId()

	return int(i)
}