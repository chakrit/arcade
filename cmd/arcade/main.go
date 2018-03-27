package main

import "log"

func main() {
	defer log.Println("stopped")
	if err := RootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
