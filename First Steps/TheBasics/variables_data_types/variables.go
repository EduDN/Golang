package main

import "fmt"

func main() {

	var saludo = "hola"
	var nombre = "Edu"
	var edad = 25

	fmt.Println(saludo, nombre, "tu edad es", edad)

	// podemos especficar el tipo de dato de la variable
	var despedida string = "Bye 👋"
	fmt.Println(despedida)

	var booleano = true
	fmt.Println(booleano)

	// declaramos varias vriables en una sola línea
	var nacimiento, mes = 2001, "enero"
	fmt.Println(nacimiento, mes)

	// otra forma de declarar variables
	// de esta forma go infiere el tipo de la misma
	fruta := "Plátano 🍌"
	fmt.Println(fruta)

}
