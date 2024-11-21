package main

import (
	"fmt"
	"time"
)

/*
 Программа для скачивания большого файла по частям
 Дано: количество частей файла, функция, которая умеет скачивать часть файла
 Задание: написать код, который скачает все части файлов
*/

func main() {
	start := 1
	end := 100
	chunkSize := 5

	var errs []error

	// your solution here

	if errs != nil {
		for _, v := range errs {
			fmt.Println(v.Error())
		}
	}
}

func download(part int) error {
	// some download and save file logic
	time.Sleep(time.Millisecond * 500)

	if part%20 == 0 {
		return fmt.Errorf("error downloading part %d", part)
	}

	return nil
}
