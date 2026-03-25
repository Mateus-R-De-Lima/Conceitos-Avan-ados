package slice

import (
	"fmt"
)

// ExemploMake demonstra o uso da função make com mapas e slices
// em um roteiro didático para quem está começando.
func ExemploMake() {
	fmt.Println("CONCEITOS AVANÇADOS DE GO - MAKE")
	fmt.Println("=")
	fmt.Println()

	fmt.Println("EXEMPLO DE CRIAÇÃO DE SLICE COM MAKE")
	exemploDeComoUsarMake()
	fmt.Println("==========================================================")

	fmt.Println("EXEMPLO DE INICIALIZAÇÃO DE MAPA (LITERAL)")
	exemploDeInicializacaoDeSliceComMake()
	fmt.Println("==========================================================")

	fmt.Println("EXEMPLO DE ATRIBUIÇÃO E LEITURA NO MAPA")
	exemploMakeComAtualizacao()
	fmt.Println("==========================================================")

	fmt.Println("EXEMPLO DE VERIFICAÇÃO DE EXISTÊNCIA DE CHAVE")
	exemploMakeComAtualizacaoEVerificacao()
	fmt.Println("==========================================================")

	fmt.Println("EXEMPLO DE REMOÇÃO DE CHAVES DO MAPA")
	exemploDeComoDeletarDoSlice()
	fmt.Println("==========================================================")

	fmt.Println("EXEMPLO DE LIMPAR TODO O MAPA (sem torná-lo nil)")
	exemploDeComoDeletarTudoSemAlterarACapacidade()
}

// exemploDeComoUsarMake mostra como criar um slice e um mapa com make
func exemploDeComoUsarMake() {
	// make para slice: tipo []int, length 0, capacity 5
	slice := make([]int, 0, 5)
	fmt.Println("Slice criado com make:", slice)
	fmt.Println("Length e Capacity do slice:", len(slice), cap(slice))

	// make para mapa: tipo map[string]string, capacidade inicial (opcional) 100
	m := make(map[string]string, 100)
	fmt.Println("Mapa criado com make é nil?", m == nil)
}

// exemploDeInicializacaoDeSliceComMake demonstra a inicialização literal de mapa
func exemploDeInicializacaoDeSliceComMake() {
	m := map[string]string{
		"nome":      "João",
		"sobrenome": "Silva",
	}
	fmt.Println("Mapa criado com literal:", m)
	fmt.Println("Mapa literal é nil?", m == nil)
}

// exemploMakeComAtualizacao adiciona itens ao mapa e lê um valor
func exemploMakeComAtualizacao() {
	m := make(map[string]string)
	m["nome"] = "João"
	m["sobrenome"] = "Silva"

	valor := m["nome"]
	fmt.Println("Valor da chave 'nome':", valor)
	fmt.Println("Mapa não é nil:", m != nil)
}

// exemploMakeComAtualizacaoEVerificacao guarda o valor e testa existência da chave
func exemploMakeComAtualizacaoEVerificacao() {
	m := make(map[string]string)
	m["nome"] = "João"
	m["sobrenome"] = "Silva"

	valor, ok := m["nome"]
	fmt.Println("Valor da chave 'nome':", valor)
	fmt.Println("Chave 'nome' existe:", ok)

	valorInexistente, okInexistente := m["idade"]
	fmt.Println("Valor da chave 'idade':", valorInexistente)
	fmt.Println("Chave 'idade' existe:", okInexistente)
}

// exemploDeComoDeletarDoSlice mostra como remover um item do mapa
func exemploDeComoDeletarDoSlice() {
	m := make(map[string]string)
	m["nome"] = "João"
	m["sobrenome"] = "Silva"

	fmt.Println("Antes de delete, nome:", m["nome"], "existe:", m["nome"] != "")
	delete(m, "nome")
	fmt.Println("Depois de delete, nome:", m["nome"], "existe:", m["nome"] != "")
	fmt.Println("Chave 'nome' existe (ok idiomático):")
	_, ok := m["nome"]
	fmt.Println(ok)
}

// exemploDeComoDeletarTudoSemAlterarACapacidade usa clear introduzido no Go 1.21
func exemploDeComoDeletarTudoSemAlterarACapacidade() {
	m := make(map[string]string)
	m["nome"] = "João"
	m["sobrenome"] = "Silva"

	clear(m)

	fmt.Println("Mapa após clear:", m)
	fmt.Println("Mapa é nil:", m == nil)
	fmt.Println("Length do mapa após clear:", len(m))
	// Para mapa não existe cap(map), mas claro que a capacidade interna é reutilizável
}
