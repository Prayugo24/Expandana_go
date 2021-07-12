package routers

import (
	"github.com/gin-gonic/gin"
	"Tes/expandana_shopping_cart/config"
	"net/http"
	"fmt"
	"Tes/expandana_shopping_cart/controllers/V1"
)

func RouterMain() http.Handler {
	router := gin.New()
	clientRedis := config.SetupRedis()
	pong, err := clientRedis.Ping(clientRedis.Context()).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("Redis Message :",pong)
	v1AddNewItem := &V1.V1AddNewItemController{Status: 200}
	v1DeleteItem := &V1.V1DeleteItemsController{Status: 200}
	v1DeleteChart := &V1.V1DeleteChartController{Status: 200}
	
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"data":"Welcome To Api Golang"})
	})

	router.POST("/addItems", v1AddNewItem.AddNewItemController)
	router.POST("/deleteItems", v1DeleteItem.DeleteItemsController)
	router.POST("/deleteChart", v1DeleteChart.DeleteChartController)
	
	return router
}
