package main

import (
	"fmt"
	"os"
	"runtime"
)

var params [2]string

func main() {
	printUserInfo()
	inputParameters()
	printParameters(params)
}

func printUserInfo() {
	username := os.Getenv("USER")
	if username == "" {
		username = os.Getenv("USERNAME") // для Windows
	}
	fmt.Printf("Имя пользователя: %s\n", username)
	fmt.Printf("Версия Go: %s\n", runtime.Version())
	fmt.Printf("Архитектура: %s\n", runtime.GOARCH)
	fmt.Printf("Операционная система: %s\n", runtime.GOOS)
}

func inputParameters() {
	fmt.Println("\nВвод параметров:")
	for i := 0; i < 2; i++ {
		fmt.Printf("Введите параметр %d: ", i+1)
		fmt.Scan(&params[i])
	}
}

func printParameters(params [2]string) {
	fmt.Println("\nВывод после цикла:")
	for i := 0; i < 2; i++ {
		fmt.Printf("параметр %d = %s\n", i+1, params[i])
	}
}
