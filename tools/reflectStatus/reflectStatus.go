package main

type CurrencyStatusType int

const (
	CurrencyNotActivated CurrencyStatusType = iota
	CurrencyActivated
)

// ReflectStatus is used to solve the problem that the
// Elastic search parser can reflect all term key into
// the 'real' struct or type defined in the program
// automatically
func ReflectStatus(x string) {

}

func main() {
	ReflectStatus("CurrencyNotActivated")
}
