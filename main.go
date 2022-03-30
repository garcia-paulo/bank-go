package main

import (
	"fmt"

	"github.com/garcia-paulo/bank-go/contas"
)

func PagarBoleto(conta contas.Conta, valor float64) {
	conta.Sacar(valor)
}

func main() {
	paulo := contas.ContaPoupanca{}
	paulo.Depositar(100)
	PagarBoleto(&paulo, 50)
	fmt.Println(paulo.ObterSaldo())
}
