package main

type CronField struct {
	Name        string
	Field       string
	MinValue    int
	MaxValue    int
	AllowedVals []int
}

func NewCronField(name, field string, minValue, maxValue int) *CronField {
	return &CronField{
		Name:     name,
		Field:    field,
		MinValue: minValue,
		MaxValue: maxValue,
	}
}
