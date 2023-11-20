package main 

import (
	"fmt"
	utils "github.com/GabrielMessiasdaRosa/gorilla-mux-test/src/utils"
)

func main() {
	fmt.Println("Hello, World!")
	uuid, crawlerResult := utils.GenerateUUID(), utils.TestingBlankIdentifier()
	constantTest := utils.TestingConstants()
	strings := utils.TestingMultipleContants()
	string1, string2, string3  := strings[0], strings[1], strings[2]
	fmt.Println(uuid)
	utils.Public()
	utils.UsePrivate()
	fmt.Println(crawlerResult)
	fmt.Println(constantTest)
	fmt.Println(string1)
	fmt.Println(string2)
	fmt.Println(string3)
}
