package citi




type citi struct {
	other_purchases float64
	custom_categories float64

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

}