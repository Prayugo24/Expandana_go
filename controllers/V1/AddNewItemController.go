package V1

import(
	"github.com/gin-gonic/gin"
	"Tes/expandana_shopping_cart/models"
	"fmt"
	"strconv"
)

type V1AddNewItemController struct {
	Status int
}


func (status *V1AddNewItemController) AddNewItemController (c *gin.Context){
	IdUser, _ := strconv.ParseInt(c.PostForm("params[Id_user]"),0,64)
	IdItem, _ := strconv.ParseInt(c.PostForm("params[Id_item]"),0,64)
	TotalItem, _ := strconv.ParseInt(c.PostForm("params[Total_item]"),0,64)
	NameItem := c.DefaultPostForm("params[Name_item]", "")
	paramIdUser := models.ParamIdUser{
		IdUser:IdUser,
	}
	paramChartItems := models.Item{}
    ChartItems := []models.Item{}
	getItems := models.GetItem(paramIdUser)
	addNew := true
	for i :=0; i < len(getItems); i++ {
		if getItems[i].IdItem == IdItem{
			addNew = false
			paramChartItems.IdItem = getItems[i].IdItem
			paramChartItems.TotalItem = getItems[i].TotalItem+TotalItem
			paramChartItems.NameItem = getItems[i].NameItem
			ChartItems = append(ChartItems,paramChartItems)
		}else{
			paramChartItems.IdItem = getItems[i].IdItem
			paramChartItems.TotalItem = getItems[i].TotalItem
			paramChartItems.NameItem = getItems[i].NameItem
			ChartItems = append(ChartItems,paramChartItems)
		}
	}
	if addNew {
		paramChartItems.IdItem = IdItem
		paramChartItems.TotalItem = TotalItem
		paramChartItems.NameItem = NameItem
		ChartItems = append(ChartItems,paramChartItems)
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