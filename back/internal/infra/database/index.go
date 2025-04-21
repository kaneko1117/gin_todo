package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/postgres"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"github.com/joho/godotenv"
)

func NewDB() *gorm.DB {
	if os.Getenv("GO_ENV") == "dev" {
		err := godotenv.Load()
		if err != nil {
			log.Fatalln(err)
		}
	}


	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PW"), os.Getenv("POSTGRES_HOST"), os.Getenv("POSTGRES_PORT"), os.Getenv("POSTGRES_DB"))
	fmt.Println(dsn)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // 標準出力にログを出力
		logger.Config{
			SlowThreshold: time.Second, // 遅いクエリの閾値（1秒）
			LogLevel:      logger.Info, // 出力レベル（Info以上）
			Colorful:      true,        // カラフルな出力
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
		NowFunc: func() time.Time {
			// timezoneを日本時間に設定してるけど、変わってない。
			ja, _ := time.LoadLocation("Asia/Tokyo")
			return time.Now().In(ja)
		},
		TranslateError: true,
	})

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("DB接続成功")

	return db
}

func CloseDB(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatalln(err)
	}
	sqlDB.Close()
	fmt.Println("DB切断成功")
}
