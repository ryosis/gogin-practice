package service

import (
	"fmt"
	"gogin-practice/config"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gopkg.in/yaml.v3"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*type cfg struct {
	Db *sql.DB
}*/

type cfg struct {
	Db *gorm.DB
}

var CFG cfg

func Init(BasePath string) {
	fmt.Println("Load from yml file")
	configPath := ""
	configPath = BasePath + "/config/file/config.yml"
	// fmt.Println("configPath:", configPath)
	data, err := os.ReadFile(configPath)
	// data, err := os.ReadFile("D:\\code\\workspace_go\\gogin-practice\\config\\file\\config.yml")

	if err != nil {
		log.Fatal("Error opening yaml file:", err)
	}

	var databaseConfig config.DatabaseConfig
	err = yaml.Unmarshal(data, &databaseConfig)
	if err != nil {
		log.Fatal("Error when unmarshaling yaml file:", err)
	}

	databaseConnection(databaseConfig)

	// fmt.Println("Host:", databaseConfig.Dbserver.Host)
	// fmt.Println("Port:", databaseConfig.Dbserver.Port)
	// fmt.Println("Username:", databaseConfig.Dbserver.Username)
	// fmt.Println("Password:", databaseConfig.Dbserver.Password)
	// fmt.Println("Database:", databaseConfig.Dbserver.Database)

	/*data, err := os.Open("D:\\code\\workspace_go\\gogin-practice\\config\\file\\config.yml")
	if err != nil {
		log.Fatal("Error opening yaml file:", err)
	}
	defer data.Close()

	var databaseConfig config.DatabaseConfig
	decoder := yaml.NewDecoder(data)
	err = decoder.Decode(&databaseConfig)

	if err != nil {
		panic(err)
	}

	fmt.Println("Name:", databaseConfig.Dbserver.Username)
	fmt.Println("Pwd:", databaseConfig.Dbserver.Password)*/
}

func InitEnv() {
	fmt.Println("Load from .env")
	godotenv.Load()
	fmt.Println("DB_HOST = " + os.Getenv("DB_HOST"))
	fmt.Println("DB_PORT = " + os.Getenv("DB_PORT"))
	fmt.Println("DB_NAME = " + os.Getenv("DB_NAME"))
	fmt.Println("DB_USERNAME = " + os.Getenv("DB_USERNAME"))
	fmt.Println("DB_PASSWORD = " + os.Getenv("DB_PASSWORD"))
}

func databaseConnection(databaseConfig config.DatabaseConfig) {
	db, err := gorm.Open(mysql.Open(databaseConfig.Dbserver.Username + ":" + databaseConfig.Dbserver.Password + "@tcp(" + databaseConfig.Dbserver.Host + ":" + databaseConfig.Dbserver.Port + ")/" + databaseConfig.Dbserver.Database))

	if err != nil {
		// control error
	}

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Error Database Connection", err)
	}

	sqlDB.SetMaxOpenConns(140)
	sqlDB.SetMaxIdleConns(140)
	sqlDB.SetConnMaxIdleTime(5 * 60)
	sqlDB.SetConnMaxLifetime(5 * 60)

	//defer sqlDB.Close()

	CFG.Db = db

	/*db, err := sql.Open("mysql", databaseConfig.Dbserver.Username+":"+databaseConfig.Dbserver.Password+"@tcp("+databaseConfig.Dbserver.Host+":"+databaseConfig.Dbserver.Port+")/"+databaseConfig.Dbserver.Database)

	if err != nil {
		log.Fatal("Error Database Connection", err)
	}

	db.SetMaxOpenConns(140)
	db.SetMaxIdleConns(140)
	db.SetConnMaxIdleTime(5 * 60)
	db.SetConnMaxLifetime(5 * 60)

	defer db.Close()

	CFG.Db = db*/

	fmt.Println("Database connected.............[OK]")
}
