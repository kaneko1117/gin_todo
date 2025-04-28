package main

import (
	"fmt"
	"gin_todo/internal/entity/model"
	"gin_todo/internal/entity/query"
	"gin_todo/internal/infra/database"
	"os"
	"time"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	db := database.NewDB()
	defer fmt.Println("Successfully seeded")
	defer database.CloseDB(db)

	query := query.Use(db)
	count, err := query.User.Count()
	if err != nil {
		fmt.Println("Error seeding data:", err)
		return
	}
	if count > 0 {
		fmt.Println("Data already exists, skipping seeding.")
		return
	}

	pw, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SEED_USER_PASSWORD")), 10)
	if err != nil {
		fmt.Println(err)
	}

	userData := model.User{
		UserName:  "admin",
		Password:  string(pw),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	err = query.User.Create(&userData)
	if err != nil {
		fmt.Println("Error seeding data:", err)
		return
	}

	taskData := model.Task{
		UserID:    userData.ID,
		Tasks:     "洗濯物を干す",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		IsChecked: false,
	}

	err = query.Task.Create(&taskData)
	if err != nil {
		fmt.Println("Error seeding data:", err)
		return
	}
	fmt.Println("Seeding completed successfully.")

}
