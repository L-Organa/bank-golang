package main

import "fmt"

type ContaCorrrente struct{
	titular string
	numeorAgencia int
	numeroConta int
	saldo float64
}

func main()  {
	contaDaCarol := ContaCorrrente{
					"Carolina", 1234, 555555-1, 455.54}
	fmt.Println(contaDaCarol)
}
