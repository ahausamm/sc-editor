package api

import (
	"./helpers"
	"./lib"
	"./resources"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
)

func AuthenticateRequest() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("secure_string") == nil {
			secure_string := lib.CreateSecureString()
			session.Set("secure_string", secure_string)
			session.Save()
		}
		if !lib.IsSecureStringValid(session.Get("secure_string")) {
			c.AbortWithStatus(http.StatusUnauthorized);
		}
	}
}


var router *gin.Engine

func InitAPI() {

	gin.SetMode(gin.ReleaseMode)
	router = gin.Default()

	store := sessions.NewCookieStore([]byte("secret"))
	router.Use(sessions.Sessions("goeast-staticcms", store))
	router.Use(AuthenticateRequest())

	router.Static("/static", "./application/frontend/static")
	router.LoadHTMLGlob("./application/frontend/templates/*")

	router.GET("/", func(c *gin.Context) {
		session := sessions.Default(c)
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"hash":          "asdfljadsf3sfdjl3",
			"secure_string": session.Get("secure_string"),
		})
	})

	v1 := router.Group("api/v1")
	{
		v1.GET("/users", users.GetUsers)
		v1.GET("/users/:id", users.GetUser)
		v1.POST("/users", users.PostUser)
		v1.PUT("/users/:id", users.UpdateUser)
		v1.DELETE("/users/:id", users.DeleteUser)
	}

	Run()
	return

}

func Run() {
	router.Run(":8080")
}

func Init(userId string) {
	if(management.InitManagement(userId)){
		fmt.Println("ready to use")
		InitAPI()
	} else {
		fmt.Println("Error: Management unauthorized")
	}
}
