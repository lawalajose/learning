package piscine

type food struct {
	preptime int
}

func FoodDeliveryTime(order string) int {
	kk := 0
	burger := food{}
	nuggets := food{}
	chips := food{}
	burger.preptime = 15
	chips.preptime = 10
	nuggets.preptime = 12
	if order == "burger" {
		kk = burger.preptime
	} else if order == "chips" {
		kk = chips.preptime
	} else if order == "nuggets" {
		kk = nuggets.preptime
	} else {
		return 404
	}
	return kk
}
