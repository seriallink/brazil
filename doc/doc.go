package doc

import (
    "fmt"
)

type ImplDocumento interface {
    SetCPF()
    SetCNPJ()
    IsCPF()         bool
    IsCNPJ()        bool
}

type Documento struct {
    Numero      string
    Tipo        string
    Formatado   string
    Valido      bool
}

// define os tipos de documento
var DocTipo = struct { CNPJ, CPF, Invalido string } { "CNPJ", "CPF", "Invalido" }

// garante que Documento implementa ImplDocumento
var _ ImplDocumento = (*Documento)(nil)

func New(numero string) Documento {
    return Documento{ Numero:numero, Tipo:DocTipo.Invalido, Valido:false }
}

func Doc(numero string) (doc Documento) {

    // cnpj?
    doc = CNPJ(numero)
    if doc.IsCNPJ() { return }

    // cpf?
    doc = CPF(numero)
    if doc.IsCPF() { return }

    // invalid document!
    doc = New(numero)
    return

}

func CNPJ(numero string) Documento {
    cnpj := New(numero)
    cnpj.SetCNPJ()
    cnpj.Format()
    return cnpj
}

func CPF(numero string) Documento {
    cpf := New(numero)
    cpf.SetCPF()
    cpf.Format()
    return cpf
}

func (this *Documento) Format() {
    if this.Valido && this.Tipo == DocTipo.CNPJ {
        this.Formatado = fmt.Sprintf("%s.%s.%s/%s-%s",this.Numero[0:2],this.Numero[2:5],this.Numero[5:8],this.Numero[8:12],this.Numero[12:])
    } else if this.Valido && this.Tipo == DocTipo.CPF {
        this.Formatado = fmt.Sprintf("%s.%s.%s-%s",this.Numero[0:3],this.Numero[3:6],this.Numero[6:9],this.Numero[9:])
    }
}