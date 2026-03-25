// ============================================================================
// PROJETO: Entendendo Ponteiros em Go
// Objetivo: Demonstrar a diferença entre passar valores e referências
// ============================================================================

package ponteiros

import "fmt"

func ExemploPonteiro() {
	fmt.Println("=")
	fmt.Println("CONCEITOS AVANÇADOS DE GO - PONTEIROS")
	fmt.Println("=")
	fmt.Println()

	// Exemplo 1: SEM Ponteiro (Cópia do Valor)
	fmt.Println("📌 EXEMPLO 1: Passando um Valor SEM Ponteiro")
	fmt.Println("-")

	numero := 10
	fmt.Printf("Valor inicial de 'numero': %d\n", numero)
	fmt.Printf("Endereço de memória: %p\n", &numero)
	fmt.Println()

	passarSemPonteiro(numero)

	fmt.Printf("Valor após chamar passarSemPonteiro: %d\n", numero)
	fmt.Println("✓ O valor NÃO foi alterado porque foi passada uma CÓPIA")
	fmt.Println()
	fmt.Println()

	// Exemplo 2: COM Ponteiro (Referência ao Original)
	fmt.Println("📌 EXEMPLO 2: Passando uma Referência COM Ponteiro")
	fmt.Println("-")

	fmt.Printf("Valor inicial de 'numero': %d\n", numero)
	fmt.Printf("Endereço de memória: %p\n", &numero)
	fmt.Println()

	passarComPonteiro(&numero)

	fmt.Printf("Valor após chamar passarComPonteiro: %d\n", numero)
	fmt.Println("✓ O valor FOI alterado porque foi passada uma REFERÊNCIA")
	fmt.Println()
	fmt.Println()

	// Exemplo 3: Entendendo o operador &
	fmt.Println("📌 EXEMPLO 3: Operador & (Address-of)")
	fmt.Println("-")
	explicarOperadorAmpersand()
	fmt.Println()
	fmt.Println()

	// Exemplo 4: Entendendo o operador *
	fmt.Println("📌 EXEMPLO 4: Operador * (Dereference)")
	fmt.Println("-")
	explicarOperadorAsterisco()
	fmt.Println()
}

// ============================================================================
// FUNÇÃO 1: Passagem SEM Ponteiro
// ============================================================================
// Esta função recebe uma CÓPIA do valor.
// Qualquer modificação feita aqui NÃO afeta a variável original.
func passarSemPonteiro(valor int) {
	fmt.Println("Dentro de passarSemPonteiro():")
	fmt.Printf("  Valor recebido: %d\n", valor)
	fmt.Printf("  Endereço de memória (cópia): %p\n", &valor)

	valor = 100
	fmt.Printf("  Valor modificado para: %d\n", valor)
	fmt.Println("  ⚠️ Esta mudança é apenas LOCAL, não afeta a variável original")
}

// ============================================================================
// FUNÇÃO 2: Passagem COM Ponteiro
// ============================================================================
// Esta função recebe um PONTEIRO (referência) para a variável original.
// Modificações feitas aqui AFETAM a variável original.
// O símbolo *int significa "ponteiro para um inteiro"
func passarComPonteiro(ponteiro *int) {
	fmt.Println("Dentro de passarComPonteiro():")
	fmt.Printf("  Endereço recebido (ponteiro): %p\n", ponteiro)
	fmt.Printf("  Valor armazenado nesse endereço: %d\n", *ponteiro)

	// *ponteiro significa: "acesse o valor no endereço apontado"
	*ponteiro = 100
	fmt.Printf("  Valor modificado através do ponteiro: %d\n", *ponteiro)
	fmt.Println("  ✓ Esta mudança AFETA a variável original")
}

// ============================================================================
// FUNÇÃO 3: Explicar o operador &
// ============================================================================
// O operador & retorna o ENDEREÇO DE MEMÓRIA da variável
func explicarOperadorAmpersand() {
	mensagem := "Go é incrível!"

	fmt.Println("Usando o operador & para obter o endereço de memória:")
	fmt.Printf("  Variável: mensagem = \"%s\"\n", mensagem)
	fmt.Printf("  Endereço: &mensagem = %p (endereço em hexadecimal)\n", &mensagem)
	fmt.Println()
	fmt.Println("Lógica:")
	fmt.Println("  - 'mensagem' = o VALOR armazenado")
	fmt.Println("  - '&mensagem' = o ENDEREÇO onde esse valor está guardado")
	fmt.Println("  - Útil para criar ponteiros!")
}

// ============================================================================
// FUNÇÃO 4: Explicar o operador *
// ============================================================================
// O operador * ACESSA o valor armazenado em um endereço de memória
func explicarOperadorAsterisco() {
	numero := 42
	ponteiro := &numero // Cria um ponteiro para 'numero'

	fmt.Println("Usando o operador * para acessar um valor através do ponteiro:")
	fmt.Printf("  Variável: numero = %d\n", numero)
	fmt.Printf("  Ponteiro: ponteiro = %p (endereço)\n", ponteiro)
	fmt.Printf("  Dereference: *ponteiro = %d (valor original)\n", *ponteiro)
	fmt.Println()
	fmt.Println("Lógica:")
	fmt.Println("  - 'ponteiro' = armazena um endereço")
	fmt.Println("  - '*ponteiro' = acessa o valor naquele endereço")
	fmt.Println("  - Útil para ler e modificar valores indiretamente!")
}
