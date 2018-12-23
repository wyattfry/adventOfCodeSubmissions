#!/bin/bash

SUM=0

add() {
  SUM=$(( SUM + $1 ))
  echo $SUM
}

for NUM in $(cat 1_input); do
  add $NUM
done
