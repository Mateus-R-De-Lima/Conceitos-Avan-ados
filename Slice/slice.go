package slice

import "fmt"

// ExemploSlice demonstra os conceitos básicos de slice e arrays em Go.
// O objetivo é deixar o código legível para um júnior ou alguém que volte daqui a 1 ano.
func ExemploSlice() {
	fmt.Println("CONCEITOS AVANÇADOS DE GO - SLICE")
	fmt.Println("=")

	fmt.Println("EXEMPLO 1: Slice simples literal")
	exemploSliceSimples()
	fmt.Println("--------------------------------------------------")

	fmt.Println("EXEMPLO 2: Slice referenciando um array")
	exemploSliceComReferenciaArray()
	fmt.Println("--------------------------------------------------")

	fmt.Println("EXEMPLO 3: Slice sem array declarado explicitamente")
	exemploSliceSemDeclararUmArray()
	fmt.Println("--------------------------------------------------")

	fmt.Println("EXEMPLO 4: Uso de length e capacity")
	exemploSliceLengthECapacity()
	fmt.Println("--------------------------------------------------")

	fmt.Println("EXEMPLO 5: Slice sem alocação (nil slice)")
	exemploSliceSemAlocacao()
	fmt.Println("--------------------------------------------------")

	fmt.Println("EXEMPLO 6: Slice literal vazio")
	exemploSliceLiteral()
	fmt.Println("--------------------------------------------------")

	fmt.Println("EXEMPLO 7: Construção incremental de slice")
	exemploDeComoSalvarNoSlice()
	fmt.Println("--------------------------------------------------")

	fmt.Println("EXEMPLO 8: Construção otimizada de slice com make")
	exemploDeComoSalvarNoSliceOtimizado()
}

func exemploSliceSimples() {
	// O slice é uma visão dinâmica sobre um array interno.
	numeros := []int{1, 2, 3, 4, 5}
	// Imprime o slice completo e mostra que contém 5 valores.
	fmt.Println("Slice simples:", numeros)
}

func exemploSliceComReferenciaArray() {
	// Array fixo. Seu tamanho é imutável (5).
	arrayNumeros := [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array original:", arrayNumeros)

	// Slice que aponta para parte do array.
	slice := arrayNumeros[1:4] // indices 1,2,3
	fmt.Println("Slice criado a partir do array:", slice)

	// Alterar o array altera o slice (mesma memória subjacente).
	arrayNumeros[2] = 100
	fmt.Println("Array modificado:", arrayNumeros)
	fmt.Println("Slice após modificação do array:", slice)
}

func exemploSliceSemDeclararUmArray() {
	// Declaração direta de slice; o array interno é gerenciado pelo runtime.
	slice := []string{"Go", "Python", "JavaScript"}
	fmt.Println("Slice direto sem array explícito:", slice)
}

func exemploSliceLengthECapacity() {
	// Slice criado por literal.
	slice := []int{1, 2, 3, 4, 5}

	fmt.Println("Slice:", slice)
	fmt.Println("Length (número de elementos):", len(slice))
	fmt.Println("Capacity (capacidade reservada):", cap(slice))

	// Capacity pode ser maior que len, mantendo espaço livre para crescer sem nova alocação.
}

func exemploSliceSemAlocacao() {
	// Var slice cria um slice nil (não aponta para array).
	var slice []int
	fmt.Println("Slice nil:", slice)
	fmt.Println("Length:", len(slice), "=> 0")
	fmt.Println("Capacity:", cap(slice), "=> 0")
	fmt.Println("É nil?", slice == nil)
}

func exemploSliceLiteral() {
	// Literal vazio. Slice não é nil, mas tem len=0 e cap=0.
	slice := []string{}
	fmt.Println("Slice literal vazio:", slice)
	fmt.Println("É nil?", slice == nil)
	fmt.Println("Length:", len(slice), "Capacity:", cap(slice))
}

// MOCK DE UM BANCO DE DADOS DE FILMES
var filmesNoBd = []string{
	"Matrix",
	"Inception",
	"Interstellar",
	"The Dark Knight",
	"Pulp Fiction",
	"Fight Club",
	"Forrest Gump",
	"The Shawshank Redemption",
	"The Godfather",
	"The Lord of the Rings: The Return of the King",
	"The Matrix Reloaded",
	"The Matrix Revolutions",
	"The Dark Knight Rises",
	"The Prestige",
}

func exemploDeComoSalvarNoSlice() {
	// Simulando IDs vindos de uma API e construindo um slice com append.
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var filmes []string // começa vazio, sem capacidade reservada.

	for _, id := range resultsFromApi {
		filme := filmesNoBd[id]
		fmt.Printf("antes: len=%d cap=%d\n", len(filmes), cap(filmes))
		filmes = append(filmes, filme)
		fmt.Printf("depois: len=%d cap=%d\n", len(filmes), cap(filmes))
	}

	fmt.Println("Filmes encontrados:", filmes)
}

func exemploDeComoSalvarNoSliceOtimizado() {
	// Pré-aloca o slice; evita realocações internas durante append.
	resultsFromApi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	filmes := make([]string, 0, len(resultsFromApi))

	for _, id := range resultsFromApi {
		filme := filmesNoBd[id]
		fmt.Printf("antes: len=%d cap=%d\n", len(filmes), cap(filmes))
		filmes = append(filmes, filme)
		fmt.Printf("depois: len=%d cap=%d\n", len(filmes), cap(filmes))
	}

	fmt.Println("Filmes encontrados:", filmes)
}
