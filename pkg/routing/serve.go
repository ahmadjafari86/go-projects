package routing

import (
	"blog/pkg/config"
	"fmt"
	"log"
)

func Serve() {
	router := GetRouter()
	configs := config.Get()

	err := router.Run(fmt.Sprintf("%s:%s", configs.Server.Host, configs.Server.Port))

	if err != nil {
		log.Fatalf("Error in routing!")
		return
	}
}
