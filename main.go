package main

import (
	"fmt"
	"time"
)

func main() {
	errs := DownloadFile(100, 5)
	for _, err := range errs {
		fmt.Println(err)
	}
}

func DownloadFile(partsCount, chunkSize int) []error {
	// from := 1
	// Your code here
	return nil
}

func downloadPart(part int) error {
	// some download and save file logic. Do not edit this func!!!
	time.Sleep(time.Millisecond * 250)

	if part%20 == 0 {
		return fmt.Errorf("error downloading part %d", part)
	}

	return nil
}
