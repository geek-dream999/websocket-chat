package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"go-gin-chat/internal"
	"go-gin-chat/models"
	"go-gin-chat/routes"
	"go-gin-chat/views"
	"log"
	"net/http"
)

func init() {

	/*	var config = conf.Config{
		Port:           "8322",
		UplaodFilePath: "D:\\projects\\project-sn\\websocket-chat\\go-gin-chat\\source",
		CookieKey:      "4238uihfieh49r3453kjdfg",
		ServeType:      "GoServe",
		SmToken:        "xxxxxxxxxxx",
		DebugMod:       "true",

		Dsn: "root:mypassword.123@tcp(49.233.53.131:3307)/go_gin_chat?charset=utf8mb4&parseTime=True&loc=Local",
	}*/
	internal.InitViper()
	models.InitDB()
}

func main() {
	// 关闭debug模式
	gin.SetMode(gin.ReleaseMode)

	port := viper.GetString(`app.port`)
	router := routes.InitRoute()

	//加载模板文件
	router.SetHTMLTemplate(views.GoTpl)

	//go_ws.CleanOfflineConn()

	log.Println("监听端口", "http://127.0.0.1:"+port)

	http.ListenAndServe(":"+port, router)
}
