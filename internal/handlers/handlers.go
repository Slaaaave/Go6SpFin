package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

func HandlerIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	http.ServeFile(w, r, "index.html")
}

func HandlerUpload(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Ошибка 405", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Не удалось получить файл с формы", http.StatusBadRequest)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Ошибка при чтении файла", http.StatusInternalServerError)
		return
	}

	converted, err := service.IsMorse(string(data))
	if err != nil {
		http.Error(w, "Ошибка в конвертации морзе: "+err.Error(), http.StatusUnprocessableEntity)
		return
	}

	ext := filepath.Ext(header.Filename)
	currentTime := time.Now().Format("2006-01-02_15-04-05")
	fileName := fmt.Sprintf("%s%s", currentTime, ext)

	dst, err := os.Create(fileName)
	if err != nil {
		http.Error(w, "Не удалось создать файл", http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	_, err = dst.WriteString(converted)
	if err != nil {
		http.Error(w, "Ошибка при записи в файл", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(converted))
}
