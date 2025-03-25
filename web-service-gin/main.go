package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func postAlbums(c *gin.Context) {
	var newAlbum album
	err := c.BindJSON(&newAlbum)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid Json input",
		})
		return
	}
	if newAlbum.ID == "" || newAlbum.Title == "" || newAlbum.Artist == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "All fields (id, title, artist) are required.",
		})
		return
	}

	albums = append(albums, newAlbum)
	c.JSON(http.StatusCreated, gin.H{
		"status": "success",
		"data":   newAlbum,
	})
}

func getAlbumId(c *gin.Context) {
	id := c.Param("id")
	for _, a := range albums {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

// func postAlbums(c *gin.Context) {
// 	var newAlbum album
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{
// 			"status":  "error",
// 			"message": "Invalid JSON input",
// 		})
// 		return
// 	}
// 	albums = append(albums, newAlbum)
// 	c.JSON(http.StatusCreated, gin.H{
// 		"status": "success",
// 		"data":   newAlbum,
// 	})
// }

func getAlbums(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   albums,
	})
}
func main() {
	router := gin.Default()
	router.GET("/albums/:id", getAlbumId)
	router.GET("/albums", getAlbums)
	router.POST("/albums", postAlbums)
	router.Run("localhost:8080")
}
