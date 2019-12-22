package test

import (
	"GinHello/initRouter"
	"github.com/magiconair/properties/assert"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"
)

func TestUserSave(t *testing.T) {
	username := "chenrun"
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user/"+username, nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户名: "+username+"已保存", w.Body.String())
}

func TestUserSaveByQuery(t *testing.T) {
	username := "chenrun"
	age := 20
	router := initRouter.SetupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/user?name="+username+"&age="+strconv.Itoa(age), nil)
	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "用户: "+username+"年龄: "+strconv.Itoa(age)+"已保存", w.Body.String())
}
