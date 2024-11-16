package main

import (
	"01-Basic_calculator/calculator"
	"bufio"
	"fmt"
	"os"
)

func menu() {
	var option int
	var number1 int
	var number2 int
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Selecciona la opción que desees")
	options := "1. Sumar 2 números\n" +
		"2. Restar 2 números\n" +
		"3. Multiplicar 2 números\n" +
		"4. Dividir 2 números\n"
	fmt.Println(options)
	_, err := fmt.Scanf("%d", &option)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al seleccionar la opción \"%v\" intentelo de nuevo\n", err)
		fmt.Scanln()
		menu()
	} else if option <= 0 || option > 4 {
		fmt.Fprintln(os.Stderr, "Error, el número debe ser 1, 2, 3 o 4")
		menu()
	}

	fmt.Println("Introduce el primer \"<número1> <número2>\"")
	n, err := fmt.Scanln(&number1, &number2)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error al seleccionar la opción \"%v\" intentelo de nuevo\n", err)
		reader.ReadString('\n')
		menu()
	} else if n != 2 {
		reader.ReadString('\n')
		menu()
	}

	switch option {
	case 1:
		fmt.Println(calculator.Sum(number1, number2))
	case 2:
		fmt.Println(calculator.Rest(number1, number2))
	case 3:
		fmt.Println(calculator.Mult(number1, number2))
	case 4:
		fmt.Println(calculator.Div(float64(number1), float64(number2)))
	default:
		fmt.Print("Ha sucedido un error inesperado, porfavor seleccione de nuevo")
		menu()
	}
}

func main() {
	var input string
	menu()
	for {
		fmt.Println("Para salir escriba \"exit\" o presion Enter para continuar")
		fmt.Scanln(&input)
		if input == "exit" {
			break
		} else if input == "" {
			menu()
		} else {
		}
	}
}
