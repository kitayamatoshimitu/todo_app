package main

import (
	"fmt"
	controllers "todo_app/app/Controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}
