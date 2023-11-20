package main 

import (
	"fmt"
	utils "github.com/GabrielMessiasdaRosa/gorilla-mux-test/src/utils"
)

func main() {
	fmt.Println("Hello, World!")
	uuid, crawlerResult := utils.GenerateUUID(), utils.TestingBlankIdentifier()
	fmt.Println(uuid)
	utils.Public()
	utils.UsePrivate()
	fmt.Println(crawlerResult)
}
