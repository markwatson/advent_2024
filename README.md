# advent_2024
Repository to contain solutions to [Advent of Code 2024](https://adventofcode.com/2024). All
super rough code - written in golang just for fun to experiment. This code
uses many bad practices.

## Usage
Running a day:
```bash
go run main.go 1 example
```

Create data files for each day:
```bash
./scripts/mk-data-dir.sh
```

Adding days:
```bash
./scripts/add-day.sh 1
```
Then update `./pkg/day/days.go` to call the new day.

