package main

import (
	"ginchat/router"
	"ginchat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()

	r := router.Router()
	err := r.Run(":8081")
	if err != nil {
		return
	}
}
