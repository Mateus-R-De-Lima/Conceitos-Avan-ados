package exemplogenerics

import "fmt"

// Generics em Go: Introdução Acadêmica
//
// Os generics foram introduzidos em Go 1.18 para permitir que funções, structs e interfaces
// trabalhem com tipos de dados arbitrários de forma type-safe, sem recorrer a interface{} ou reflexão.
// Isso promove código reutilizável, legível e eficiente.
//
// Sintaxe Básica:
// - Parâmetros de tipo são definidos entre colchetes: [T any]
// - 'any' é um alias para interface{} (disponível desde Go 1.18)
// - Constraints podem ser aplicados usando interfaces para restringir tipos permitidos.
//
// Benefícios:
// - Type safety em tempo de compilação.
// - Evita type assertions e conversões manuais.
// - Melhora a performance ao evitar boxing/unboxing.
//
// Limitações:
// - Não suporta specialization (como templates em C++).
// - Constraints são baseados em interfaces.

// Exemplo 1: Struct Genérica Simples
// Uma struct que pode armazenar qualquer tipo de valor.
type Caixa[T any] struct {
	Valor T
}

// Exemplo 2: Função Genérica Básica
// Função que aceita qualquer tipo e imprime seu valor.
func exemploMetodoGenerico[T any](valor T) {
	fmt.Printf("Valor recebido (tipo %T): %v\n", valor, valor)
}

// Exemplo 3: Constraints com Interfaces
// Usando constraints para restringir tipos aceitos.
// Aqui, apenas tipos que implementam a interface 'comparable' são permitidos.
// 'comparable' inclui tipos que podem ser comparados com == e !=.
func encontrarMaior[T comparable](a, b T) T {
	if a == b {
		return a // Retorna qualquer um se iguais
	}
	// Nota: Para tipos numéricos, precisaríamos de ordered, mas vamos simplificar.
	// Em Go, para comparações > <, usamos constraints como constraints.Ordered.
	return a // Exemplo simplificado
}

// Exemplo 4: Função Genérica com Múltiplos Tipos
// Função que trabalha com dois tipos genéricos.
func trocarValores[T, U any](a T, b U) (U, T) {
	return b, a
}

// Exemplo 5: Slice Genérico
// Função que filtra elementos de um slice baseado em uma condição genérica.
// Usando uma função de callback para flexibilidade.
func filtrar[T any](slice []T, condicao func(T) bool) []T {
	var resultado []T
	for _, item := range slice {
		if condicao(item) {
			resultado = append(resultado, item)
		}
	}
	return resultado
}

// Exemplo 6: Mapa Genérico
// Struct que representa um mapa genérico simples.
type Mapa[K comparable, V any] struct {
	dados map[K]V
}

// Método para adicionar ao mapa.
func (m *Mapa[K, V]) Adicionar(chave K, valor V) {
	if m.dados == nil {
		m.dados = make(map[K]V)
	}
	m.dados[chave] = valor
}

// Método para obter valor do mapa.
func (m *Mapa[K, V]) Obter(chave K) (V, bool) {
	valor, existe := m.dados[chave]
	return valor, existe
}

// Exemplo 7: Interfaces como Constraints Práticos
// Usando uma interface customizada como constraint.
// Isso permite que apenas tipos que implementam a interface sejam aceitos.
type Imprimivel interface {
	String() string
}

// Função genérica que aceita apenas tipos que implementam Imprimivel.
func Imprimir[T Imprimivel](valor T) {
	fmt.Println("Imprimindo:", valor.String())
}

// Exemplo 8: Conceito Avançado - Generics com Interfaces para Polimorfismo
// Interfaces permitem polimorfismo: tratar diferentes tipos de forma uniforme.
// Aqui, uma função que processa uma lista de itens que têm um método comum.
type Processavel interface {
	Processar() string
}

// Structs que implementam Processavel.
type Tarefa struct {
	Nome string
}

func (t Tarefa) Processar() string {
	return fmt.Sprintf("Processando tarefa: %s", t.Nome)
}

type Arquivo struct {
	Nome string
}

func (a Arquivo) Processar() string {
	return fmt.Sprintf("Processando arquivo: %s", a.Nome)
}

// Função genérica que processa uma slice de qualquer tipo que implementa Processavel.
func ProcessarItens[T Processavel](itens []T) {
	for _, item := range itens {
		fmt.Println(item.Processar())
	}
}

// Exemplo 9: Conceito Avançado - Constraints com Múltiplas Interfaces
// Combinando interfaces para constraints mais específicas.
// Usando uma interface que combina várias.
type Numerico interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

// Função que soma apenas tipos numéricos.
func Somar[T Numerico](a, b T) T {
	return a + b
}

// Exemplo 10: Prático - Função Genérica para Encontrar em Slice
// Usando generics para buscar em slices de qualquer tipo comparável.
func Encontrar[T comparable](slice []T, alvo T) (int, bool) {
	for i, v := range slice {
		if v == alvo {
			return i, true
		}
	}
	return -1, false
}

// Exemplo 11: Avançado - Generics com Type Assertions (Cuidado: Perde Type Safety)
// Embora generics promovam type safety, às vezes precisamos de type assertions.
// Isso é avançado e deve ser usado com cautela.
func ConverterParaString[T any](valor T) string {
	switch v := any(valor).(type) {
	case string:
		return v
	case int:
		return fmt.Sprintf("%d", v)
	case float64:
		return fmt.Sprintf("%.2f", v)
	default:
		return fmt.Sprintf("%v", v)
	}
}

// Exemplo 12: Tipo Produto que implementa Imprimivel
type Produto struct {
	Nome  string
	Preco float64
}

func (p Produto) String() string {
	return fmt.Sprintf("%s - R$ %.2f", p.Nome, p.Preco)
}

// Atualizando a Função Principal para Incluir Novos Exemplos
func ExemploGenerics() {
	fmt.Println("=== Generics em Go: Exemplos ===\n")

	// Exemplos existentes (1-6) permanecem os mesmos...

	// 7. Interfaces como Constraints
	fmt.Println("\n7. Interfaces como Constraints Práticos:")

	produto := Produto{Nome: "Livro", Preco: 29.99}
	Imprimir(produto)

	// 8. Polimorfismo com Interfaces
	fmt.Println("\n8. Polimorfismo com Interfaces Genéricas:")
	tarefas := []Tarefa{
		{Nome: "Limpar quarto"},
		{Nome: "Fazer lição"},
	}
	arquivos := []Arquivo{
		{Nome: "foto.jpg"},
		{Nome: "documento.pdf"},
	}
	ProcessarItens(tarefas)
	ProcessarItens(arquivos)

	// 9. Constraints com Múltiplas Interfaces
	fmt.Println("\n9. Constraints com Tipos Numéricos:")
	resultadoInt := Somar(10, 20)
	fmt.Printf("Soma de ints: %d\n", resultadoInt)
	resultadoFloat := Somar(3.5, 2.1)
	fmt.Printf("Soma de floats: %.2f\n", resultadoFloat)

	// 10. Busca em Slice Genérica
	fmt.Println("\n10. Busca em Slice Genérica:")
	numeros := []int{1, 2, 3, 4, 5}
	if indice, encontrado := Encontrar(numeros, 3); encontrado {
		fmt.Printf("Número 3 encontrado no índice: %d\n", indice)
	}
	palavras := []string{"gato", "cachorro", "pássaro"}
	if indice, encontrado := Encontrar(palavras, "cachorro"); encontrado {
		fmt.Printf("Palavra 'cachorro' encontrada no índice: %d\n", indice)
	}

	// 11. Type Assertions em Generics
	fmt.Println("\n11. Type Assertions (Avançado - Use com Cuidado):")
	fmt.Println("Convertendo int:", ConverterParaString(42))
	fmt.Println("Convertendo string:", ConverterParaString("teste"))
	fmt.Println("Convertendo float:", ConverterParaString(3.14159))
	fmt.Println("Convertendo struct:", ConverterParaString(produto))

	fmt.Println("\n=== Fim dos Exemplos de Generics ===")
}
