package main

import (
	"github.com/zaahidali/task_manager_api/router"
)

func main() {
	r := router.SetupRouter()

	r.Run()
}
