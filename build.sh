#!/usr/bin/env sh

for CMD in `ls cmd/*.go`; do
  FILENAME=$(basename $CMD)
  EXE_NAME=${FILENAME%.*}

  go build -o ./dist/$EXE_NAME $CMD
done