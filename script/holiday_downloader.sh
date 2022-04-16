set -xe

declare -A urls=(
  ["jp"]="https://www8.cao.go.jp/chosei/shukujitsu/syukujitsu.csv"
)

for key in "${!urls[@]}"; do
  original_path="nholidays/${key}/national_holidays.csv"
  download_path="nholidays/${key}/tmp.csv"

  wget -O "${download_path}" "${urls[${key}]}"
  nkf -w --overwrite "${download_path}"

  if [ ! -e "${original_path}" ];then
    echo "${original_path} not exist"
    mv "${download_path}" "${original_path}"
  elif [ -n "$(diff "${download_path}" "${original_path}")" ]; then
    echo "diff found ${download_path} ${original_path}"
    mv "${download_path}" "${original_path}"
  else
    echo "no updated"
    rm "${download_path}"
  fi

done
