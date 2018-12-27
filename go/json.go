package main

import (
	"encoding/json"
	"fmt"
)
/*
tag标签
 */
type Product struct {
	Name string `json:"name"`
	ProductID int64 `json:"product_id,string""`
	Number int `json:"number"`
	Price float64 `json:"price"`
	IsOnSale bool `json:"onsale"`
}
/*
json.Marshal()/UnMarshal()
 */

 func main(){
 	p := Product{
 		"xiaomi 6",
 		10001,
 		10000,
 		2000.0,
 		true,
	}
 	data,err := json.Marshal(p)
 	fmt.Println(string(data),err)
 	r := Product{}
 	json.Unmarshal(data,&r)
 	fmt.Printf("%+v\n",r)
 }