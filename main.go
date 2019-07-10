package main

import (
	"ginLearn/pkg"
	log "ginLearn/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

func init() {
	pkg.Setup("")
	log.InitLog("info", pkg.AppConfig.LogPath)
	//err := redigo.Init()
	//if err != nil {
	//	panic(err)
	//}
}
func main() {
	gin.SetMode(gin.DebugMode)
	router := gin.Default()
	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	router.Use(gin.Logger())
	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	router.Use(gin.Recovery())

	//redigo.SetStrWithExpire("name", "fang", 60*60)
	//value, err := redigo.GetStr("name")
	//if err != nil {
	//	logger.Info("Read redis key name failed")
	//	panic("Read redis key name failed")
	//}
	//logger.Info("redis Name: %s", value)
	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	router.GET("/welcome", func(c *gin.Context) {
		firstname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")

		c.String(http.StatusOK, "Hello %s %s", firstname, lastname)
	})
	// Listen and serve on 0.0.0.0:8080
	err := router.Run(":8090")
	if err != nil {
		panic(err)
	}
}
