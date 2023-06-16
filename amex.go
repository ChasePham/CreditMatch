package amex
import ("math")
type amex struct {
	supermarkets float64
	online_retail float64
	gas_stations float64
	restaurants float64
	flights_portal float64
	hotels float64
	ubers float64
	digital_entertainment float64
}

func (a amex) blueCash() float64 {
	return a.online_retail * .03 + a.gas_stations * .03
}

func (a amex) goldCard() float64 {
	return a.restaurants * 4 * .007 + a.supermarkets * 4 * .007 - 250
}

func (a amex) platinum() float64 {
	var remain_hotels float64 = math.Max(0, a.hotels - 200)
	var remain_flights float64 = math.Max(0,a.flights_portal - 200)
	return remain_hotels * 5 * .007 + remain_flights * 5 * .007 - 695
}

func (a amex) total() float64 {
	var uber_cash float64 = 320.0
	var AET_credit float64 = 240.0
	var remain_ubers float64 = math.Max(0,a.ubers - uber_cash)
	var remain_AET float64 = math.Max(0, a.digital_entertainment - AET_credit)
	return blueCash() + goldCard() + platinum() - remain_ubers - remain_AET
}