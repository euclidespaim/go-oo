package contas

import "go-orientado/clientes"

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	Saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {
	podeSacar := valorDoSaque > 0 && valorDoSaque <= c.Saldo
	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque realizado com sucesso \n"
	} else {
		return "Saldo insuficiente \n"
	}
}

func (c*ContaCorrente) Depositar(valorDoDeposito float64) (string, float64) {
	if valorDoDeposito > 0.0 {
		c.Saldo += valorDoDeposito
		return "Depósito realizado com sucesso:", c.Saldo
	} else {
		return "Valor do depósito menor que zero:", c.Saldo
	}
}

func (c*ContaCorrente) Transferir(valorDaTransferencia float64, contaDestino *ContaCorrente) bool {
	if valorDaTransferencia < c.Saldo && valorDaTransferencia > 0 {
		c.Saldo -= valorDaTransferencia
		contaDestino.Depositar(valorDaTransferencia)
		return true
	} else {
		return false
	}
}