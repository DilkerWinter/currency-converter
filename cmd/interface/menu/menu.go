
package menu

import (
	"fmt"
	"time"
	"os"
	"strings"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/huh/spinner"
	"currency-converter/cmd/interface/utils"
	"currency-converter/cmd/controller"
	"strconv"
)

func Menu() {
	utils.ClearTerminal()

	var catppuccin *huh.Theme = huh.ThemeCatppuccin()
	var theme *huh.Theme = catppuccin

	var currencyToBeConverted string
	var currencyToConvert []string
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

	selectCurrencyDestiny := huh.NewMultiSelect[string]().
		Title("Choose the currency to convert to").
		Description("Press X to select one or more currencies and press enter to confirm.").
		Options(
			huh.NewOption("BRL - Real", "BRL"),
			huh.NewOption("USD - Dollar", "USD"),
			huh.NewOption("EUR - Euro", "EUR"),
			huh.NewOption("JPY - Yen", "JPY"),
		).
		Validate(func(c []string) error {
			if len(c) <= 0 {
				return fmt.Errorf("Select at least one currency to convert")
			}
			return nil
		}).
		Value(&currencyToConvert).
		WithTheme(theme)

	if err := selectCurrencyDestiny.Run(); err != nil {
		fmt.Println("Error during destination currency selection:", err)
		return
	}

	for {
		selectAmountToConvert := huh.NewInput().
			Title("Insert the amount of money to convert to: " + strings.Join(currencyToConvert, ", ")).
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

	


  var result controller.ConversionResult  
  var err error

  action := func() {
    time.Sleep(2 * time.Second)
    result, err = controller.CurrencyConverter(amount, currencyToBeConverted, currencyToConvert)
    if err != nil {
        fmt.Println("Error during conversion:", err)
        return
    }
}

  if err := spinner.New().
    Title(fmt.Sprintf("Converting %s %.2f to: %s", utils.ConvertCurrencyToSimbol(currencyToBeConverted), amount, strings.Join(currencyToConvert, " | "))).
    Action(action).
    Run(); err != nil {
    fmt.Println("Error during spinner:", err)
    os.Exit(1)
  }

  ShowResults(result, currencyToBeConverted, amount)
}





