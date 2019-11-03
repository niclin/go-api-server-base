package conf

import (
	"os"

	"go101/model"

	"github.com/joho/godotenv"
)

func Init() {
	godotenv.Load()

	model.Database(os.Getenv("MYSQL_DSN"))
}
