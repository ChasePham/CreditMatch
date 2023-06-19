package citi

import "math"


type Citi struct {
	Other_purchases float64
	Custom_categories float64
	Travel_portal_purchases float64 // Excludes air travel ONLY
	Other_travel_purchases float64 // Includes air travel with everything else from portal purchases

}

func (c Citi) doubleCash() float64 {
	return c.Other_purchases * .02
}

func (c Citi) customCash() float64 {
	var result float64 = 0
	result += 6000 * 5 * .01 + (c.Custom_categories - 6000) * .01
	return result
}

func (c Citi) premier() float64 {
	var remaining_portal_purchases float64 = math.Max(0,c.Travel_portal_purchases - 100)
	return remaining_portal_purchases * 10 * .01 + c.Other_travel_purchases * 3 * .01
}

func (c Citi) total() float64 {
	return c.doubleCash() + c.customCash() + c.premier()
}