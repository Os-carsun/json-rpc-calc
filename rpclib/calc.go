package main

import (
	"errors"
	"fmt"
	"math/big"
)

type Pair struct {
	A, B string
}

type Calculator struct {
	Variable map[string]*big.Float
}

func (calc *Calculator) Create(pair *Pair) error {
	value, _, err := big.ParseFloat(pair.B, 10, 80, big.ToZero)

	if err != nil {
		return err
	}

	_, ok := calc.Variable[pair.A]

	if !ok {
		calc.Variable[pair.A] = value
	} else {
		return errors.New("Variable has been already created")
	}

	return nil
}

func main() {
	calc := Calculator{Variable: make(map[string]*big.Float)}

	pair := &Pair{"test", "3.1415961231231231231231231123"}
	err := calc.Create(pair)
	if err != nil {
		fmt.Printf(">%v<", err)
	}
	err = calc.Create(pair)
	if err != nil {
		fmt.Printf(">>%v<<", err)
	}
	fmt.Printf("%v %T", calc.Variable["test"], calc.Variable["test"])
}
