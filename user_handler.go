package main

import (
	"encoding/json"
	"net/http"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "请求格式错误", http.StatusBadRequest)
		return
	}

	if err := CreateUser(u.Username, u.Password); err != nil {
		http.Error(w, "注册失败，可能用户名重复", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"message":"注册成功"}`))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	var u User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		http.Error(w, "请求格式错误", http.StatusBadRequest)
		return
	}

	ok, err := CheckUserCredentials(u.Username, u.Password)
	if err != nil || !ok {
		http.Error(w, "用户名或密码错误", http.StatusUnauthorized)
		return
	}

	w.Write([]byte(`{"message":"登录成功"}`))
}
