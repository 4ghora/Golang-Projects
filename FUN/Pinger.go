package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func pingWebsite(url string, count int) {
	for i := 0; i < count || count == 0; i++ {
		startTime := time.Now()

		response, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error pinging %s: %v\n", url, err)
			return
		}
		defer response.Body.Close()

		elapsedTime := time.Since(startTime)

		fmt.Printf("%s responded with status code %d in %s\n", url, response.StatusCode, elapsedTime)

		if count > 0 && i+1 < count {
			time.Sleep(1 * time.Second) // sleep for 1 second between pings
		}
	}
}

func main() {
	var website string
	fmt.Print("Enter the website URL to ping: ")
	fmt.Scanln(&website)

	if website == "" {
		fmt.Println("Please enter a valid website URL.")
		return
	}

	var countOption int
	fmt.Println("\nChoose how many times to ping:")
	fmt.Println("1. 1 time")
	fmt.Println("2. 4 times")
	fmt.Println("3. Infinite")
	fmt.Println("4. Custom")

	fmt.Print("Enter your choice: ")
	fmt.Scanln(&countOption)

	interruptChannel := make(chan os.Signal, 1)
	signal.Notify(interruptChannel, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-interruptChannel
		fmt.Println("\nCtrl+C detected. Program execution completed.")
		os.Exit(0)
	}()

	switch countOption {
	case 1:
		pingWebsite(website, 1)
	case 2:
		pingWebsite(website, 4)
	case 3:
		fmt.Println("Pinging indefinitely. Press Ctrl+C to stop.")
		pingWebsite(website, 0)
	case 4:
		var customCount int
		fmt.Print("Enter the custom count: ")
		fmt.Scanln(&customCount)
		pingWebsite(website, customCount)
	default:
		fmt.Println("Invalid choice. Please choose a valid option.")
	}

	fmt.Println("\nProgram execution completed.")
}
