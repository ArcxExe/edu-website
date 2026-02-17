package handlers

import (
    "net/http"
    "edu-website/db"
)

func loginHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method == "POST" {
        username := r.FormValue("username")
        password := r.FormValue("password")
        
        // Проверка пароля (хэширование можно добавить через bcrypt)
        var storedPassword string
        err := db.DB.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&storedPassword)
        if err != nil || password != storedPassword {
            http.Error(w, "Неверный логин или пароль", http.StatusUnauthorized)
            return
        }
        
        // Сохранение сессии (можно использовать куки или JWT)
        http.Redirect(w, r, "/dashboard", http.StatusSeeOther)
    } else {
        // Отображение формы входа
        tmpl, _ := template.ParseFiles("templates/login.html")
        tmpl.Execute(w, nil)
    }
}
