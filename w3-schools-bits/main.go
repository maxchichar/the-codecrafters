package main

import(
	"fmt"
)

func main()  {
	var SportCars = [7]string{"Porsche 718", "Porsche 911", "Lotus emira", "Alphine A110", "BMW M3/M4", "Aston martin vantage", "BMW M2"}

	SportCars[5] = "Ferrari 296 speciale A"
	fmt.Println(SportCars)
}