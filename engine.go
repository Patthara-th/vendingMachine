package vendingMachine

import (	
	"sort"
)

type vendingMachine struct {
	current     int
	coin        map[string]int
	items       map[string]int
	changecoins map[int]string
	value       []int
	insertedcoin		[]string
}

func NewVendingMachine(x, y map[string]int) *vendingMachine {

	value := make([]int, 0)
	var changecoins map[int]string
	changecoins = make(map[int]string)
	for k, v := range x {
		changecoins[v] = k
		value = append(value, v)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(value)))
	return &vendingMachine{0, x, y, changecoins, value,make([]string,0)}

}

func (v *vendingMachine) InsertCoin(c string) {

	v.current += v.coin[c]
	v.insertedcoin = append(v.insertedcoin,c)

}

func (v *vendingMachine) GetInsertedMoney() int {

	return v.current

}

func (v *vendingMachine) SelectSD() string {

	if v.current >= v.items["SD"] {
		v.insertedcoin = make([]string,0)
		v.current -= v.items["SD"]
		if v.current > 0 {
			return "SD" + v.getchange()
		}
		return "SD"
	}	
	return "Not Enough Coins"
}

func (v *vendingMachine) SelectCC() string {

	if v.current >= v.items["CC"] {
		v.insertedcoin = make([]string,0)
		v.current -= v.items["CC"]
		if v.current > 0 {
			return "CC" + v.getchange()
		}
		return "CC"
	}	
	return "Not Enough Coins"
}

func (v *vendingMachine) getchange() string {
		var change string
	for _, x := range v.value {
		for v.current >= x {
			change += ", " + v.changecoins[x]
			v.current -= x
		}
	}
	return change
}

func (v *vendingMachine) CoinReturn() string {

	var change string
	
	for _, x := range v.insertedcoin {
			change += ", " + x			
	}
	v.insertedcoin = make([]string,0)
	return change[2:len(change)]

}