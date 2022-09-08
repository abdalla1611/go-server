package main

import 
(
	"github.com/gin-gonic/gin"
	"net/http"
)

type album struct{
	ID string  `json:"id"`
	Title string `json:"title"`
	Artist string `json:"artist"`
	Price float64 `json:"price"`
}
func getAlbums(c *gin.Context){
	c.IndentedJSON(http.StatusOK , albums)
}
func addAlbum(c *gin.Context){
	var newAlbum album
	if err:=c.BindJSON(&newAlbum); err != nil {
		// we have a problem
		return
	}
	albums = append(albums, newAlbum)
	c.IndentedJSON(http.StatusCreated ,newAlbum)
}
func getAlbumById(c *gin.Context){
	id := c.Param("id")
	for _ , i := range albums{
		if i.ID == id{
			c.IndentedJSON(http.StatusOK , i)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound , gin.H{"message":"album not found"})

}

// albums slice to seed record album data.
var albums = []album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func main(){
	router := gin.Default()
	router.GET("/albums" , getAlbums)
	router.GET("/albums/:id" , getAlbumById)
	router.POST("/albums" , addAlbum)
	router.Run("localhost:8080")
}

