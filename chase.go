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
	return other_travel_purchases * .05 + drug_stores * .03
}

func (c Chase) sapphirePreferred() float64 {
	var total float64 = -95.0
	total += 50
	
}

func (c Chase) sapphireReserved() float64 {

}