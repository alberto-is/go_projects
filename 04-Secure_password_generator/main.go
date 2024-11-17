package main

import (
	caractergen "04-Secure_password_generator/Caracter_gen"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	length         int
	includeSpecial bool
	includeNumbers bool
	includeUpper   bool
	marks          bool
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, `

Generador de Contraseñas Seguras v1.0
------------------------------------
Uso: %s [opciones]

Opciones:
	-l, --length  int     Longitud de la contraseña (mínimo 4) (default 12)
	-s, --special         Incluir caracteres especiales
	-n, --numbers         Incluir números (default true)
	-u, --upper          Incluir letras mayúsculas (default true)
	-m, --marks          Incluir puntos y comas (default true)

Ejemplos:
	%s -l 16             # Contraseña de 16 caracteres
	%s --length 16 -s    # 16 caracteres con especiales

	z
`, os.Args[0], os.Args[0], os.Args[0])
	}

	flag.IntVar(&length, "length", 12, "Longitud de la contraseña (mínimo 4)")
	flag.IntVar(&length, "l", 12, "Longitud de la contraseña (mínimo 4)")

	flag.BoolVar(&includeSpecial, "special", false, "Incluir caracteres especiales")
	flag.BoolVar(&includeSpecial, "s", false, "Incluir caracteres especiales")

	flag.BoolVar(&includeNumbers, "numbers", true, "Incluir números")
	flag.BoolVar(&includeNumbers, "n", true, "Incluir números")

	flag.BoolVar(&includeUpper, "upper", true, "Incluir letras mayúsculas")
	flag.BoolVar(&includeUpper, "u", true, "Incluir letras mayúsculas")

	flag.BoolVar(&marks, "marks", true, "Incluir puntos y comas")
	flag.BoolVar(&marks, "m", true, "Incluir puntos y comas")

	flag.Parse()

	if length < 4 {
		log.Fatal("La longitud mínima de la contraseña debe ser 4")
	}
	// Generar la contraseña
	password, err := caractergen.GenRandomPassword(length, includeSpecial, includeNumbers, includeUpper, marks)
	if err != nil {
		log.Fatalf("Error generando la contraseña: %v", err)
	}

	fmt.Println("Contraseña generada:", password)
}
