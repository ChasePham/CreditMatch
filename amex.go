package amex

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