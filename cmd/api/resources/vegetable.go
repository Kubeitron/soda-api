package resources

type Vegetable struct {
	Color    string  `json:"color" xml:"color"`
	Weight   float64 `json:"weight" xml:"weight"`
	Name     string  `json:"name" xml:"name"`
	Vitamins string  `json:"vitamins" xml:"vitamins"`
	Calories int     `json:"calories" xml:"calories"`
}
