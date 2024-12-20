package main

import (
	"net/http"

	"go-to-do/commands"
	"go-to-do/configs"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
)

func main() {
	configs.LoadConfigs()
	cmd := &cobra.Command{}

	cmd.AddCommand(commands.DropTables())
	cmd.AddCommand(commands.Migrate())
	cmd.AddCommand(commands.Seed())
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}

	if configs.App.Env != "local" {
		MapURL()
	}
}

func MapURL() {
	router := gin.Default()
	// healthcheck path
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	err := router.Run()
	if err != nil {
		panic(err.Error() + "MapURL router not able to run")
	}
}
