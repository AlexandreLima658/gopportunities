package main

import (
	"github.com/AlexandreLima658/gopportunities/config"
	"github.com/AlexandreLima658/gopportunities/router"
)

var (
	logger *config.Logger
)

func main() {

	logger = config.GetLogger("main")

	err := config.Init()

	if err != nil {
		logger.Errorf("Config initialization error: %v", err)
		return
	}

	router.Init()

}
