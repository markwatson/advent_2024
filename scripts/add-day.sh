#!/bin/bash

day=$1
if [ -z $day ]; then
  echo "Usage: $0 <day>"
  exit 1
fi

# Code
cp -r pkg/days/day0 pkg/days/day$day
mv pkg/days/day$day/day0.go pkg/days/day$day/day$day.go
sed -i '' -e "s/ay0/ay$day/g" pkg/days/day$day/day$day.go

# This doesn't really work:
# Update days.go
# cat ./pkg/days/days.go | gsed "/RegisterAllDays() {/a Register("$day", &day$day.Day$day{})" > ./pkg/days/days2.go
# mv ./pkg/days/days2.go ./pkg/days/days.go
# goimports pkg/days/days.go

# Add input
mkdir data/day$day
touch data/day$day/example.txt
touch data/day$day/input.txt

