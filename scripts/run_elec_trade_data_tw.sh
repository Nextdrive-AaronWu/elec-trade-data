#!/bin/bash

# --- 專案根目錄路徑，通常是這個腳本所在的目錄的上一層 ---
SCRIPT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" &> /dev/null && pwd )"
PROJECT_ROOT=$(dirname "$SCRIPT_DIR")

# 從 .env 檔案加載環境變數 (僅限開發環境或本地測試，生產環境應使用更安全的秘密管理方式)
# 如果在生產環境運行，此處的 .env 文件應該被忽略，且環境變數會從其他地方注入
if [ -f "$PROJECT_ROOT/.env" ]; then
    echo "Loading environment variables from $PROJECT_ROOT/.env"
    source "$PROJECT_ROOT/.env"
fi

# 設定要處理的日期，例如總是前一天
# 可以在 crontab 傳入，或在此處計算
# export START_DATE=${1:-$(date -d "yesterday" +%Y-%m-%d)} # 如果沒有傳參數，則預設為昨天

# --- 建立 logs 目錄 ---
LOG_DIR="$PROJECT_ROOT/logs"
mkdir -p "$LOG_DIR"

# --- 程式路徑 ---
# 假設編譯後的執行檔在專案根目錄下的 bin/ 資料夾
# 請確保你的 Go build 命令將執行檔輸出到這個位置
EXECUTABLE_PATH="$PROJECT_ROOT/bin/elec-trade-data-tw"

# 檢查執行檔是否存在
if [ ! -f "$EXECUTABLE_PATH" ]; then
    echo "Error: Executable not found at $EXECUTABLE_PATH"
    exit 1
fi

# --- 執行 Go 程式 ---
echo "Starting elec-trade-data-tw at $(date)"
# 將 Go 程式的輸出導向日誌文件
$EXECUTABLE_PATH "$@" >> "$PROJECT_ROOT/logs/elec-trade.log" 2>&1

# --- 錯誤處理 ---
EXIT_CODE=$?
if [ $EXIT_CODE -ne 0 ]; then
    echo "elec-trade-data-tw failed with exit code $EXIT_CODE at $(date)"
    exit $EXIT_CODE
else
    echo "elec-trade-data-tw completed successfully at $(date)"
fi
