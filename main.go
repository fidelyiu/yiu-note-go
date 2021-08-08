package main

import (
	"github.com/fidelyiu/yiu-note/core/router"
)

func main() {
	r := router.InitRouter()
	_ = r.Run(":8080")
}
