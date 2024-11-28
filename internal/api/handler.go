package api 

import (
	"fmt"
	"net/http"
	"log"
	"os"
	"path/filepath"

	"github.com/naluneotlichno/rvc/internal/player"
)

// Директория для загрузки файлов
const uploadDir = "./uploads"

// Старт HTTP сервера
func StartServer() {

	// Создаем папку для загрузки файлов, если ее нет
	if err := os.Mkdir(uploadDir, os.ModePerm); err != nil {
		log.Fatalf("Не удалось создать папку для загрузки файлов: %v", err)
	}

	// Интерфейс CLI
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprint(w, `
			<h1>Добро пожаловать в RVC!</h1>
			<p>Введите путь к музыкальному файлу и нажмите "Играть":</p>
			<form action="/play" method="post">
				<input type="text" name="file" placeholder="Введите путь к файлу" style="width:300px;">
				<button type="submit">Играть</button>
			</form>
		`)
	})

	// Добавляем эндпоинт для проигрывания музыки
	http.HandleFunc("/play", func(w http.ResponseWriter, r *http.Request) {

		// Проверка метода POST. Если не он -- 405 error
		if r.Method != http.MethodPost {
			http.Error(w, 
				"Метод не поддерживается, используйте POST", 
				http.StatusMethodNotAllowed)
			return
		}

		// Загружаем файл
		file, _, err := r.FormFile("file")
		if err != nil {
			http.Error(w, fmt.Sprintf(
				"Не удалось загрузить файл: %v", err), 
				http.StatusInternalServerError
			)
			return
		}
		defer file.Close()

		//  Создаем файл на сервере
		dst, err := os.Create(
			filepath.Join(uploadDir, "uploaded_music"
			))
		if err != nil {
			http.Error(w, 
				fmt.Sprintf("Ошибка при сохранении файла: %v",err),
				http.StatusInternalServerError
			)
			return
		}	
		defer dst.Close()


		file := r.URL.Query().Get("file")

		// Если file не указан -- 400 error
		if file == "" {
			http.Error(w, 
				`Параметр 'file' не может быть пустым.
				 Укажите путь к файлу.`, 
				 http.StatusBadRequest)
			return
		}

		// Проигрывание файла
		player.PlayMusic(file)

		// Отправляем ответ пользователя
		w.Header().Set("Content-Type", "text/html; charset=utf-8")
		fmt.Fprintf(w, `
			<h1>Файл проигрывается: %s</h1>
			<a href="/">Назад</a>
		`, file)
	})

	// Запускаем сервер. 
	log.Fatal(http.ListenAndServe(":8080", nil))
}
