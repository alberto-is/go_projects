package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime)
	log.SetPrefix("Error:")

	var ch chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-ch
		fmt.Println("\n Recibido:", sig)
		fmt.Printf("Cerrando programa\n")
		os.Exit(0)
	}()

	for {

		fmt.Println("Indique los dos números que quiere dividir")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		args := scanner.Text()
		words := strings.Split(args, " ")
		if len(words) != 2 {
			log.Printf("Número de argumentos errone %d", len(words))
		} else if words[1] == "0" {
			log.Panic("Se ha intentado dividir por 0")
		} else {

			number1, _ := strconv.ParseFloat(words[0], 32)
			number2, _ := strconv.ParseFloat(words[1], 32)

			fmt.Println("El resutlaod es: ", number1/number2)

		}

	}

}
