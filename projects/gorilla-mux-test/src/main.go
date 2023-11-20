package main 

import (
	"fmt"
	utils "github.com/GabrielMessiasdaRosa/gorilla-mux-test/src/utils"
)

func main() {
	fmt.Println("Hello, World!")
	uuid, deuBoa := utils.GenerateUUID(), utils.DeuBoa()
	fmt.Println(uuid)
	fmt.Println(deuBoa)
}
