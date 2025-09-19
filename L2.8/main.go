package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

// Функция NPT сервер  обернута  для проверки vet  и golint
func NTPTime() error {
	currentTime, err := ntp.Time("0.beevik-ntp.pool.ntp.org")
	if err != nil {
		return fmt.Errorf("failed to get time: %w", err)
	}
	fmt.Println("Текущее время :", currentTime)
	return nil
}

func main() {
	if err := NTPTime(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
