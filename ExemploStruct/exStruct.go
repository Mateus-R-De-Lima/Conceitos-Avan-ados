package ExemploStruct

import (
	"encoding/json"
	"fmt"
)

// ============================================================================
// DEFINIÇÃO DE STRUCTS
// ============================================================================

// Pessoa representa uma pessoa com seus dados básicos
type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     int
}

// User representa um usuário com credenciais e tags JSON para serialização
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

// ============================================================================
// MÉTODOS (Receivers) DA STRUCT Pessoa
// ============================================================================

// Saudacao retorna uma frase com o nome completo da pessoa.
func (p Pessoa) Saudacao() string {
	return fmt.Sprintf("Olá, %s %s!", p.Nome, p.Sobrenome)
}

// AnosFaltam retorna quantos anos faltam até a pessoa chegar em uma idade alvo.
func (p Pessoa) AnosFaltam(alvo int) int {
	if alvo <= p.Idade {
		return 0
	}
	return alvo - p.Idade
}

// AtualizarNome atualiza o nome da pessoa (método com receiver pointer)
func (p *Pessoa) AtualizarNome(novoNome string) {
	if p == nil {
		return
	}
	p.Nome = novoNome
}

// ============================================================================
// FUNÇÕES AUXILIARES (Construtores)
// ============================================================================

// NovoPessoa cria uma nova instância de Pessoa com os valores fornecidos
func NovoPessoa(nome, sobrenome string, idade int) Pessoa {
	return Pessoa{Nome: nome, Sobrenome: sobrenome, Idade: idade}
}

// AtualizarIdade atualiza a idade de uma pessoa através de ponteiro
func AtualizarIdade(p *Pessoa, novaIdade int) {
	if p == nil {
		return
	}
	p.Idade = novaIdade
}

// ============================================================================
// FUNÇÃO PRINCIPAL
// ============================================================================

// ExemploStruct demonstra diferentes formas de usar structs em Go
func ExemploStruct() {
	fmt.Println("==========================================================")

	fmt.Println("           EXEMPLOS DE STRUCTS EM GO")
	fmt.Println("==========================================================")

	exemploDeclaracaoSimples()
	exemploStructVazio()
	exemploDeclaracaoParcial()
	exemploPonteiros()
	exemploCopia()
	exemploComTags()

	fmt.Println("==========================================================")

	fmt.Println("           FIM DOS EXEMPLOS")
	fmt.Println("==========================================================")

}

// ============================================================================
// EXEMPLOS INDIVIDUAIS
// ============================================================================

// exemploDeclaracaoSimples mostra como declarar uma struct com todos os campos
func exemploDeclaracaoSimples() {
	fmt.Println("\n[1] DECLARAÇÃO SIMPLES COM TODOS OS CAMPOS")
	fmt.Println("==========================================================")

	pessoa1 := NovoPessoa("João", "Silva", 30)
	fmt.Printf("Struct: %+v\n", pessoa1)
	fmt.Printf("Nome completo: %s\n", pessoa1.Saudacao())
	fmt.Printf("Anos faltando para 40: %d\n", pessoa1.AnosFaltam(40))
}

// exemploStructVazio mostra uma struct com campos vazios (valores zero)
func exemploStructVazio() {
	fmt.Println("\n[2] STRUCT COM VALORES VAZIOS (Zero Values)")
	fmt.Println("==========================================================")

	pessoa2 := Pessoa{}
	fmt.Printf("Struct vazio: %+v\n", pessoa2)
	fmt.Printf("Tipo: %T\n", pessoa2)
	fmt.Println("Nota: string vazia=\"\", int vazio=0")
}

// exemploDeclaracaoParcial mostra como atribuir apenas alguns campos
func exemploDeclaracaoParcial() {
	fmt.Println("\n[3] DECLARAÇÃO PARCIAL (Alguns campos)")
	fmt.Println("==========================================================")

	pessoa3 := Pessoa{Nome: "Maria", Sobrenome: "Oliveira", Idade: 25}
	fmt.Printf("Struct com valores nomeados: %+v\n", pessoa3)

	fmt.Println()
	pessoa3Incompleta := Pessoa{Nome: "Carlos"}
	fmt.Printf("Struct incompleta (campos faltantes vazios): %+v\n", pessoa3Incompleta)
}

// exemploPonteiros mostra o uso de ponteiros para modificar structs
func exemploPonteiros() {
	fmt.Println("\n[4] TRABALHANDO COM PONTEIROS")
	fmt.Println("==========================================================")

	// Criar uma struct via ponteiro
	pessoa4 := &Pessoa{Nome: "Carlos", Sobrenome: "Santos", Idade: 25}
	fmt.Printf("Struct via ponteiro (antes): %+v\n", pessoa4)

	// Modificar através do ponteiro
	AtualizarIdade(pessoa4, 28)
	fmt.Printf("Struct via ponteiro (depois): %+v\n", pessoa4)

	// Usar método com receiver pointer
	pessoa4.AtualizarNome("Roberto")
	fmt.Printf("Após atualizar nome: %+v\n", pessoa4)
	fmt.Println("Nota: Ponteiros permitem modificações da struct original")
}

// exemploCopia mostra que structs são copiadas por valor em Go
func exemploCopia() {
	fmt.Println("\n[5] CÓPIA DE STRUCTS (Por Valor)")
	fmt.Println("==========================================================")

	pessoa5 := NovoPessoa("Ana", "Costa", 28)
	fmt.Printf("Original: %+v\n", pessoa5)

	// Copiar a struct (cópia por valor, não referência)
	pessoa5Copia := pessoa5
	pessoa5Copia.Nome = "Beatriz"

	fmt.Printf("Cópia modificada: %+v\n", pessoa5Copia)
	fmt.Printf("Original intacta: %+v\n", pessoa5)
	fmt.Println("Nota: Modificar a cópia NÃO afeta o original")
}

// exemploComTags mostra como usar struct tags para serialização JSON
func exemploComTags() {
	fmt.Println("\n[6] STRUCT TAGS - USO COM JSON")
	fmt.Println("==========================================================")

	user := User{
		Username: "johndoe",
		Password: "senha123",
		Email:    "john@example.com",
	}

	fmt.Printf("Struct original: %+v\n\n", user)

	// Converter para JSON
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Erro ao converter para JSON:", err)
		return
	}

	fmt.Printf("Como JSON: %s\n", jsonBytes)
	fmt.Println("\nNota: As tags `json:\"fieldname\"` definem como serializar para JSON")
}
