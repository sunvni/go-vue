package db

import (
	"log"
	"os"

	"github.com/joho/godotenv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"govue/models/entity"
)


func gormConnect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
    DB_HOST := os.Getenv("DB_HOST")
    DB_PORT := os.Getenv("DB_PORT")
    DB_USERNAME := os.Getenv("DB_USERNAME")
    DB_PASSWORD := os.Getenv("DB_PASSWORD")
    DB_NAME := os.Getenv("DB_NAME")
    // MySQLだと文字コードの問題で"?parseTime=true"を末尾につける必要がある
    CONNECT := DB_USERNAME + ":" + DB_PASSWORD + "@tcp(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?parseTime=true"
    db, err := gorm.Open("mysql", CONNECT)

    if err != nil {
        panic(err.Error())
    }
    return db
}

// DBの初期化
func Init() {
    db := gormConnect()

    // コネクション解放解放
    defer db.Close()
    db.AutoMigrate(&entity.Tweet{}) //構造体に基づいてテーブルを作成
}

// データインサート処理
func Insert(content string, title string) {
    db := gormConnect()

    defer db.Close()
    // Insert処理
    db.Create(&entity.Tweet{Content: content, Title: title})
}

//DB更新
func Update(id int, tweetText string) {
    db := gormConnect()
    var tweet entity.Tweet
    db.First(&tweet, id)
    tweet.Content = tweetText
    db.Save(&tweet)
    db.Close()
}

// 全件取得
func GetAll() []entity.Tweet {
    db := gormConnect()

    defer db.Close()
    var tweets []entity.Tweet
    // FindでDB名を指定して取得した後、orderで登録順に並び替え
    db.Order("created_at desc").Find(&tweets)
    return tweets
}

//DB一つ取得
func GetOne(id int) entity.Tweet {
    db := gormConnect()
    var tweet entity.Tweet
    db.First(&tweet, id)
    db.Close()
    return tweet
}

//DB削除
func Delete(id int) {
    db := gormConnect()
    var tweet entity.Tweet
    db.First(&tweet, id)
    db.Delete(&tweet)
    db.Close()
}