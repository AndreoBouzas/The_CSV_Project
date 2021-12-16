package Reader

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

var path string = "files/colaboradores.csv"

type Colaboradores struct {
	UnidadeFederativa string
	Cidade            string
	Nome              string
	Idade             int
	Cargo             string
	Salario           float64
}

const (
	unidadeFederativa int = iota
	cidade
	nome
	idade
	cargo
	salario
)

func InitFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	reader := csv.NewReader(file)
	for {
		row, err := reader.Read()
		if err == io.EOF {
			log.Fatal(err)
		}
		if err != nil {
			log.Fatal(err)
		}
		idad, err := strconv.Atoi(row[idade])
		if err != nil {
			log.Fatal(err)
		}
		sala := float64(salario)
		fmt.Println(Colaboradores{
			UnidadeFederativa: row[unidadeFederativa],
			Cidade:            row[cidade],
			Nome:              row[nome],
			Idade:             idad,
			Cargo:             row[cargo],
			Salario:           sala,
		})
	}

}
