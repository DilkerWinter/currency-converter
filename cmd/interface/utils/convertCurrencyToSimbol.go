package utils

func ConvertCurrencyToSimbol(currency string) string{
  switch currency{
  case "BRL":
    return "R$"
  case "USD":
    return "$"
  case "EUR":
    return "€"
  case "JPY":
    return "¥"
  default:
    return ""
  }
}
