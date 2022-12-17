package main

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-lambda-go/lambda"
)

type MyEvent struct {
	MonthNumber uint
	WeekDay     uint
}

func HandleRequest(ctx context.Context, event MyEvent) (string, error) {
	if event.MonthNumber > 12 {
		return "", fmt.Errorf("the provided nmonth number is greater than 12. Got: %d", event.MonthNumber)
	}
	m := time.Month(event.MonthNumber)
	w := time.Weekday(event.WeekDay)
	return fmt.Sprintf("The month is %s and the day is %s", m.String(), w.String()), nil
}

func main() {
	lambda.Start(HandleRequest)
}
