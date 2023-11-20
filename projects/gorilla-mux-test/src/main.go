package main 

import (
	"fmt"
	utils "github.com/GabrielMessiasdaRosa/gorilla-mux-test/src/utils"
)

func main() {
	fmt.Println("Hello, World!")
	uuid:= utils.GenerateUUID()
	fmt.Println(uuid)
	utils.Public()
	utils.UsePrivate()
}
