#!/bin/bash
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m' # No Color

for ((i=1;i<100;i++)) # цикл будет длится 100 раз потом завершится или завершится если найдет неверный ответ
do
linear_stats1=$(./linear-stats) # запускаем linear-stats
linear_stats2=$(go run main.go data.txt) # тут запускаете свой .go файл

echo -e "---Linear-Stats---"
echo -e "$linear_stats1\n"

echo -e "-----main.go-----"
echo -e "$linear_stats2\n"

if [[ $linear_stats1 == $linear_stats2 ]]
  then
    printf "${GREEN}------EQUAL------\n\n${NC}"
    else 
    printf "${RED}----NOT-EQUAL----\n\n${NC}"
    break
fi
done