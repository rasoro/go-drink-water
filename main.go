package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/adhocore/gronx/pkg/tasker"
)

func main() {
	interval := GetMinutesFromArgs()
	RemindmeToHydratate(interval)
}

func RemindmeToHydratate(minuteInterval int) {
	log.Printf("You will be reminded to drink water every %v minute(s)", minuteInterval)

	interval := fmt.Sprintf("*/%v * * * *", minuteInterval)

	taskr := tasker.New(tasker.Option{
		Verbose: false,
	})

	notification := NewNotification(
		"It's time to hydratation",
		"go drink water my fella",
	)

	taskr.Task(interval, func(ctx context.Context) (int, error) {
		log.Println("Go drink whater now")
		notification.Notify()
		return 0, nil
	})
	taskr.Until(10 * time.Hour)
	taskr.Run()
}

func GetMinutesFromArgs() int {
	if len(os.Args) > 1 {
		if minutes, _ := strconv.Atoi(os.Args[1]); minutes != 0 {
			return minutes
		}
	}
	return 30
}
