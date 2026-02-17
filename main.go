package main

import (
    "net/http"
    "log"
)

func main() {
    // Статические файлы
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    // Роуты
    http.HandleFunc("/", indexHandler)
    http.HandleFunc("/lesson/", lessonHandler)
    http.HandleFunc("/login", loginHandler)
    http.HandleFunc("/dashboard", dashboardHandler)

    log.Println("Сервер запущен на http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
    // Отображение списка уроков
}
