package entities

type Transaction struct {
	Idnumber   int
	Amount     float32
	Utility    string
	Vendorcode string
	// Affected   map[string]string
}
