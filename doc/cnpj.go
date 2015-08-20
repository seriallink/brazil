package doc

import (
    "fmt"
    "regexp"
    "strings"
    "strconv"
)

func (this *Documento) IsCNPJ() bool {
    return this.Tipo == DocTipo.CNPJ && this.Valido
}

func (this *Documento) SetCNPJ() {

    // set documento = cnpj (mesmo se for invalido)
    this.Tipo = DocTipo.CNPJ

    // remove caracteres nao numericos
    reg,_ := regexp.Compile("[^0-9]")
    this.Numero = reg.ReplaceAllString(this.Numero,"")

    // CNPJ vazio
    if this.Numero == "" {
        this.Valido = false; return
    }

    // CPNJ > 14 caracteres
    if len(this.Numero) > 14 {
        this.Valido = false; return
    }

    // elimina CNPJs invalidos conhecidos
    for _, c := range "0123456789" {
        if strings.Trim(this.Numero,string(c)) == "" {
            this.Valido = false; return
        }
    }

    // pad zeros a esquerda
    this.Numero = fmt.Sprintf("%014s",this.Numero)

    // valida DVs
    for x:=0; x<=1; x++ {

        // primeiro dv
        tamanho := len(this.Numero) - 2

        // os digitos nao mudam
        digitos := this.Numero[tamanho:]

        // segund dv (avanca uma casa)
        if x== 1 { tamanho++ }

        numeros := this.Numero[:tamanho]
        posicao := tamanho - 7

        // calculo e resultado da validacao
        calculo,resultado := 0,0

        for _, numero := range numeros {
            n,_ := strconv.Atoi(string(numero))
            calculo += n * posicao
            posicao--
            if posicao < 2 { posicao = 9 }
        }

        if mod := calculo % 11; mod >= 2 {
            resultado = 11 - calculo % 11
        }

        if digito,_ := strconv.Atoi(digitos[x:x+1]); digito != resultado {
            this.Valido = false; return
        }

    }

    // cnpj valido
    this.Valido = true
}