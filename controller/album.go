package controller

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"

	"restapi/service/albumService"
)

type album struct {
	ID     string `valid:"MaxSize(50)"`
	Title  string `valid:"MaxSize(50)"`
	Artist string `valid:"MaxSize(50)"`
	Price  float64
}

func GetAlbums(c *gin.Context) {
	albums, _ := albumService.GetAlbums()
	c.IndentedJSON(http.StatusOK, albums)
}

func GetAlbum(c *gin.Context) {
	valid := validation.Validation{}
	id := c.Param("id")

	a := album{ID: id}

	ok, _ := valid.Valid(&a)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album id"})
		return
	}

	albService := albumService.Album{ID: id}
	album, err := albService.GetAlbum()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Provided id deon't exists"})
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func PostAlbum(c *gin.Context) {
	valid := validation.Validation{}
	var newAlbum album

	if err := c.BindJSON(&newAlbum); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Some error while parsing the data"})
		return
	}

	ok, _ := valid.Valid(&newAlbum)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album data"})
		return
	}

	albService := albumService.Album{ID: newAlbum.ID, Title: newAlbum.Title, Artist: newAlbum.Artist, Price: newAlbum.Price}
	album, err := albService.PostAlbum()

	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Provided id already exists"})
		return
	}

	c.IndentedJSON(http.StatusCreated, album)
}

func UpdateAlbum(c *gin.Context) {
	valid := validation.Validation{}
	var updateAlbumData album

	if err := c.BindJSON(&updateAlbumData); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "Some error while parsing the data"})
		return
	}

	id := c.Param("id")

	a := album{ID: id, Title: updateAlbumData.Title, Artist: updateAlbumData.Artist, Price: updateAlbumData.Price}

	ok, _ := valid.Valid(&a)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album data"})
		return
	}

	albService := albumService.Album{ID: id, Title: updateAlbumData.Title, Artist: updateAlbumData.Artist, Price: updateAlbumData.Price}
	album, err := albService.UpdateAlbum()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, album)
}

func DeleteAlbum(c *gin.Context) {
	valid := validation.Validation{}
	id := c.Param("id")

	a := album{ID: id}

	ok, _ := valid.Valid(&a)

	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid album id"})
		return
	}

	albService := albumService.Album{ID: id}
	albums, err := albService.DeleteAlbum()

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, err)
		return
	}

	c.IndentedJSON(http.StatusOK, albums)
}
