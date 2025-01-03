package api

import (
	"github.com/alpgozbasi/dms-ocr/internal/api/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

func SetupRouter(db *sqlx.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	docHandler := handlers.NewDocumentHandler(db)

	router.POST("/documents", docHandler.CreateDocument)
	router.GET("/documents", docHandler.ListDocuments)

	router.POST("/documents/upload", docHandler.UploadFile)
	
	return router
}
