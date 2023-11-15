package main

import(
	"fmt"
	"time"
)

func main() {
	fmt.Println("-----MRT TEXT------")

	city := []string{"Lebak Bulus", "Istora", "Setiabudi", "Bundaran HI"}

	for i := 0; i < len(city); i++ {
		fmt.Print(" >>>>> ")
		fmt.Print(city[i])
		time.Sleep(time.Second *4)
	}
 }