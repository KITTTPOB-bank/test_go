package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/KITTTPOB-bank/hospitalapi/initializers"
	"github.com/KITTTPOB-bank/hospitalapi/router"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/handlers"
)

func init() {
	initializers.LoadEnvs()
	initializers.ConnectDB()
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	router.SetupRouter(r)
	http.Handle("/", r)
	http.ListenAndServe(getPort(),
		handlers.CORS(
			handlers.AllowedOrigins([]string{"*"}),
			// handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Accept", "AllowedOrigins", "AllowedHeaders", "AllowedMethods", "Content-Length", "Accept-Encoding", "X-Requested-With"}),
		)(r))

}
func getPort() string {

	port := os.Getenv("PORT")
	fmt.Println(port)
	if port == "" {
		port = "8010"
	}
	return ":" + port
}
