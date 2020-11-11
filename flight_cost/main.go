package main

/*
	The purpose of this program is to provide a "highly accurate" (read: distance not negative) estimation
	for the cost of domestic airfare travel using lat/lon coordinates.
*/
import (
	"fmt"

	"github.com/AnonymousRecess/flight_cost/distance"
	"github.com/AnonymousRecess/flight_cost/point"
	"github.com/AnonymousRecess/flight_cost/rasm"
	"github.com/AnonymousRecess/flight_cost/statemap"
)

/*
	main is the point of entry to every program in Go
*/
func main() {
	fmt.Println("This program shows the average cost to fly between U.S states\n" +
		"Enter two U.S States for program start")
	fmt.Print("First State: ")
	var state1 string
	fmt.Scanln(&state1) // Read user input from stdin
	fmt.Print("Second State: ")
	var state2 string
	fmt.Scanln(&state2)

	var scoord1 point.Point = statemap.StateLocations[state1]
	var scoord2 point.Point = statemap.StateLocations[state2]

	distanceInK := distance.Length(scoord1.X, scoord1.Y, scoord2.X, scoord2.Y) // shorthand initialization - only works within functions (Gets the distance between the two states)

	fmt.Print("\nThe distance between these two states is approximately ")
	fmt.Println(fmt.Sprintf("%.2f km", distanceInK)) // Sprintf is very similar to C, the % indicates data type and the leading arguments are substituted in order

	fare := (50 + ((distanceInK / 1.609) * 0.11)) // Linear relationship for airefare per mile provided by rome2rio @ https://www.rome2rio.com/blog/2013/01/02/170779446/#:~:text=Fare%20%3D%20%2450%20%2B%20(Distance%20*,11%20cents%20per%20mile%20travelled.
	fmt.Println(fmt.Sprintf("The projected cost using a domestic flight average is $%.2f \n", fare))
	for airline, profit := range rasm.AirlineProfit { // format for a foreach loop over a map
		money := (distanceInK / 1.609) * profit
		fmt.Println(fmt.Sprintf("%s made a profit of $%.2f", airline, money))
	}
}
