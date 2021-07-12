package V1

import(
	"github.com/gin-gonic/gin"
	"Tes/expandana_shopping_cart/models"
	"fmt"
	"strconv"
)

type V1DeleteChartController struct {
	Status int
}


func (status *V1DeleteChartController) DeleteChartController (c *gin.Context){
	IdUser, _ := strconv.ParseInt(c.PostForm("params[Id_user]"),0,64)
	paramIdUser := models.ParamIdUser{
		IdUser:IdUser,
	}
	response := models.RemoveKeys(paramIdUser)
	fmt.Println(response)
	c.JSON(200, gin.H{"status": 200, "response":response})
	return
}

