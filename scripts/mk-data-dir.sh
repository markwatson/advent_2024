#!/bin/bash

mkdir data/

for i in {1..25}
do
  mkdir data/day$i
  touch data/day$i/example.txt
  touch data/day$i/input.txt
done
