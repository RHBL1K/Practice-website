package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/smtp"
	"time"


	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


const (
	SMTPServer      = "smtp.gmail.com"
	SMTPPort        = "587"
	SenderEmail     = ""
		// Нужно написать почту отправителя
	SenderPassword  = ""
		// Нужно написать пароль от почты отправителя
	ReceiverEmail   = ""
		// Нужно написать почту получателя
	MongoDBURI      = "mongodb+srv://Aldiyar:Nursultan2005@cluster0.8jzkf.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	DatabaseName    = "contactFormDB"
	CollectionName  = "messages"
)


type FormData struct {
	Name    string `bson:"name"`
	Phone   string `bson:"phone"`
	Message string `bson:"message"`
	Time    string `bson:"time"`
}


func connectMongoDB() (*mongo.Client, error) {
	clientOptions := options.Client().ApplyURI(MongoDBURI)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	fmt.Println(" Успешное подключение к MongoDB")
	return client, nil
}

func main() {
	
	client, err := connectMongoDB()
	if err != nil {
		log.Fatal("Ошибка подключения к MongoDB:", err)
	}
	defer client.Disconnect(context.TODO())


	collection := client.Database(DatabaseName).Collection(CollectionName)

	// Обслуживание статики
	fs := http.FileServer(http.Dir("static"))
	http.Handle("/", fs)

	// Обработчик формы
	http.HandleFunc("/submit", func(w http.ResponseWriter, r *http.Request) {
		handleFormSubmission(w, r, collection)
	})

	fmt.Println("🌍 Сервер запущен на http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func handleFormSubmission(w http.ResponseWriter, r *http.Request, collection *mongo.Collection) {
	if r.Method != http.MethodPost {
		http.Error(w, "Метод запрещен", http.StatusMethodNotAllowed)
		return
	}

	// Получаем данные из формы
	name := r.FormValue("name")
	phone := r.FormValue("phone")
	message := r.FormValue("message")
	currentTime := time.Now().Format("2006-01-02 15:04:05")

	// Сохраняем в MongoDB
	formData := FormData{
		Name:    name,
		Phone:   phone,
		Message: message,
		Time:    currentTime,
	}

	_, err := collection.InsertOne(context.TODO(), formData)
	if err != nil {
		log.Println("Ошибка сохранения в MongoDB:", err)
		http.Error(w, "Ошибка при сохранении данных", http.StatusInternalServerError)
		return
	}

	fmt.Println(" Данные успешно сохранены в MongoDB")

	// Отправляем письмо
	emailBody := fmt.Sprintf("Имя: %s\nТелефон: %s\nСообщение:\n%s\nВремя: %s", name, phone, message, currentTime)
	auth := smtp.PlainAuth("", SenderEmail, SenderPassword, SMTPServer)

	err = smtp.SendMail(SMTPServer+":"+SMTPPort, auth, SenderEmail, []string{ReceiverEmail}, []byte("Subject: Новое сообщение с сайта\n\n"+emailBody))
	if err != nil {
		log.Println("Ошибка отправки почты:", err)
		http.Error(w, "Ошибка при отправке сообщения", http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, "Сообщение успешно отправлено и сохранено в базе!")
}
