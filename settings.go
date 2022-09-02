package main

type settings struct {
	DbHost  string
	DbPort  string
	DbPass  string
	WebPort string
}

func NewSettings() settings {
	return settings{
		"localhost",
		"6379",
		"",
		"8080",
	}
}
