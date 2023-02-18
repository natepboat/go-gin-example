package service

import (
	"example/gin-example/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

var albums = model.AlbumRoot{}

func init() {
	albums = model.AlbumRoot{
		Albums: []model.Album{
			{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
			{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
			{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
		},
	}
}

func ListAlbum(c *gin.Context) {
	var reponseType = c.Query("type")
	if reponseType == "xml" {
		c.XML(http.StatusOK, albums)
	} else {
		c.IndentedJSON(http.StatusOK, albums)
	}
}
