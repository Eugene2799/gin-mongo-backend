package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"strings"
)

// GetHTML 获取数据
func GetHTML(c *gin.Context) {

	Url := c.PostForm("url")
	Url1 := strings.Split(Url, "//")[1]
	UrlOrigin := "https://" + strings.Split(Url1, "/")[0]

	client := &http.Client{}
	req, _ := http.NewRequest("GET", Url, nil)
	for k,v :=range c.Request.Header {
		//fmt.Printf(k)
		if k == "User-Agent" {
			req.Header.Set(k, v[0])
		}
		if k == "Origin" {
			//	使用origin v[0]值来限制请求来源
			fmt.Printf(v[0])
		}
	}
	req.Header.Set("Sec-Fetch-Dest", "document")
	req.Header.Set("Sec-Fetch-Mode", "navigate")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Origin", UrlOrigin)
	res, err := client.Do(req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	body, err := ioutil.ReadAll(res.Body)

	res.Body.Close()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": 500,
			"msg":    err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": 200,
		"msg":    "Success",
		"data":   string(body),
	})
}

