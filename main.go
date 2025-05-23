package main

import (
	"net/http"
	"os"

	_ "github.com/KITTTPOB-bank/hospitalapi/docs"
	"github.com/KITTTPOB-bank/hospitalapi/initializers"
	"github.com/KITTTPOB-bank/hospitalapi/router"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description ตัวอย่างรูปแบบการใส่ข้อมูล "Bearer <token>"
func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	router.SetupRouter(r)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	http.ListenAndServe(getPort(),
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Accept", "AllowedOrigins", "AllowedHeaders", "AllowedMethods", "Content-Length", "Accept-Encoding", "X-Requested-With"}),
		)(r))
}

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8010"
	}
	return ":" + port
}
