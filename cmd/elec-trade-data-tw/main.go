package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"

	"github.com/Nextdrive-AaronWu/elec-trade-data/internal/api"
	"github.com/Nextdrive-AaronWu/elec-trade-data/internal/db"
)

func main() {
	_ = godotenv.Load(".env")

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		log.Fatal("DATABASE_URL not set")
	}

	pool, err := pgxpool.New(context.Background(), dbURL)
	if err != nil {
		log.Fatalf("failed to connect db: %v", err)
	}
	defer pool.Close()

	// 日期判斷：CLI > env > 今日
	var startDate string
	if len(os.Args) > 1 && os.Args[1] != "" {
		startDate = os.Args[1]
	} else if os.Getenv("START_DATE") != "" {
		startDate = os.Getenv("START_DATE")
	} else {
		startDate = (func() string {
			return time.Now().AddDate(0, 0, 1).Format("2006-01-02")
		})()
	}

	// 印出本次處理的日期
	log.Printf("Processing date: %s", startDate)

	// 呼叫 API 取得資料
	resp, err := api.FetchDailyData(startDate)
	if err != nil {
		log.Fatalf("fetch api error: %v", err)
	}
	if resp.Code != 200 {
		log.Fatalf("api response code not 200: %v", resp.Msg)
	}
	if len(resp.Data) == 0 {
		log.Println("no data to insert")
		return
	}

	// 批次入庫
	err = db.InsertTradeDataBatch(context.Background(), pool, resp.Data)
	if err != nil {
		log.Fatalf("db insert error: %v", err)
	}
	log.Printf("insert success, count: %d", len(resp.Data))
}
