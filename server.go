package main

import (
	"fmt"
	"server/api"
)

func init(){
	fmt.Println("Server is getting started!")
}

func main(){

	defer func(){ fmt.Println("closing server") }()
	api.InitServer(":8080")

}
