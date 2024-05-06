package main

import (
	"crypto/tls"
	"net/http"
	"os"

	"github.com/cloudflare/cfssl/log"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// init function to initialize useful variables, executes only once
func init() {
	// read environment variables
	log.Info("Reading environment variables...")
	err := godotenv.Load()
	// when not able to load .env file, throw an FATAL error
	if err != nil {
		log.Fatal(err)
	}
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "application/json")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, X-Max")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	}
}

func setupRouter() *gin.Engine {
	log.Info("initializing gin router...")

	// set Production mode
	gin.SetMode(gin.ReleaseMode)

	// initialize gin router with default setting
	router := gin.Default()

	// router.LoadHTMLGlob("templates/**/*.html")

	// register error logger
	router.Use(gin.ErrorLogger())

	router.Use(CORSMiddleware())
	// register all REST API endpoints

	router.GET("/", func(c *gin.Context) {
		c.Writer.Header().Set("Content-Type", "text/html")
		c.File("templates/index.html")

	})

	router.GET("/poodle.js", func(c *gin.Context) {
		c.File("templates/poodle.js")
	})

	return router

}

// main function of the REST API server
func main() {

	// start the REST API server here
	router := setupRouter()

	go func() {

	}()

	port := os.Getenv("APP_PORT")

	if port == "" {
		port = "9000"
	}
	port = ":" + port
	log.Infof("starting server on port: %v", port)

	privateKey := os.Getenv("SSL_PRIVATE_KEY_FILE_PATH")
	publicKey := os.Getenv("SSL_PUBLIC_KEY_FILE_PATH")
	if privateKey == "" || publicKey == "" {
		log.Fatal("missing private key and public key")
	}
	server := &http.Server{
		Addr:    port,
		Handler: router,
		TLSConfig: &tls.Config{
			// MinVersion: tls.VersionSSL30,
			// MaxVersion: tls.VersionSSL30,
			MinVersion: tls.VersionTLS10,
			MaxVersion: tls.VersionTLS10,
			// PreferServerCipherSuites: true,
		},
	}
	server.ListenAndServeTLS(publicKey, privateKey)

	log.Infof("Server stopped on port: %v", port)

}

func OnSignin(c *gin.Context) {

	reqBody := &struct {
		UserID   string `json:"userID"`
		Password string `json:"password"`
	}{}

	if err := c.Bind(reqBody); err != nil {
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": false,
		"token":   "",
		"message": "userID & password doesn't match",
	})

}
