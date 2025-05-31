package handlers

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

// GetFiles handles the file listing endpoint
func GetFiles(c *gin.Context) {
	// Basic file listing - you might want to implement proper file browsing
	path := c.DefaultQuery("path", "/")

	files, err := os.ReadDir(path)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var fileList []map[string]interface{}
	for _, file := range files {
		info, _ := file.Info()
		fileList = append(fileList, map[string]interface{}{
			"name":     file.Name(),
			"is_dir":   file.IsDir(),
			"size":     info.Size(),
			"modified": info.ModTime(),
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"path":  path,
		"files": fileList,
	})
}
