# Advent of Code 2024
Repository to contain solutions to [Advent of Code 2024](https://adventofcode.com/2024). All
super rough code - written in golang just for fun to experiment. This code
uses many bad practices.

> **Note: Spoilers containing Advent of Code solutions (obviously)**

## Setup

You just need [go 1.23](https://go.dev/) installed.

## Usage
Running a day:
```bash
go run main.go 1 example
```

Note the input files are not included
since they're unique to each user and [they ask not to share them](https://adventofcode.com/2024/about#faq_copying). You can register and download them at the site.

Create data files for each day:
```bash
./scripts/mk-data-dir.sh
```

Adding days:
```bash
./scripts/add-day.sh 1
```
Then update `./pkg/day/days.go` to call the new day.

