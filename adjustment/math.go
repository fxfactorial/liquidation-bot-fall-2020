package adjustment

import (
	"fmt"
	"log"
	"math/big"
	"watch_for_arb/adjustment/decimal"
)

func Scale_down_by(amount *big.Int, scale_factor uint8) *big.Int {
	flt, _, _ := big.ParseFloat(
		fmt.Sprintf(`1e%d`, scale_factor), 10, 0, big.ToNearestEven,
	)
	var i = new(big.Int)
	i, _ = flt.Int(i)

	return new(big.Int).Div(amount, i)
}

func Token_amount_to_big_int(in_amount int, scale_factor uint8) *big.Int {
	as_amount := fmt.Sprintf(`%de%d`, in_amount, scale_factor)
	flt, _, err := big.ParseFloat(as_amount, 10, 0, big.ToNearestEven)

	if err != nil {
		log.Fatal(err)
	}

	var in_amount_ = new(big.Int)
	in_amount_, _ = flt.Int(in_amount_)
	return in_amount_
}

var (
	one_hundred = decimal.NewDec(100)
	two         = decimal.NewDec(2)
)

func Relative_percentage_diff(price_one, price_two decimal.Dec) decimal.Dec {
	numerator := price_one.Sub(price_two).Abs()
	denominator := price_one.Add(price_two).Quo(two)
	return numerator.Quo(denominator).Mul(one_hundred)
}
