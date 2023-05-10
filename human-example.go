package main

import (
	"fmt"
	"time"
)

const (
    //TimeFormat to format date into
    TimeFormat = "02/01/2006"
)

type AttributesFunc func(*Attributes)

type Attributes struct {
	name string
	lastname string
	eyesColor string
	hairColor string
	height float32
	Weight float32
	birthday time.Time
}

type Human struct {
	Attributes
}

func defaultAttributes() Attributes {
	return Attributes{
		name: string("John"),
		lastname: string("Doe"),
		eyesColor: string("brown"),
		hairColor: string("black"),
		height: float32(1.75),
		Weight: float32(70.00),
		birthday: getDOB(1900, 12, 31),
	}
}

func setEyesColor(color string) AttributesFunc {
	return func(a *Attributes) {
		a.eyesColor = color
	}
}

func setHairColor(color string) AttributesFunc {
	return func(a *Attributes) { 
		a.hairColor = "brown" 
	}
}

func setHeight(n float32) AttributesFunc {
	return func(a *Attributes) { a.height = n }
}

func setWeight(n float32) AttributesFunc {
	return func(a *Attributes) { a.Weight = n }
}

func setBirthday(t time.Time) AttributesFunc {
	return func(a *Attributes) { a.birthday = t}
}

func newHuman(attributes ...AttributesFunc) *Human {
	a := defaultAttributes()
	for _, fn := range attributes {
		fn(&a)
	}
	return &Human{
		Attributes: a,
	}
}

func main() {
	fmt.Println("````````````````````````````````````````````````````````````````````")
	// customized human
	p := newHuman(
		setEyesColor("green"),
		setHairColor("brown"),
		setHeight(float32(177.00)),
		setWeight(float32(71.00)),
		setBirthday(time.Date(1985, time.Month(3), 1, 0, 0, 0, 0, time.UTC)),
	)
	fmt.Println("Human: ", p)
	fmt.Println("birthday formatted: ", p.birthday.Format(TimeFormat))
	fmt.Println("````````````````````````````````````````````````````````````````````")

	// default human
	p = newHuman()
	fmt.Println("Human: ", p)
	fmt.Println("birthday formatted: ", p.birthday.Format(TimeFormat))
	fmt.Println("````````````````````````````````````````````````````````````````````")

}

// helper
func getDOB(year, month, day int) time.Time {
    dob := time.Date(year, time.Month(month), day, 0, 0, 0, 0, time.UTC)
    return dob
}