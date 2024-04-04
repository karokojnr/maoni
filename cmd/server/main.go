package main

import "fmt"

// Run - responsible for the instantiation and startup of our go application
func Run() error {
	fmt.Println("Starting Maoni...")
	return nil
}

func main() {
	fmt.Println("Maoni API")
	if err := Run(); err != nil {
		fmt.Println(err)
	}
}
