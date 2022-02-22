package utils

// import (
// 	"encoding/json"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// //默认英文
// const defaultLanguage = "zh"
// const defaultCountry = "CN"

// var defaultRetLanguage = defaultLanguage + "-" + defaultCountry

// func HttpLanguage(l string) string {
// 	length := len(l)
// 	if length < 2 {
// 		return defaultRetLanguage
// 	}
// 	language := l[:2] //

// 	switch language {
// 	case "en":
// 		return "en"
// 	case defaultLanguage:
// 		if length >= 5 {
// 			country := l[3:5]
// 			return defaultLanguage + "-" + country
// 		}
// 		return defaultRetLanguage
// 	default:
// 		return language
// 	}
// }

// func HttpResponseError(c *gin.Context, err error) {
// 	var serviceErr ServiceError

// 	JsonErr := json.Unmarshal([]byte(err.Error()), &serviceErr)

// 	if JsonErr != nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":   500,
// 			"detail": err.Error(),
// 		})
// 		return
// 	}

// 	c.JSON(http.StatusOK, gin.H{
// 		"code":   serviceErr.Code,
// 		"detail": serviceErr.Detail,
// 	})
// }

// func HttpResponse403(c *gin.Context) {
// 	c.JSON(http.StatusOK, gin.H{
// 		"code":   403,
// 		"detail": "没有权限访问，请先登录",
// 	})
// }

// func HttpResponseSend(c *gin.Context, status int, data interface{}, msg string) {
// 	if data == nil {
// 		c.JSON(http.StatusOK, gin.H{
// 			"code":   status,
// 			"detail": msg,
// 		})
// 	} else {
// 		c.JSON(http.StatusOK, data)
// 	}

// }
