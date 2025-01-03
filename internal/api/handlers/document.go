package handlers

import (
	"github.com/alpgozbasi/dms-ocr/internal/models"
	"github.com/alpgozbasi/dms-ocr/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"net/http"
	"time"
)

type DocumentHandler struct {
	DB *sqlx.DB
}

func NewDocumentHandler(db *sqlx.DB) *DocumentHandler {
	return &DocumentHandler{
		DB: db,
	}
}

func (h *DocumentHandler) CreateDocument(c *gin.Context) {
	var req struct {
		FileName string `json:"file_name"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	doc := models.Document{
		FileName:  req.FileName,
		FilePath:  "",
		OCRText:   "",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `
		INSERT INTO documents (file_name, file_path, ocr_text, created_at, updated_at)
		VALUES (:file_name, :file_path, :ocr_text, :created_at, :updated_at)
		RETURNING id
	`
	rows, err := h.DB.NamedQuery(query, doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create document"})
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&doc.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to scan document ID"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "Document created successfully",
		"document_id": doc.ID,
	})
}

func (h *DocumentHandler) ListDocuments(c *gin.Context) {
	var docs []models.Document

	err := h.DB.Select(&docs, "SELECT * FROM documents ORDER BY id DESC")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch documents"})
		return
	}

	c.JSON(http.StatusOK, docs)
}

func (h *DocumentHandler) UploadFile(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to upload file"})
		return
	}

	localPath, err := storage.SaveLocalFile(fileHeader) // generate a local path but not actually move the file yet
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to save file"})
	}

	err = c.SaveUploadedFile(fileHeader, localPath)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save file"})
		return
	}

	fileName := c.PostForm("file_name")
	if fileName == "" {
		fileName = fileHeader.Filename
	}

	doc := models.Document{
		FileName:  fileName,
		FilePath:  localPath,
		OCRText:   "", // will be filled after OCR
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	query := `
		INSERT INTO documents (file_name, file_path, ocr_text, created_at, updated_at)
		VALUES (:file_name, :file_path, :ocr_text, :created_at, :updated_at)
		RETURNING id
	`
	rows, err := h.DB.NamedQuery(query, doc)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to store document record"})
		return
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&doc.ID); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan new document ID"})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"message":     "File uploaded successfully",
		"document_id": doc.ID,
		"file_path":   doc.FilePath,
	})
}
