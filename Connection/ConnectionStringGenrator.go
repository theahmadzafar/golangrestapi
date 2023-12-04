package conns

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func GetConString() string {
	err := godotenv.Load()
	if err != nil {
		return fmt.Sprintf("Error loading .env file: %s", err)
	}
	Host := os.Getenv("HOST")
	Database := os.Getenv("DATABASE")
	DBPort := os.Getenv("DBPORT")
	User := os.Getenv("USER")
	Password := os.Getenv("PASSWORD")
	constring := fmt.Sprintf("host=%s user=%s dbname=%s  sslmode=disable password=%s  port=%s", Host, User, Database, Password, DBPort)

	return constring
}
