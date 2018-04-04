package models

type ExchangeRate struct {
	Model
	CurrencyCodeFrom       string
	CurrencyCodeTo         string
	ConversionValue        int64
	InverseConversionValue int64
}
