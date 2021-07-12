package V1

import(
	"github.com/gin-gonic/gin"
	"Tes/expandana_shopping_cart/models"
	"fmt"
	"strconv"
	// "net/http"
)

type V1DeleteItemsController struct {
	Status int
}


func (status *V1DeleteItemsController) DeleteItemsController (c *gin.Context){
	IdUser, _ := strconv.ParseInt(c.PostForm("params[Id_user]"),0,64)
	IdItem, _ := strconv.ParseInt(c.PostForm("params[Id_item]"),0,64)
	paramIdUser := models.ParamIdUser{
		IdUser:IdUser,
	}
	paramChartItems := models.Item{}
    ChartItems := []models.Item{}
	getItems := models.GetItem(paramIdUser)
	for i :=0; i < len(getItems); i++ {
		if getItems[i].IdItem != IdItem{
			paramChartItems.IdItem = getItems[i].IdItem
			paramChartItems.TotalItem = getItems[i].TotalItem
			paramChartItems.NameItem = getItems[i].NameItem
			ChartItems = append(ChartItems,paramChartItems)
		}
	}
	parasmItems := models.ParamsItem{
		IdUser:IdUser,
		Items:ChartItems,
	}
	response := models.AddNewCartItems(parasmItems)
	fmt.Println(response)
	c.JSON(200, gin.H{"status": 200, "response":response})
	return
}