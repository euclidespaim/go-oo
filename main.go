package main

import "fmt"

type ContaCorrente struct {
	titular       string
	numeroAgencia int
	numeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque realizado com sucesso \n"
	} else {
		return "Saldo insuficiente \n"
	}
}

func main() {
	contaDoElon := ContaCorrente{
		titular: "Elon Musk", 
		numeroAgencia: 321, 
		numeroConta: 1234567, 
		saldo: 100.21,
	}
	fmt.Println(contaDoElon.Sacar(100))
	fmt.Println("Saldo na conta do titular: ", contaDoElon.titular, "= ", contaDoElon.saldo)
	fmt.Println()

	contaDoBezos := ContaCorrente{"Jeff Bezos", 123, 1234567, 200.00}
	fmt.Println("Saldo na conta do titular: ", contaDoBezos.titular, "= ", contaDoBezos.saldo)
	fmt.Println()

	var contaDoStark *ContaCorrente
	contaDoStark = new(ContaCorrente)
	contaDoStark.titular = "Tony Stark"
	contaDoStark.numeroAgencia = 123
	fmt.Println(*contaDoStark)
	fmt.Println()
}

