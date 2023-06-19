package capital_one


type Capital struct {
	Dining float64
	Entertainment float64
	Streaming_services float64
	Groceries float64
	Other_purchases float64
	Hotels_and_car_rentals float64
	Flights float64
}


func (c Capital) savor() float64 {
	return c.Dining * .03 + c.Entertainment * .03 + c.Streaming_services * .03 + c.Groceries * .03
}

func (c Capital) venturex() float64 {
	return c.Hotels_and_car_rentals * 10 * .01 + c.Flights * 5 * .01 + c.Other_purchases * 2 * .01 + (300 + 100 - 395)
}

func (c Capital) Total() float64 {
	return c.savor() + c.venturex()
}