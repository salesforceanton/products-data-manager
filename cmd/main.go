package main

import (
	"fmt"

	"github.com/salesforceanton/products-data-manager/internal/config"
	"github.com/salesforceanton/products-data-manager/internal/logger"
)

func main() {
	// Intialize App configuration
	cfg, err := config.InitConfig()
	if err != nil {
		logger.LogServerIssue(err)
		return
	}

	fmt.Println(cfg)
}
