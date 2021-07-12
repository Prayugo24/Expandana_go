package models

import (
	"Tes/expandana_shopping_cart/config"
	"encoding/json"
	"fmt"
	"strconv"
)
type ParamsItem struct{
	IdUser int64
	Items []Item
}
type ParamIdUser struct {
	IdUser int64
}
type ResponseRemoveKeys struct{
	Status int
	Message string
}
type Item struct {
	IdItem int64 `json:"id_item"`
	TotalItem int64 `json:"total_item"`
	NameItem string `json:"name_item"`
}


func AddNewCartItems(params ParamsItem) (respon []Item){
	clientRedis := config.SetupRedis()
	idUser := strconv.FormatInt(int64(params.IdUser), 10)
	fmt.Println(idUser)
	json, err := json.Marshal(params.Items)
    if err != nil {
        fmt.Println(err)
    }
    err = clientRedis.Set(clientRedis.Context(),idUser, json, 0).Err()
    if err != nil {
        fmt.Println(err)
    }
	paramIdUser :=ParamIdUser{
		IdUser:params.IdUser,
	}
    respon = GetItem(paramIdUser)
	return respon
}

func GetItem(params ParamIdUser)(response []Item){
	clientRedis := config.SetupRedis()
	idUser := strconv.FormatInt(int64(params.IdUser), 10)
	val, err := clientRedis.Get(clientRedis.Context(),idUser).Result()
    if err != nil {
        fmt.Println(err)
    }
	
	response = []Item{}
	json.Unmarshal([]byte(val), &response)
	
	return  response
}

func RemoveKeys(params ParamIdUser)(response ResponseRemoveKeys){
	clientRedis := config.SetupRedis()
	idUser := strconv.FormatInt(int64(params.IdUser), 10)
	err := clientRedis.Del(clientRedis.Context(), idUser).Err()
	if err != nil {
		response = ResponseRemoveKeys{
			Status : 400,
			Message : "Delete Failed",
		}
	}else{
		response = ResponseRemoveKeys{
			Status : 200,
			Message : "Delete Succes",
		}
	}
	return response
}
