package chase


type Chase struct {
	flex_cat float64 
	travel_portal_purchases float64
	other_travel_purchases float64
	drug_stores float64
	restaurants float64
	online_groceries float64
	streaming_services float64
	other_purchases float64

}

func (c Chase) freedomflex() float64 { 
	return c.flex_cat * .05
}


func (c Chase) freedomUnlimited() float64 {
	return c.other_purchases * .05 + c.drug_stores * .03
}


//FIXME: There is some math error in here I believe
func (c Chase) sapphirePreferred() float64 {
	var total float64 = -95.0
	var temp_travel float64 = c.travel_portal_purchases
	temp_travel -= 50

	var total_points float64 = temp_travel * .5
	total_points += c.other_travel_purchases * 2
	total_points += c.restaurants * 3
	total_points += c.streaming_services * 3
	total_points += c.online_groceries * 3
	var total_spent float64 = temp_travel + c.other_travel_purchases + c.restaurants + c.streaming_services + c.online_groceries
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
	return total_points * .01 - temp_travel
}

func (c Chase) sapphireReserved() float64 {

}