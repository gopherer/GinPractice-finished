package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type People struct {
	Name string `form:"name" json:"name"`
}

func main() {
	engine := gin.Default()
	var p People
	engine.POST("/ping", func(context *gin.Context) {
		fmt.Println(context.FullPath())
		context.Bind(&p)
		fmt.Println(context.ContentType())
		context.String(200, "hello")
		fmt.Println(p)
	})

	if err := engine.Run(); err != nil {
		log.Fatal(err)
	} // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
