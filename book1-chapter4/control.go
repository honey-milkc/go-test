package book1_chapter4

import (
	"fmt"
	"math"
)

type Coster interface {
	Cost() float64
}

func displayCost(c Coster) {
	fmt.Println("cost: ", c.Cost())
}

type Item struct {
	name     string
	price    float64
	quantity int
}

func (t Item) Cost() float64 {
	return t.price * float64(t.quantity)
}

func (t Item) String() string {
	return fmt.Sprintf("[%s] %.0f", t.name, t.Cost())
}

type DiscountItem struct {
	Item
	discountRate float64
}

func (t DiscountItem) Cost() float64 {
	return t.Item.Cost() * (1.0 - t.discountRate/100)
}

func (t DiscountItem) String() string {
	return fmt.Sprintf("%s => %.0f(%.0f%s DC)", t.Item.String(), t.Cost(), t.discountRate, "%")
}

type Rental struct {
	name         string
	feePerDay    float64
	periodLength int
	RentalPeriod
}

type RentalPeriod int

const (
	Days RentalPeriod = iota
	Weeks
	Months
)

func (p RentalPeriod) ToDays() int {
	switch p {
	case Weeks:
		return 7
	case Months:
		return 30
	default:
		return 1
	}
}

func (r Rental) Cost() float64 {
	return r.feePerDay * float64(r.ToDays()*r.periodLength)
}

func (r Rental) String() string {
	return fmt.Sprintf("[%s] %.0f", r.name, r.Cost())
}

type Itemer interface {
	Coster
	fmt.Stringer
}

type Order struct {
	Itemer
	taxRate float64
}

func (o Order) Cost() float64 {
	return o.Itemer.Cost() * (1.0 + o.taxRate/100)
}

func (o Order) String() string {
	return fmt.Sprintf("Total price: %.0f(tax rate: %.2f)\n\tOrder details: %s",
		o.Cost(), o.taxRate, o.Itemer.String())
}

func DoItem() {
	shoes := Item{"Women's Walking Shoes", 30000, 2}
	eventShoes := DiscountItem{
		Item{"Sport Shoes", 50000, 3},
		10.00,
	}

	// fmt.Println(shoes.Cost())
	// fmt.Println(eventShoes.Cost())
	// fmt.Println(eventShoes.Item.Cost())
	displayCost(shoes)
	displayCost(eventShoes)

	shirt := Item{"Men's Slim-Fit Shirt", 25000, 3}
	video := Rental{"Interstellar", 1000, 3, Days}

	displayCost(shirt)
	displayCost(video)

	fmt.Println(Order{shirt, 10.00})
	fmt.Println(Order{video, 5.00})
}

type shaper interface {
	area() float64
}

func describe(s shaper) {
	fmt.Println("area : ", s.area())
}

type rect struct {
	width  float64
	height float64
}

func (r rect) area() float64 {
	return r.height * r.width
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
}

func DoInterface() {
	r := rect{3, 4}
	c := circle{2.5}
	describe(r)
	describe(c)

	var v interface{}

}
