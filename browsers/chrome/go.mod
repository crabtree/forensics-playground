module github.com/crabtree/forensics-playground/browsers/chrome

go 1.19

require github.com/crabtree/forensics-playground/browsers/shared v0.0.0

replace github.com/crabtree/forensics-playground/browsers/shared v0.0.0 => ../shared

require github.com/mattn/go-sqlite3 v1.14.16
