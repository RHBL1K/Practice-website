package main

import (
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func sendEmail(name, phone, message string) error {
	from := "aldikoshka@gmail.com"  // Укажите ваш email
	password := "" // Укажите пароль от почты

	to := []string{"aldiyarkuandyk68@gmail.com"} // Один получатель
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"

	auth := smtp.PlainAuth("", from, password, smtpHost)

	body := fmt.Sprintf("Name: %s\nPhone: %s\nMessage: %s", name, phone, message)

	msg := []byte("Subject: Contact Form Submission\n" +
		"MIME-Version: 1.0\n" +
		"Content-Type: text/plain; charset=UTF-8\n\n" +
		body)

	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, from, to, msg)
	if err != nil {
		return err
	}
	return nil
}

func formHandler(w http.ResponseWriter, r *http.Request) {
	
    w.Header().Set("Access-Control-Allow-Methods", "POST")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:5500")

    if r.Method == http.MethodOptions { // Предварительный CORS-запрос
        w.WriteHeader(http.StatusOK)
        return
    }

    if r.Method != http.MethodPost {
        http.Error(w, "Метод не поддерживается", http.StatusMethodNotAllowed)
        return
    }

    err := r.ParseForm()
    if err != nil {
        http.Error(w, "Ошибка обработки формы", http.StatusBadRequest)
        return
    }

    name := r.FormValue("name")
    phone := r.FormValue("phone")
    message := r.FormValue("message")

    if name == "" || phone == "" || message == "" {
        http.Error(w, "Все поля должны быть заполнены", http.StatusBadRequest)
        return
    }

    if err := sendEmail(name, phone, message); err != nil {
        log.Println("Ошибка отправки email:", err)
        http.Error(w, "Ошибка отправки email", http.StatusInternalServerError)
        return
    }

    w.Write([]byte("Сообщение успешно отправлено!"))
}


func main() {
	http.HandleFunc("/submit", formHandler)
	port := os.Getenv("PORT")
	if port == "" {
		port = "5500"
	}
	
	fmt.Println("Сервер запущен на порту:", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
