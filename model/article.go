package model

import (
	"GinHello/initDB"
)

type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (article Article) TableName () string {
	return "article"
}

func (article Article) Insert() int {
	//result, e := initDB.Db.Exec("insert into article (type, content) values (?, ?);", a.Type, a.Content)
	create := initDB.Db.Create(&article)
	if create.Error != nil {
		return 0
	}

	return 1
}

func (article Article) FindAll() []Article {
	var articles []Article
	initDB.Db.Find(&articles)
	return articles
}

func (article Article) FindById() Article {
	initDB.Db.First(&article, article.Id)
	return article
}

func (article Article) DeleteOne() {
	initDB.Db.Delete(article)
}