package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	server.GET("/", func(c *gin.Context) {
		c.Header("content-type", "text/html")
		c.File("html/index.html")
	})
	server.GET("/install.plist", func(c *gin.Context) {
		c.File("plist/install.plist")
	})
	server.GET("/LINE++.ipa", func(c *gin.Context) {
		c.File("ipa_files/LINE++.ipa")
	})
	server.RunTLS(":44111", "keys/cert.pem", "keys/key.pem")
}
