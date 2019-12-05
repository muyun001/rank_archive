package settings

import (
	"github.com/joho/godotenv"
	"log"
	"os"
	"strconv"
)

var (
	DbConnection string
	DbUsername   string
	DbPassword   string
	DbHost       string
	DbPort       string
	DbDatabase   string
)

var Debug bool

func init() {
	checkEnv()
	LoadSetting()
}

func checkEnv() {
	_ = godotenv.Load()
	needChecks := []string{
		"DB_CONNECTION", "DB_HOST", "DB_PORT", "DB_DATABASE", "DB_USERNAME", "DB_PASSWORD",
	}

	for _, envKey := range needChecks {
		if os.Getenv(envKey) == "" {
			log.Fatalf("env %s missed", envKey)
		}
	}
}

func LoadSetting() {
	DbConnection = os.Getenv("DB_CONNECTION")
	DbUsername = os.Getenv("DB_USERNAME")
	DbPassword = os.Getenv("DB_PASSWORD")
	DbHost = os.Getenv("DB_HOST")
	DbPort = os.Getenv("DB_PORT")
	DbDatabase = os.Getenv("DB_DATABASE")

	debug := os.Getenv("DEBUG")
	if debug != "" && debug != "false" && debug != "0" {
		Debug = true
	}
}

func loadIntFatal(e string) int {
	intVar, err := strconv.Atoi(os.Getenv(e))
	if err != nil {
		log.Fatalf("env %s invalid\n", e)
	}

	return intVar
}
