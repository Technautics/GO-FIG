package app

import (
	"fmt"

	"github.com/Technautics/GO-FIG/config"
)

// The application logic lives here !

func StartApp(cfg *config.Config) {
	fmt.Println("🔧 Loaded Configuration :")
	fmt.Println("----------------")
	fmt.Println("App_Name", cfg.App_Name)
	fmt.Println("Environment", cfg.Env)
	fmt.Println("Port", cfg.Port)
	fmt.Println("Debug", cfg.Debug)
	fmt.Println("Timeout_Seconds", cfg.Timeout_seconds)
}
