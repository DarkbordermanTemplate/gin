package endpoints

import (
	"log"
	"practice/connection"

	"github.com/gin-gonic/gin"
)

type Fruit struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

// @Summary Get fruit count by url queries
// @produce application/json
// @param name query string true "Fruit name"
// @Success 200 {object} Fruit
// @Router /fruit [get]
func Get_by_query(context *gin.Context) {
	value, ok := connection.MEMORY[context.Query("name")]
	if !ok {
		context.String(400, "Bad Request")
		return
	}
	context.JSON(200, gin.H{
		"name":  context.Query("name"),
		"count": value,
	})
}

// @Summary Get fruit count by url paths
// @produce application/json
// @param name path string true "Fruit name"
// @Success 200 {object} Fruit
// @Router /fruit/{name} [get]
func Get_by_param(context *gin.Context) {
	value, ok := connection.MEMORY[context.Param("name")]
	if !ok {
		context.String(400, "Bad Request")
		return
	}
	context.JSON(200, gin.H{
		"name":  context.Param("name"),
		"count": value,
	})
}

// @Summary Post fruit count by JSON
// @produce text/plain
// @Param name body Fruit true "ok"
// @Success 200 {object} string "OK"
// @Router /fruit [post]
func Post_by_JSON(context *gin.Context) {
	json := Fruit{}
	context.BindJSON(&json)
	_, ok := connection.MEMORY[json.Name]
	log.Print(ok)
	if ok {
		context.String(400, "Bad Request")
		return
	}
	connection.MEMORY[json.Name] = json.Count
	context.String(200, "OK")
}
