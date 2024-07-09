# Cron Parser

This is a simple Golang application that parses a cron string and expands each field to show the times at which it will run. It handles the standard cron format with five time fields plus a command.

## Usage

```sh
# build the main.go
go build -o main     

 # run the executable with cmd argument  
./main "*/15 0 1,15 * 1-5 /usr/bin/find"       
```