package menu

import (
  "fmt"
	"currency-converter/cmd/interface/utils"
	"currency-converter/cmd/controller"
  "github.com/atotto/clipboard"
	"github.com/charmbracelet/huh"
)



func ShowResults(result controller.ConversionResult, currencyToBeConverted string, amount float64) {
    if result.Converted == nil {
        fmt.Println("No conversion result available.")
        return
    }

    options := []huh.Option[string]{}
    var selectedValue float64   

    for currency, value := range result.Converted {
        option := fmt.Sprintf("%s: %s %.2f", currency, utils.ConvertCurrencyToSimbol(currency), value)
        options = append(options, huh.NewOption(option, currency)) 
    }

    var selectedResult string
    title := fmt.Sprintf("%s %.2f (%s) Converted to:", utils.ConvertCurrencyToSimbol(currencyToBeConverted), amount, currencyToBeConverted)
    selectResult := huh.NewSelect[string]().
        Title(title).
        Description("Choose the result to copy to clipboard").
        Options(options...).
        Value(&selectedResult).
        WithTheme(huh.ThemeCatppuccin())

    if err := selectResult.Run(); err != nil {
        fmt.Println("Error during selection:", err)
        return
    }

    for currency, value := range result.Converted {
        if currency == selectedResult {
            selectedValue = value 
            break
        }
    }

    err := clipboard.WriteAll(fmt.Sprintf("%.2f", selectedValue))
    if err != nil {
        fmt.Println("Error copying to clipboard:", err)
        return
    }

    fmt.Printf(
      "Selected result '%s' currency (value: %s %.2f) has been copied to clipboard.\n",
      selectedResult,
      utils.ConvertCurrencyToSimbol(selectedResult),
      selectedValue)
}

