package car

type Person struct {
	Name     string
	Lastname string
}

type CarPublic struct {
	Label string
	Owner Person
	year  int16
}
