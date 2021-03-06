package static

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Mount sets up static routes
func Mount(router *gin.Engine) {
	mountFile(router, "/bundle.js")
	mountFile(router, "/bundle.css")
	mountFile(router, "/global.css")
	mountFile(router, "/palm.png")
	mountFile(router, "/sw.js")
	mountFile(router, "/index.html")
	router.GET("/", serveClient)
	router.GET("/users/:userId", serveClient)
	router.GET("/reports/dip", serveClient)
}

func serveClient(c *gin.Context) {
	c.Status(http.StatusOK)
	c.File("./client/public/index.html")
}

func mountFile(router *gin.Engine, fileName string) {
	router.GET(fileName, func(c *gin.Context) {
		c.Status(http.StatusOK)
		c.File("./client/public" + fileName)
	})
}
