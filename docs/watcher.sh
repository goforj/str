#!/usr/bin/env bash

go install github.com/bokwoon95/wgo@latest

echo "Watching for .go file changes to regenerate documentation..."

wgo -verbose -file=.go -xdir examples \
  go -C docs run ./examplegen :: \
  go -C docs run ./readme
