package main

import (
	"math"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
    myApp := app.New()
    window := myApp.NewWindow("Calculator")

    var firstNumber float64
    var operation string

    input := widget.NewEntry()
    input.SetPlaceHolder("Enter a number")

    makeNumButton := func(num string) *widget.Button {
        return widget.NewButton(num, func() {
            input.SetText(input.Text + num)
        })
    }

	buttons := make([]*widget.Button, 10)
	for i := 0; i <= 9; i++ {
		buttons[i] = makeNumButton(strconv.Itoa(i))
	}

	plus := widget.NewButton("+", func() {
       firstNumber, _ = strconv.ParseFloat(input.Text, 64)
       operation = "+"
       input.SetText("")
   })

	minus := widget.NewButton("-", func() {
		firstNumber, _ = strconv.ParseFloat(input.Text, 64)
		operation = "-"
		input.SetText("")
	})

	multiply := widget.NewButton("*", func() {
		firstNumber, _ = strconv.ParseFloat(input.Text, 64)
		operation = "*"
		input.SetText("")
	})

	divide := widget.NewButton("/", func() {
		firstNumber, _ = strconv.ParseFloat(input.Text, 64)
		operation = "/"
		input.SetText("")
	})

	mod := widget.NewButton("%", func() {
		firstNumber, _ = strconv.ParseFloat(input.Text, 64)
		operation = "%"
		input.SetText("")
	})

	power := widget.NewButton("^", func() {
		firstNumber, _ = strconv.ParseFloat(input.Text, 64)
		operation = "^"
		input.SetText("")
	})

   	equals := widget.NewButton("=", func() {
       secondNumber, _ := strconv.ParseFloat(input.Text, 64)
       var result float64

       switch operation {
       case "+":
           result = firstNumber + secondNumber
       case "-":
           result = firstNumber - secondNumber
       case "*":
           result = firstNumber * secondNumber
	   case "/":
           result = firstNumber / secondNumber
       case "%":
           result = math.Mod(firstNumber, secondNumber)
       case "^":
           result = math.Pow(firstNumber, secondNumber)
       }

       input.SetText(strconv.FormatFloat(result, 'f', -1, 64))
   })

   	clear := widget.NewButton("C", func() {
       input.SetText("")
       firstNumber = 0
       operation = ""
   })

   	numPad := container.NewGridWithColumns(3,
       buttons[7], buttons[8], buttons[9],
       buttons[4], buttons[5], buttons[6],
       buttons[1], buttons[2], buttons[3],
       buttons[0], clear, equals,
   )

   	operations := container.NewVBox(plus, minus, multiply, divide, mod, power)

   	content := container.NewVBox(
       input,
       container.NewHBox(
           numPad,
           operations,
       ),
   )

   window.SetContent(content)
   window.Resize(fyne.NewSize(300, 400))
   window.Show()
   myApp.Run()
}
