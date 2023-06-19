package amex
import ("math")
type Amex struct {
	Supermarkets float64
	Online_retail float64
	Gas_stations float64
	Restaurants float64
	Flights_portal float64
	Hotels float64
	Ubers float64
	Digital_entertainment float64
}

func (a Amex) blueCash() float64 {
	return a.Online_retail * .03 + a.Gas_stations * .03
}

func (a Amex) goldCard() float64 {
	return a.Restaurants * 4 * .007 + a.Supermarkets * 4 * .007 - 250
}

func (a Amex) platinum() float64 {
	var remain_hotels float64 = math.Max(0, a.Hotels - 200)
	var remain_flights float64 = math.Max(0,a.Flights_portal - 200)
	return remain_hotels * 5 * .007 + remain_flights * 5 * .007 - 695
}

func (a Amex) total() float64 {
	var uber_cash float64 = 320.0
	var AET_credit float64 = 240.0
	var remain_ubers float64 = math.Max(0,a.Ubers - uber_cash)
	var remain_AET float64 = math.Max(0, a.Digital_entertainment - AET_credit)
	return a.blueCash() + a.goldCard() + a.platinum() - remain_ubers - remain_AET
}