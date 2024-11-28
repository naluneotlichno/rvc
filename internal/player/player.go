package player

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Команда для воспроезвидения музыки
var PlayCommand = &cobra.Command{
	Use:	"play [file]",
	Short:	"Проигрывает музыкальный файл",
	Args:	cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		filePath := args[0]
		PlayMusic(filePath)
	},
}

// Воспроизведение музыки
func PlayMusic(filePath string) {
	fmt.Printf("Проигрываю музыкальный файл: %s\n", filePath)
	// Тут потом добавим реальную логику работы с аудиофайлом
}