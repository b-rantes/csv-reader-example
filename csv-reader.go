package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("dados.csv")
	defer file.Close()

	if err != nil {
		fmt.Println("Erro ao abrir o arquivo csv: ", err)
		return
	}

	reader := csv.NewReader(file)
	vendas := make(map[int][]Venda)

	for {

		row, err := reader.Read()

		if err == io.EOF {
			fmt.Println("Arquivo acabou")
			break
		}

		if err != nil {
			fmt.Println("Erro ao ler linha: ", err)
			return
		}

		dadosVenda := strings.Split(row[0], ";")
		mes, _ := strconv.Atoi(dadosVenda[0])
		nome := dadosVenda[1]
		valor, _ := strconv.ParseFloat(dadosVenda[2], 32)
		venda := Venda{Mes: mes, Vendedor: nome, Valor: valor}
		vendas[venda.Mes] = append(vendas[venda.Mes], venda)
	}
	vendas[0] = nil

	for chave, valor := range vendas {
		fmt.Printf("A media de vendas do mês %d foi: %.2f", chave, calculaMediaVendas(valor))
		fmt.Println("")
	}
}

func calculaMediaVendas(v []Venda) float64 {
	somaVendas := 0.00
	quantidadeVendas := len(v)
	for i := 0; i < quantidadeVendas; i++ {
		somaVendas += v[i].Valor
	}

	return somaVendas / float64(quantidadeVendas)
}
