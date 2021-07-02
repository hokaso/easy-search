package router

import (
	a "easy-search/api"
	"github.com/gin-gonic/gin"
	"log"
)

func Init() {
	r := gin.Default()
	setupRouter(r)

	// Run the server
	if err := r.Run(":3000"); err != nil {
		log.Panicf("startup service failed, err:%v\n", err)
	}
}

func setupRouter(r *gin.Engine) {
	root := r.Group("/")
	//if gin.IsDebugging() {
	//	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//}
	//root.GET("", func(c *gin.Context) {
	//	c.String(200, "Welcome to Jax.")
	//})

	{
		api := root.Group("/api")
		api.GET("/qq", a.QueryUserQq)
		api.GET("/phone", a.QueryUserPhone)
	}
}
