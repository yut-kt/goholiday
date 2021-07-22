#!/usr/local/bin/bash

set -xe

declare -A urls=(
  ["jp"]="https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
)

for key in "${!urls[@]}"; do
  original_path="nholidays/${key}.csv"
  download_path="nholidays/tmp_${key}.csv"

  wget -O "${download_path}" "${urls[${key}]}"
  nkf -w --overwrite "${download_path}"
  # File update check
  if [ -z "$(diff ${original_path} ${download_path})" ]; then
    rm "${download_path}"
  else
    mv "${download_path}" "${original_path}"
  fi

done
