package chase

import "math"

type Chase struct {
	Flex_cat float64 
	Travel_portal_purchases float64
	Other_travel_purchases float64
	Drug_stores float64
	Restaurants float64
	Online_groceries float64
	Streaming_services float64
	Other_purchases float64
	Hotels_and_car_rentals float64
}

func (c Chase) freedomflex() float64 { 
	return c.Flex_cat * .05
}


func (c Chase) freedomUnlimited() float64 {
	return c.Other_purchases * .05 + c.Drug_stores * .03
}


//FIXME: There is some math error in here I believe
func (c Chase) sapphirePreferred() float64 {
	var temp_travel float64 = c.Travel_portal_purchases
	temp_travel -= 50

	var total_points float64 = math.Max(0,temp_travel * 5)
	total_points += c.Other_travel_purchases * 2
	total_points += c.Restaurants * 3
	total_points += c.Streaming_services * 3
	total_points += c.Online_groceries * 3
	total_points += c.Hotels_and_car_rentals * 5
	var total_spent float64 = temp_travel + c.Other_travel_purchases + c.Restaurants + c.Streaming_services + c.Online_groceries + c.Hotels_and_car_rentals
	total_points += total_spent * .1



	var additional_value float64 = temp_travel / .0125

	if additional_value > total_points {
		additional_value -= total_points
		total_points = 0
		temp_travel -= additional_value * .0125
	} else {
		total_points -= additional_value
		temp_travel = 0
	}

	// Travel portal is already discounted within this method
	return total_points * .01 - math.Max(0,temp_travel) - 95
}

func (c Chase) sapphireReserved() float64 {
	var temp_travel float64 = c.Travel_portal_purchases
	temp_travel -= 300

	var total_points float64 = math.Max(0,temp_travel * 5)
	total_points += c.Other_travel_purchases * 2
	total_points += c.Restaurants * 3
	total_points += c.Streaming_services * 3
	total_points += c.Online_groceries * 3
	total_points += c.Hotels_and_car_rentals * 10
	var total_spent float64 = temp_travel + c.Other_travel_purchases + c.Restaurants + c.Streaming_services + c.Online_groceries + c.Hotels_and_car_rentals
	total_points += total_spent * .1



	var additional_value float64 = temp_travel / .015

	if additional_value > total_points {
		additional_value -= total_points
		total_points = 0
		temp_travel -= additional_value * .015
	} else {
		total_points -= additional_value
		temp_travel = 0
	}

	// Travel portal is already discounted within this method
	return total_points * .01 - math.Max(0,temp_travel) - 550
}


func (c Chase) total() float64 {
	var flex float64 = c.freedomflex()
	var unlimited float64 = c.freedomUnlimited()
	var preferred float64 = c.sapphirePreferred()
	var reserved float64 = c.sapphireReserved()
	var result float64 = 0
	result += flex + unlimited + math.Max(preferred,reserved)
	return result
}