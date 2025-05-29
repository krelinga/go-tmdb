#! /usr/bin/bash

for file in testdata/*; do
    go run ./tools/print_testdata "$file" | fgrep -q '"Game of Thrones"' && echo "$file contains 'Game of Thrones'"
done