
package controller

import (
  "currency-converter/cmd/service"
  "fmt"
)


type ConversionResult struct {
  SourceCurrency string
  Amount         float64
  Converted      map[string]float64
}

func CurrencyConverter(amount float64, currencyToBeConverted string, currencyToConvert []string) (ConversionResult, error) {
  rates, err := service.GetExchangeRateFromApi()
  if err != nil {
    return ConversionResult{}, fmt.Errorf("Failed to obtain exchange rates: %v", err)
  }

  rateToUSD, exists := rates[currencyToBeConverted]
  if !exists {
    return ConversionResult{}, fmt.Errorf("Exchange rate to %s not found", currencyToBeConverted)
  }

  amountInUSD := amount / rateToUSD
  convertedRates := make(map[string]float64)

  for _, currency := range currencyToConvert {
    rateFromUSD, exists := rates[currency]
    if !exists {
      fmt.Printf("Exchange rate to %s not found.\n", currency)
      continue
    }

    convertedAmount := amountInUSD * rateFromUSD
    convertedRates[currency] = convertedAmount
  }

  return ConversionResult{
    SourceCurrency: currencyToBeConverted,
    Amount:         amount,
    Converted:      convertedRates,
  }, nil
}

