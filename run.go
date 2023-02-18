package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("car.json")
	if err != nil {
		fmt.Println(err)
	}

	bytes, err := io.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}

	car := Car{}
	err = json.Unmarshal(bytes, &car)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(car)
}
