package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: cronparser \"<cron_expression>\"")
		return
	}

	cronExpression := os.Args[1]
	parts := strings.Fields(cronExpression)

	if len(parts) != 6 {
		fmt.Println("Invalid cron expression format. Expected format: \"<minute> <hour> <day of month> <month> <day of week> <command>\"")
		return
	}

	minuteField := NewCronField("minute", parts[0], 0, 59)
	hourField := NewCronField("hour", parts[1], 0, 23)
	dayOfMonthField := NewCronField("day of month", parts[2], 1, 31)
	monthField := NewCronField("month", parts[3], 1, 12)
	dayOfWeekField := NewCronField("day of week", parts[4], 0, 6)

	minuteField.Expand()
	hourField.Expand()
	dayOfMonthField.Expand()
	monthField.Expand()
	dayOfWeekField.Expand()

	fmt.Printf("%-14s%s\n", "minute", joinInts(minuteField.AllowedVals))
	fmt.Printf("%-14s%s\n", "hour", joinInts(hourField.AllowedVals))
	fmt.Printf("%-14s%s\n", "day of month", joinInts(dayOfMonthField.AllowedVals))
	fmt.Printf("%-14s%s\n", "month", joinInts(monthField.AllowedVals))
	fmt.Printf("%-14s%s\n", "day of week", joinInts(dayOfWeekField.AllowedVals))
	fmt.Printf("%-14s%s\n", "command", parts[5])
}
