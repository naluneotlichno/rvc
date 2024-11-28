package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/naluneotlichno/rvc/internal/player"
	"github.com/naluneotlichno/rvc/internal/api"
	"github.com/spf13/cobra"
)

func main() {

	// Запускаем сервер в горутине
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("Запуск RVC сервера на http://localhost:8080")
		api.StartServer()
	}()

	// Главная команда для запуска CLI
	rootCmd := &cobra.Command{
		Use:    "rvc",
		Short:	"CLI это утонченно.",
		Long: 	`RVC - музыкальный сервер. Он предоставляет REST API для управления музыкой`,
	}

	// Подключаем команду Play
	rootCmd.AddCommand(player.PlayCommand)

	// Запуск CLI
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	// Ожидание завершения работы сервера
	wg.Wait()
}