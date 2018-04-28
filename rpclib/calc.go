package rpclib

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

func (calc *Calculator) Init() {
	calc.Variable = make(map[string]*big.Float)
}
func (calc *Calculator) Create(pair *Pair) error {
	value, _, err := big.ParseFloat(pair.B, 10, 80, big.ToZero)

	if err != nil {
		return err
	}

	if _, ok := calc.Variable[pair.A]; !ok {
		calc.Variable[pair.A] = value
	} else {
		return errors.New("Variable has been already created")
	}

	return nil
}

func (calc *Calculator) Delete(name string) error {

	if _, ok := calc.Variable[name]; ok {
		delete(calc.Variable, name)
	} else {
		return errors.New("Variable not exsit")
	}
	return nil
}

func (calc *Calculator) Update(pair *Pair) error {

	if _, ok := calc.Variable[pair.A]; ok {
		value, _, err := big.ParseFloat(pair.B, 10, 80, big.ToZero)

		if err != nil {
			return err
		}
		calc.Variable[pair.A] = value
	} else {
		return errors.New("Variable not exsit")
	}

	return nil
}

func (calc *Calculator) DoCal(pair *Pair, method string) (*big.Float, error) {

	var a, b, result *big.Float
	result = big.NewFloat(0)
	if _, ok := calc.Variable[pair.A]; ok {
		a = calc.Variable[pair.A]
	} else {
		value, _, err := big.ParseFloat(pair.B, 10, 80, big.ToZero)

		if err != nil {
			return big.NewFloat(0), err
		}
		a = value
	}
	if _, ok := calc.Variable[pair.B]; ok {
		b = calc.Variable[pair.B]
	} else {
		value, _, err := big.ParseFloat(pair.B, 10, 80, big.ToZero)

		if err != nil {
			return big.NewFloat(0), err
		}
		b = value
	}

	switch method {
	case "add":
		return result.Add(a, b), nil
	case "sub":
		return result.Sub(a, b), nil
	case "mul":
		return result.Mul(a, b), nil
	case "div":
		return result.Quo(a, b), nil
	default:
		return result, errors.New("no such method")
	}
}

func main() {
	calc := Calculator{Variable: make(map[string]*big.Float)}

	pair := &Pair{"test", "3.1415961231231231231231231123"}

	//create test
	err := calc.Create(pair)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	err = calc.Create(pair)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("test:{value:%v type:%T}\n", calc.Variable["test"], calc.Variable["test"])

	//update test
	pair = &Pair{"test", "123456678"}
	err = calc.Update(pair)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	pair = &Pair{"test123", "3.1415961231231231231231231123"}
	err = calc.Update(pair)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("test:{value:%v type:%T}\n", calc.Variable["test"], calc.Variable["test"])

	//add
	pair = &Pair{"test", "1234567"}
	result, err := calc.DoCal(pair, "add")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("result: {value:%v, type:%T}\n", result, result)

	//sub
	pair = &Pair{"test", "1234567"}
	result, err = calc.DoCal(pair, "sub")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("result: {value:%v, type:%T}\n", result, result)
	//mul
	pair = &Pair{"test", "1234567"}
	result, err = calc.DoCal(pair, "mul")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("result: {value:%v, type:%T}\n", result, result)
	//div
	pair = &Pair{"test", "1234567"}
	result, err = calc.DoCal(pair, "div")
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("result: {value:%v, type:%T}\n", result, result)
	//delete test
	err = calc.Delete("blahblah")
	if err != nil {
		fmt.Println(err)
	}

	err = calc.Delete("test")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("test:{value:%v type:%T}\n", calc.Variable["test"], calc.Variable["test"])
}
