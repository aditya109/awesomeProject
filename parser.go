package main

import (
	"strconv"
	"strings"
)

// Expand expands the cron field into a list of allowed values
func (cf *CronField) Expand() {
	// Handle the wildcard character '*', which means every value in the range
	if cf.Field == "*" {
		for i := cf.MinValue; i <= cf.MaxValue; i++ {
			cf.AllowedVals = append(cf.AllowedVals, i)
		}
		return
	}

	// Split the field by commas to handle lists of values or ranges
	for _, part := range strings.Split(cf.Field, ",") {
		// Handle step values (e.g., "*/15" or "1-5/2")
		if strings.Contains(part, "/") {
			split := strings.Split(part, "/")
			startEnd := split[0]
			step, _ := strconv.Atoi(split[1])

			start := cf.MinValue
			end := cf.MaxValue

			// If there's a specific start value, set it
			if startEnd != "*" {
				start, _ = strconv.Atoi(startEnd)
			}

			// Add values to AllowedVals based on the step
			for i := start; i <= end; i += step {
				cf.AllowedVals = append(cf.AllowedVals, i)
			}
		} else if strings.Contains(part, "-") {
			// Handle ranges (e.g., "1-5")
			split := strings.Split(part, "-")
			start, _ := strconv.Atoi(split[0])
			end, _ := strconv.Atoi(split[1])

			// Add all values in the range to AllowedVals
			for i := start; i <= end; i++ {
				cf.AllowedVals = append(cf.AllowedVals, i)
			}
		} else {
			// Handle single values (e.g., "5")
			val, _ := strconv.Atoi(part)
			cf.AllowedVals = append(cf.AllowedVals, val)
		}
	}
}
