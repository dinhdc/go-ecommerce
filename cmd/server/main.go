package main

import (
	"github.com/dinhdc/go-ecommerce/internal/routers"
)

func main() {
	r := routers.NewRouter()
	r.Run(":8002")
}
