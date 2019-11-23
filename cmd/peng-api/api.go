package main

import (
	"github.com/labstack/echo"
	"net/http"
	"yonghochoi.com/depthon-2019/model/emoticon"
	"yonghochoi.com/depthon-2019/model/user"
)

var emoticonDatas []*emoticon.Emoticon

func InitSample() {
	emoticonDatas = []*emoticon.Emoticon{
		emoticon.New("good", "기분 좋음", "http://localhost"),
		emoticon.New("angry", "화남", "http://localhost"),
		emoticon.New("boring", "따분함", "http://localhost"),
		emoticon.New("sad", "슬픔", "http://localhost"),
		emoticon.New("tired", "지침", "http://localhost"),
	}
}

func GetEmoticonDatas(c echo.Context) error {
	return c.JSON(http.StatusOK, emoticonDatas)
}

func GetMoods(c echo.Context) error {
	return c.String(http.StatusOK, version)
}

func Join(c echo.Context) error {
	var user user.User
	if err := c.Bind(&user); err != nil {
		return err
	}

	user.Join()
	return c.JSON(http.StatusOK, user)
}
