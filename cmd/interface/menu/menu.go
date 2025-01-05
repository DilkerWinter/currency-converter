package menu

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"currency-converter/cmd/interface/utils"
	"strconv"
)

func Menu() {
	utils.ClearTerminal()

	var catppuccin *huh.Theme = huh.ThemeCatppuccin()
	var theme *huh.Theme = catppuccin

	var currencyToBeConverted string
	var currencyToConvert string
	var amountInput string
	var amount float64 

	selectCurrency := huh.NewSelect[string]().
		Title("Choose a currency to be converted").
		Options(
			huh.NewOption("BRL - Real", "BRL"),
			huh.NewOption("USD - Dollar", "USD"),
			huh.NewOption("EUR - Euro", "EUR"),
			huh.NewOption("JPY - Yen", "JPY"),
		).
		Value(&currencyToBeConverted).
		WithTheme(theme)

	if err := selectCurrency.Run(); err != nil {
		fmt.Println("Error during selection:", err)
		return
	}

	selectCurrencyDestiny := huh.NewSelect[string]().
		Title("Choose the currency to convert to").
		Options(
			huh.NewOption("BRL - Real", "BRL"),
			huh.NewOption("USD - Dollar", "USD"),
			huh.NewOption("EUR - Euro", "EUR"),
			huh.NewOption("JPY - Yen", "JPY"),
		).
		Value(&currencyToConvert).
		WithTheme(theme)

	if err := selectCurrencyDestiny.Run(); err != nil {
		fmt.Println("Error during destination currency selection:", err)
		return
	}

	
for {
	selectAmountToConvert := huh.NewInput().
		Title("Insert the amount of money to convert to: " + currencyToConvert).
		Placeholder("Ex: 252.39").
		Description("Amount in " + currencyToBeConverted).
		Value(&amountInput).
		Validate(func(input string) error {
			_, err := strconv.ParseFloat(input, 64)
			if err != nil {
				return fmt.Errorf("Invalid amount input. Please enter a valid number.")
			}
			return nil
		})

	if err := selectAmountToConvert.Run(); err != nil {
		fmt.Println("Error during amount input:", err)
		return
	}

	var err error
	amount, err = strconv.ParseFloat(amountInput, 64)
	if err == nil {
		break 
	}
}

	fmt.Println("Selected currency to be converted:", currencyToBeConverted)
	fmt.Println("Selected currency to convert to:", currencyToConvert)
	fmt.Printf("Amount to convert: %s %.2f\n", utils.ConvertCurrencyToSimbol(currencyToConvert), amount)
}
