package main

import (
	"github.com/gin-gonic/gin"
	"github.com/99designs/gqlgen/handler"
	"github.com/yarncraft/mini-graphql/src"
	db "github.com/yarncraft/mini-graphql/src/db"

)

func graphqlHandler() gin.HandlerFunc {
	h := handler.GraphQL(src.NewExecutableSchema(src.Config{Resolvers: &src.Resolver{}}))

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func playgroundHandler() gin.HandlerFunc {
	h := handler.Playground("GraphQL", "/query")

	return func(c *gin.Context) {
		h.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	db.GetConnection()
	defer db.Close()

	gin.SetMode(gin.ReleaseMode)
    r := gin.Default()
	
	r.POST("/query", graphqlHandler())
	r.GET("/", playgroundHandler())
    r.Run()
}  