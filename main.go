package main

import "fmt"

func main() {
	var (
		priceOneApple = 5.99
		priceOnePier  = 7
		ourBalance    = 23
	)

	fmt.Println("Ціні за 8 яблук та 9 груш складає ", priceOneApple*9+float64(priceOnePier)*8)
	fmt.Println("Ми можемо купити ", ourBalance/priceOnePier, " грушу/груші/груш")
	fmt.Println("Ми можемо купити ", (ourBalance)%int(priceOneApple), "яблуко/яблука/яблук")
	if 2*(float64(priceOnePier)+priceOneApple) < float64(ourBalance) {
		fmt.Println("Покупка вдала")
	} else {
		fmt.Println("Покупка не вдала")
	}
}
