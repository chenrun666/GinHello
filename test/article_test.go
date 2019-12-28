package test

import (
	"GinHello/model"
	"bytes"
	"encoding/json"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInsertArticle(t *testing.T) {
	article := model.Article{
		Id: 1,
		Type: "gin",
		Content: "hello gin",
	}
	marshal, _ := json.Marshal(article)
	w := httptest.NewRecorder()
	req := httptest.NewRequest(http.MethodPost, "/article", bytes.NewBufferString(string(marshal)))
	req.Header.Add("Content-Type", "application/json")
	router.ServeHTTP(w, req)
	assert.Equal(t, w.Code, http.StatusOK)
	assert.Equal(t, "{id:-1}", w.Body.String())
}
