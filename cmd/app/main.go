package main

import (
	"flag"
	"github.com/BurntSushi/toml"
	"log"
	"smth/internal/app"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "config-path", "config/server.toml", "path to config file")
}

func main() {
	flag.Parse()

	_, err := toml.DecodeFile(configPath, configPath)

	config := app.NewConfig()

	s := app.New(config)
	if err = s.Start(); err != nil {
		log.Fatal(err)
	}

	//router := gin.Default()
	//router.Static("/assets/", "static/")
	//router.LoadHTMLGlob("templates/*.html")
	//
	//router.GET("/", mainPage)
	//router.GET("/auth", singUpPage)
	//router.POST("/auth/singup", toSingUpPage)

}

//func mainPage(c *gin.Context) {
//	c.HTML(200, "index.html", gin.H{
//		"title": "main page",
//	})
//}
//
//func singUpPage(c *gin.Context) {
//	c.HTML(http.StatusOK, "singup.html", gin.H{
//		"title": "sing-up page",
//	})
//}
//
//func toSingUpPage(c *gin.Context) {
//	user := service.User{Name: c.PostForm("name"), Email: c.PostForm("email"), Password: c.PostForm("password")}
//	fmt.Println(user)
//	user.Create()
//	c.Redirect(301, "/")
//}
