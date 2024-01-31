package main

import (
	"fmt"
	"os"
)

type OdiaEnStruct struct {
	EN_key string
	OD_val string
}

func main() {
	fmt.Println("Hello World!")
	////fir
	data, err := os.ReadFile("En-Od-v3.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(len(data))
}
