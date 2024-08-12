package seeder

import (
	articleModel "blog/internal/modules/article/models"
	userModel "blog/internal/modules/user/models"
	"blog/pkg/database"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func Seed() {
	db := database.Connection()

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte("secret"), 12)
	if err != nil {
		log.Fatal("hash password error")
		return
	}

	user := userModel.User{Name: "random name", Email: "random@gmail.com", Password: string(hashedPassword)}
	db.Create(&user)
	log.Printf("User created successfully with email address %s", user.Email)

	for i := 1; i <= 10; i++ {
		article := articleModel.Article{Title: fmt.Sprintf("Random title %d", i), Content: fmt.Sprintf("Random content %d", i), UserID: 1}
		db.Create(&article)
		log.Printf("Article created successfully with title %s", article.Title)
	}

	log.Println("Seeder done.")
}
