package citi

import "math"



type citi struct {
	other_purchases float64
	custom_categories float64
	travel_portal_purchases float64 // Excludes air travel ONLY
	other_traveL_purchases float64 // Includes air travel with everything else from portal purchases

}

func (c citi) doubleCash() float64 {
	return c.other_purcahses * .02
}

func (c citi) customCash() float64 {
	var result float64 = 0
	result += 6000 * 5 * .01 + (c.custom_categories - 6000) * .01
	return result
}

func (c citi) premier() float64 {
	var result float64 = 0
	var remaining_portal_purchases float64 = math.Max(0,c.travel_portal_purchases - 100)
	return remaining_portal_purchases * 10 * .01 + c.other_travel_purchases * 3 * .01
}