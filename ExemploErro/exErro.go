package exemploerro

import (
	"errors"
	"fmt"
)

// Em Go, os erros são tratados de forma diferente das exceções em outras linguagens.
// A interface error é um tipo built-in que representa um erro.
// Qualquer tipo que implemente o método Error() string é considerado um erro.
// Erros são valores retornados pelas funções, não lançados como exceções.
// Isso torna o tratamento de erros explícito e previsível.

// Definindo um tipo de erro customizado.
// MeuErro implementa a interface error ao definir o método Error().
type MeuErro struct {
	Mensagem string
}

// Método Error() que satisfaz a interface error.
func (e *MeuErro) Error() string {
	return e.Mensagem
}

// Função principal que demonstra o tratamento de erros.
func ExemploErro() {
	fmt.Println("=== Tratamento de Erros em Go ===")
	fmt.Println("Erros em Go são valores retornados pelas funções.")
	fmt.Println("Devemos sempre verificar se o erro é nil antes de prosseguir.\n")

	// Exemplo 1: Função que retorna erro
	fmt.Println("1. Exemplo de função que retorna erro:")
	resultado, err := raizQuadrada(-4)
	if err != nil {
		fmt.Println("Erro:", err.Error())
	} else {
		fmt.Printf("Resultado da raiz quadrada: %.2f\n", resultado)
	}

	resultado, err = raizQuadrada(16)
	if err != nil {
		fmt.Println("Erro:", err.Error())
	} else {
		fmt.Printf("Resultado da raiz quadrada: %.2f\n", resultado)
	}

	// Exemplo 2: Divisão com verificação de erro
	fmt.Println("\n2. Exemplo de divisão:")
	quociente, err := dividir(10, 0)
	if err != nil {
		fmt.Println("Erro:", err.Error())
	} else {
		fmt.Printf("Resultado da divisão: %d\n", quociente)
	}

	quociente, err = dividir(10, 2)
	if err != nil {
		fmt.Println("Erro:", err.Error())
	} else {
		fmt.Printf("Resultado da divisão: %d\n", quociente)
	}

	// Exemplo 3: Validação de tipos de erro com errors.Is
	fmt.Println("\n3. Validação de tipos de erro com errors.Is:")
	exemploDeValidacaoDeErro()

	// Exemplo 4: Validação de tipos específicos com errors.As
	fmt.Println("\n4. Validação de tipos específicos com errors.As:")
	exemploDeValidacaoDeErro2()

	// Exemplo 5: Panic e Recover
	fmt.Println("\n5. Panic e Recover:")
	fmt.Println("Panic é usado para situações irrecuperáveis, como bugs de programação.")
	fmt.Println("Recover pode ser usado em defer para capturar panics.")
	exemploPanicRecover()
}

// Função que calcula uma "raiz quadrada" simples (apenas exemplo, não precisa ser precisa).
// Retorna erro se o número for negativo.
func raizQuadrada(x float64) (float64, error) {
	if x < 0 {
		// Retornando um erro customizado
		return 0, &MeuErro{Mensagem: "Não pode calcular a raiz quadrada de um número negativo!"}
	}
	// Implementação simplificada (não é a raiz real, apenas para exemplo)
	return x / 2, nil // Retorna nil para indicar ausência de erro
}

// Função de divisão que retorna erro se o divisor for zero.
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, &MeuErro{Mensagem: "Não pode dividir por zero!"}
	}
	return a / b, nil
}

// Definindo um erro global usando errors.New.
// Isso é útil para erros comuns que podem ser reutilizados.
var ErrNotFound = errors.New("not found")

// Função que demonstra o uso de errors.Is para verificar tipos de erro.
func exemploDeValidacaoDeErro() {
	err := funcaoQueRetornaErro()

	// errors.Is verifica se o erro é do tipo especificado ou envolve ele.
	if errors.Is(err, ErrNotFound) && err != nil {
		fmt.Println("Erro identificado como 'Not Found'")
	} else {
		fmt.Println("Nenhum erro ou erro diferente encontrado")
	}
}

// Função que simula o retorno de um erro.
func funcaoQueRetornaErro() error {
	// Simulando um erro
	return ErrNotFound
}

// Função que demonstra o uso de errors.As para verificar tipos específicos de erro.
func exemploDeValidacaoDeErro2() {
	err := funcaoQueRetornaErroEspecifico()

	// errors.As verifica se o erro pode ser atribuído ao tipo especificado.
	var meuErro *MeuErro
	if err != nil && errors.As(err, &meuErro) {
		fmt.Println("Erro do tipo MeuErro:", meuErro.Mensagem)
	} else {
		fmt.Println("Erro não é do tipo MeuErro")
	}
}

// Função que retorna um erro específico.
func funcaoQueRetornaErroEspecifico() error {
	return &MeuErro{Mensagem: "Este é um erro específico!"}
}

// Função que demonstra panic e recover.
// Panic interrompe a execução normal e começa a desenrolar a pilha de chamadas.
// Recover pode capturar o panic em uma função defer.
func exemploPanicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recuperado de panic:", r)
		}
	}()

	fmt.Println("Chamando função que causa panic...")
	funcaoQueCausaPanic()
	fmt.Println("Esta linha não será executada se houver panic.")
}

// Função que causa panic (apenas para exemplo).
func funcaoQueCausaPanic() {
	panic("Algo deu errado! Este é um panic.")
}
