package handlers

import (
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/service"
)

// хэндлер "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	i, err := os.ReadFile("../index.html")
	if err != nil {
		http.Error(w, "File reading error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(i)

}

// хэндлер "/upload"
func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Invalid parse", http.StatusInternalServerError)
		return
	}

	file, handler, err := r.FormFile("myFile")
	if err != nil {
		http.Error(w, "Unable to get file", http.StatusInternalServerError)
		return
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		http.Error(w, "Error read file", http.StatusInternalServerError)
		return
	}

	convertation, err := service.AutoConvertation(string(data))
	if err != nil {
		http.Error(w, "Error convertation file", http.StatusInternalServerError)
		return
	}

	ext := filepath.Ext(handler.Filename)
	name := time.Now().UTC().String()

	// windows не давал сохранить файл из-за недопустимых символов в имени файла
	name = strings.ReplaceAll(name, ":", "-")
	name = strings.ReplaceAll(name, " ", "_")
	opName := name + ext

	err = os.WriteFile(opName, []byte(convertation), 0644)
	if err != nil {
		http.Error(w, "Error save file", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.Write([]byte(convertation))

}
