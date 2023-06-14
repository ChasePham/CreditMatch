package capital_one


type capital struct {
	dining float64
	entertainment float64
	streaming_services float64
	groceries float64
	other_purchases float64
	hotels_and_car_rentals float64
	flights float64
}


func (c capital) savor() float64 {
	return c.dining * .03 + c.entertainment * .03 + c.streaming_services * .03 + c.groceries * .03
}

func (c capital) venturex() float64 {
	
}