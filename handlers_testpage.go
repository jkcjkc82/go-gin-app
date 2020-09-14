package main

import (
	"github.com/gin-gonic/gin"
)

func showTestpage(c *gin.Context) {
	render(c, gin.H{
		"title":   "test page",
		"payload": "test page payload"}, "testpage.html")

}
