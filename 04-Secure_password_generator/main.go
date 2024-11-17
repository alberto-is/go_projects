package main

import (
	caractergen "04-Secure_password_generator/Caracter_gen"
	"flag"
	"fmt"
	"log"
)

func main() {

	length := flag.Int("length", 12, "Longitud de la contraseña (mínimo 4)")
	includeSpecial := flag.Bool("special", false, "Incluir caracteres especiales")
	includeNumbers := flag.Bool("numbers", true, "Incluir números")
	includeUpper := flag.Bool("upper", true, "Incluir letras mayúsculas")
	marks := flag.Bool("marks", true, "Incluir putnos y comas")
	flag.Parse()

	if *length < 4 {
		log.Fatal("La longitud mínima de la contraseña debe ser 4")
	}
	// Generar la contraseña
	password, err := caractergen.GenRandomPassword(*length, *includeSpecial, *includeNumbers, *includeUpper, *marks)
	if err != nil {
		log.Fatalf("Error generando la contraseña: %v", err)
	}

	fmt.Println("Contraseña generada:", password)
}
