package main

//go:generate jsonenums -type=CountryType

type CountryType int

const (
	CountryNotActiviated CountryType = iota
	CountryActivated
)
