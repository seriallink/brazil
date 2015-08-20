package doc

import (
    "fmt"
    "regexp"
    "strings"
    "strconv"
)

func (this *Documento) IsCPF() bool {
    return this.Tipo == DocTipo.CPF && this.Valido
}

func (this *Documento) SetCPF() {

    // set documento = cpf (mesmo se for invalido)
    this.Tipo = DocTipo.CPF

    // remove caracteres nao numericos
    reg,_ := regexp.Compile("[^0-9]")
    this.Numero = reg.ReplaceAllString(this.Numero,"")

    // CNPJ vazio
    if this.Numero == "" {
        this.Valido = false; return
    }

    // CPF > 11 caracteres
    if len(this.Numero) > 11 {
        this.Valido = false; return
    }

    // elimina CPFs invalidos conhecidos
    for _, c := range "0123456789" {
        if strings.Trim(this.Numero,string(c)) == "" {
            this.Valido = false; return
        }
    }

    // pad zeros a esquerda
    this.Numero = fmt.Sprintf("%011s",this.Numero)

    // armazena as somatorias para os calculos dos digitos verificadores
    sum1, sum2 := 0, 0

    // somatoria dos digitos do cpf (primeiro digito)
    for i:=0; i<9; i++ {
        digito,_ := strconv.Atoi(this.Numero[i:i+1])
        sum1 += digito * (10 - i)
    }

    // calculo do primeiro digito
    rev := 11 - (sum1 % 11)
    if rev == 10 || rev == 11 {rev=0}

    // validacao do primeiro digito
    if digito, _ := strconv.Atoi(this.Numero[9:10]); digito != rev {
        this.Valido = false; return
    }

    // somatoria dos digitos do cpf (segundo digito)
    for i:=0; i<10; i++ {
        digito,_ := strconv.Atoi(this.Numero[i:i+1])
        sum2 += digito * (11 - i)
    }

    // calculo do segundo digito
    rev = 11 - (sum2 % 11)
    if rev == 10 || rev == 11 {rev=0}

    // validacao do segundo digito
    if digito, _ := strconv.Atoi(this.Numero[10:11]); digito != rev {
        this.Valido = false; return
    }

    // cnpj valido
    this.Valido = true
}