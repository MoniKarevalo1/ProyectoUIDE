/*
@autor: Mónica Arévalo
@fecha: 30/10/2024
@descrupción: Clase 1 de programación go
*/

package main

// librerías a importar y a utilizar dentro del archivo
// librería fmt permite imprimir en pantalla

import (
	"fmt"
)

func main() {
	fmt.Println("Bienvenidos a la asignatura de programación orientada a objetos en go")
	//variables var
	var nombre string
	// asignar un dato
	// nombre ="Milton"
	fmt.Println("Ingrese su nombre:")
	//leer el dato desde el teclado , fmt.Scanln() &variable capturar ingresos en go
	fmt.Scanln(&nombre)
	fmt.Printf("Hola, %s!\n", nombre)

	//Vamos a completar de manera directa los siguientes tipos de datos
	var numero1 int = 42
	var decimal float64 = 3.1416
	var texto1 string = "Bienvenido a la UIDE"
	var boolean bool = false

	fmt.Println("===Tipos de Datos Básicos===")
	fmt.Printf("El tipo de variable es: %d\n", numero1)
	fmt.Printf("El tipo de variable es: %.2f\n", decimal)
	fmt.Printf("El tipo de variable es: %s\n", texto1)
	fmt.Printf("El tipo de variable es: %t\n", boolean)

	//asignación especial :=
	numero2 := 90
	decimal2 := 3.14
	texto2 := "Hola Copito"
	boolean2 := true

	fmt.Println("===Tipos de Datos Básicos Parte 2===")
	fmt.Printf("El tipo de variable es: %d\n", numero2)
	fmt.Printf("El tipo de variable es: %.2f\n", decimal2)
	fmt.Printf("El tipo de variable es: %s\n", texto2)
	fmt.Printf("El tipo de variable es: %t\n", boolean2)

	//Operadores Básicos
	// + - * / % logicos
	// operadores incrementales ++ --
	// operadores logicos and - && or || - not!
	// operadores de comparacion mayor mayor= menor= == !=
	// operadores de asignación = += -= /= %=

	// calcular el promedio de 3 números

	nota1 := 10
	nota2 := 8
	nota3 := 3

	promedio := (nota1 + nota2 + nota3) / 3
	fmt.Println("El promedio de las notas es:", promedio, "", "es su nota final")

	fmt.Println("====Claculaora Básica===")
	var num1, num2 float64
	var operacion string

	fmt.Print("Ingrese el primer nombre: ")
	fmt.Scanln(&num1)

	fmt.Print("Ingrese el segundo nombre: ")
	fmt.Scanln(&num2)

	fmt.Print("Seleccione la operación que quiere realizar [+, -, *, /]: ")
	fmt.Scanln(&operacion)

	fmt.Println("los datos ingresados son:")
}
