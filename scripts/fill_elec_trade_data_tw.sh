#!/bin/bash
# 補齊 2021-07-01 到 2025-11-06 的資料

START="2021-07-01"
END="2021-08-31"


# 日期遞增函式
# macOS 用：date -j -f "%Y-%m-%d" "$d" "+%Y-%m-%d" -v+1d
# Linux 用：date -I -d "$d + 1 day"
next_date() {
  # macOS 寫法（順序需正確）
  echo $(date -v+1d -j -f "%Y-%m-%d" "$1" "+%Y-%m-%d")
  # Linux 寫法（如需支援，請改用下行）
  # echo $(date -I -d "$1 + 1 day")
}

BIN_PATH="$(dirname "$0")/../bin/elec-trade-data-tw"

CUR="$START"
while [[ "$CUR" < "$END" || "$CUR" == "$END" ]]; do
  echo "Processing $CUR"
  "$BIN_PATH" "$CUR"
  sleep 1
  CUR="$(next_date "$CUR")"
done
