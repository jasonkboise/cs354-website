package rasm

/*
	AirlineProfit is a map where the name of the Airline Service is the key
	and the value is U.S ¢ per mile (converted to $) as a domestic RASM(revenue per available seat mile)
	Info Provided by Statista.com ~https://www.statista.com/statistics/527810/us-airlines-domestic-revenue-per-asm/~
*/
var AirlineProfit = map[string]float64{
	"Alaska":    0.165,
	"Delta":     0.15,
	"American":  0.145,
	"United":    0.14,
	"JetBlue":   0.124,
	"Southwest": 0.12,
	"Hawaiin":   0.119,
	"Allegiant": 0.097,
	"Spirit":    0.096,
	"Frontier":  0.067, // Not a Sponser.. unless
}
