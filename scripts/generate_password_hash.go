package main

import (
	"flag"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	// Пароль из аргумента командной строки
	password := flag.String("password", "admin123", "пароль для хэширования")
	flag.Parse()

	if *password == "" {
		log.Fatal("Пароль не может быть пустым. Используй -password=твой_пароль")
	}

	// Генерация хэша
	hash, err := bcrypt.GenerateFromPassword([]byte(*password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalf("Ошибка при генерации хэша: %v", err)
	}

	fmt.Println("Пароль:", *password)
	fmt.Println("Bcrypt-хэш:")
	fmt.Println(string(hash))
}