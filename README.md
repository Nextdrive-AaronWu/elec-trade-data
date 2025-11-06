# elec-trade-data

## 專案簡介
自動抓取台電 API（每日電力交易資料），存入 PostgreSQL 資料庫，支援批次補資料與每日排程。

## 環境需求
- Go 1.20+（建議最新版）
- PostgreSQL 12+
- 依賴套件：pgx, godotenv

## 安裝與設定
1. 安裝依賴：
	```sh
	go mod tidy   # 下載並整理所有 go.mod 依賴
	```
2. 安裝開發工具（建議）：
	```sh
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install golang.org/x/tools/cmd/goimports@latest
	```
2. 設定 `.env` 檔案：
	```env
	API_BASE_URL="https://etp.taipower.com.tw/api/infoboard/settle_value/query"
	DATABASE_URL="postgres://user:pass@host:port/dbname?sslmode=disable"
	START_DATE="2025-11-06" # 可選，預設今日
	```

## 資料庫 migration
1. 執行 migrations/postgres/20251106000000_create_elec_trade_data_tw_dtable.up.sql 建表
2. 主鍵 id 為 SERIAL，自動遞增
3. tran_date + tran_hour 為 unique constraint，重複執行自動覆蓋舊資料

## 執行方式
1. 開發：
	```sh
	go run ./cmd/elec-trade-data-tw [YYYY-MM-DD]
	```
2. build binary：
	```sh
	go build -o bin/elec-trade-data-tw ./cmd/elec-trade-data-tw
	```
3. 執行 binary：
	```sh
	./bin/elec-trade-data-tw [YYYY-MM-DD]
	```
4. CLI 日期參數優先，未指定則用 .env 或今日

## 批次補資料
執行 `scripts/fill_elec_trade_data_tw.sh`，自動補齊 2021-07-01 到 2025-11-06 資料。

## crontab/自動化
可將 binary 加入 crontab，每日自動執行：
```
0 2 * * * /path/to/elec-trade-data-tw
```

## 錯誤處理與重試
- 資料重複時自動覆蓋（ON CONFLICT UPDATE）
- 失敗天可重跑，不影響資料正確性
