package exemplointerface

import (
	"fmt"
	"reflect"
)

// ============================================================================
// DEFINIÇÃO DE INTERFACES
// ============================================================================

// Animal define o comportamento básico de um animal
type Animal interface {
	Som() string
}

// Voador define o comportamento de voar
type Voador interface {
	Voar() string
}

// Nadador define o comportamento de nadar
type Nadador interface {
	Nadar() string
}

// Ave define uma ave que pode voar e fazer som
type Ave interface {
	Animal // Incorpora interface Animal
	Voador // Incorpora interface Voador
}

// ============================================================================
// STRUCTS QUE IMPLEMENTAM AS INTERFACES
// ============================================================================

type Cachorro struct {
	Nome string
}

type Gato struct {
	Nome string
}

type Pato struct {
	Nome string
}

type Peixe struct {
	Nome string
}

// ============================================================================
// IMPLEMENTAÇÃO DOS MÉTODOS (SATISFYING INTERFACES)
// ============================================================================

// Implementação para Cachorro
func (c Cachorro) Som() string {
	return "Au Au"
}

// Implementação para Gato
func (g Gato) Som() string {
	return "Miau"
}

// Implementação para Pato (implementa múltiplas interfaces)
func (p Pato) Som() string {
	return "Quack"
}

func (p Pato) Voar() string {
	return "Estou voando como um pato!"
}

func (p Pato) Nadar() string {
	return "Estou nadando como um pato!"
}

// Implementação para Peixe
func (p Peixe) Nadar() string {
	return "Estou nadando como um peixe!"
}

// ============================================================================
// FUNÇÕES QUE USAM INTERFACES
// ============================================================================

// oQueOAnimalDiz demonstra polimorfismo básico
func oQueOAnimalDiz(a Animal) {
	fmt.Printf("O animal diz: %s\n", a.Som())
}

// descreverAnimal usa reflexão para mostrar informações sobre o tipo
func descreverAnimal(a Animal) {
	fmt.Printf("Tipo: %s, Valor: %+v, Som: %s\n", reflect.TypeOf(a), a, a.Som())
}

// ============================================================================
// FUNÇÃO PRINCIPAL COM TODOS OS EXEMPLOS
// ============================================================================

func ExemploInterface() {
	fmt.Println("==========================================================")

	fmt.Println("              EXEMPLOS DE INTERFACES EM GO")
fmt.Println("==========================================================")


	exemploBasico()
	exemploPolimorfismo()
	exemploInterfaceVazia()
	exemploTypeAssertion()
	exemploTypeAssertionPanic()
	exemploTypeSwitch()
	exemploComposicao()
	exemploCasoPratico()

	fmt.Println("==========================================================")

	fmt.Println("              FIM DOS EXEMPLOS")
	fmt.Println("==========================================================")

}

// ============================================================================
// EXEMPLOS INDIVIDUAIS
// ============================================================================

// exemploBasico mostra a definição e uso básico de interface
func exemploBasico() {
	fmt.Println("\n[1] DEFINIÇÃO E USO BÁSICO DE INTERFACE")
	fmt.Println("==========================================================")


	dog := Cachorro{Nome: "Rex"}
	cat := Gato{Nome: "Whiskers"}

	fmt.Printf("Cachorro: %+v\n", dog)
	fmt.Printf("Gato: %+v\n", cat)

	fmt.Println("\nChamando método Som() diretamente:")
	fmt.Printf("Cachorro faz: %s\n", dog.Som())
	fmt.Printf("Gato faz: %s\n", cat.Som())

	fmt.Println("\nNota: Qualquer struct que implementa Som() string satisfaz a interface Animal")
}

// exemploPolimorfismo demonstra como interfaces permitem polimorfismo
func exemploPolimorfismo() {
	fmt.Println("\n[2] POLIMORFISMO COM INTERFACES")
	fmt.Println("==========================================================")


	// Slice de Animal (interface) pode conter diferentes tipos
	animais := []Animal{
		Cachorro{Nome: "Rex"},
		Gato{Nome: "Whiskers"},
		Pato{Nome: "Donald"},
	}

	fmt.Println("Iterando sobre diferentes animais:")
	for i, animal := range animais {
		fmt.Printf("%d. ", i+1)
		oQueOAnimalDiz(animal)
	}

	fmt.Println("\nNota: Mesmo loop trata tipos diferentes de forma uniforme")
}

// exemploInterfaceVazia mostra o uso da interface{} (interface vazia)
func exemploInterfaceVazia() {
	fmt.Println("\n[3] INTERFACE VAZIA (interface{})")
	fmt.Println("==========================================================")


	// interface{} pode armazenar qualquer tipo
	var coisa interface{}

	coisa = "Olá, mundo!"
	fmt.Printf("String: %v (tipo: %T)\n", coisa, coisa)

	coisa = 42
	fmt.Printf("Int: %v (tipo: %T)\n", coisa, coisa)

	coisa = true
	fmt.Printf("Bool: %v (tipo: %T)\n", coisa, coisa)

	coisa = Cachorro{Nome: "Buddy"}
	fmt.Printf("Struct: %+v (tipo: %T)\n", coisa, coisa)

	fmt.Println("\nNota: interface{} é como 'any' em outras linguagens, aceita qualquer valor")
}

// exemploTypeAssertion demonstra como fazer type assertion segura
func exemploTypeAssertion() {
	fmt.Println("\n[4] TYPE ASSERTION (Verificação Segura de Tipo)")
	fmt.Println("==========================================================")


	var animal Animal = Cachorro{Nome: "Max"}

	// Type assertion segura (retorna valor e bool indicando sucesso)
	if dog, ok := animal.(Cachorro); ok {
		fmt.Printf("✅ É um cachorro! Nome: %s, Som: %s\n", dog.Nome, dog.Som())
	} else {
		fmt.Println("❌ Não é um cachorro")
	}

	// Testando com tipo errado
	if cat, ok := animal.(Gato); ok {
		fmt.Printf("✅ É um gato! Nome: %s\n", cat.Nome)
	} else {
		fmt.Println("❌ Não é um gato (como esperado)")
	}

	fmt.Println("\nNota: Type assertion segura usa 'valor, ok := interface.(Tipo)'")
}

// exemploTypeAssertionPanic demonstra o PANIC que pode ocorrer com type assertion
func exemploTypeAssertionPanic() {
	fmt.Println("\n[5] TYPE ASSERTION COM PANIC (PERIGO!)")
	fmt.Println("==========================================================")


	var animal Animal = Cachorro{Nome: "Rex"}

	fmt.Println("Tentando fazer type assertion sem verificar...")

	// ⚠️  PERIGOSO: Type assertion sem verificação pode causar PANIC!
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("🚨 PANIC CAPTURADO: %v\n", r)
			fmt.Println("Nota: Sempre use a forma segura 'valor, ok := interface.(Tipo)'")
		}
	}()

	// Esta linha causaria panic se o tipo estivesse errado
	cat := animal.(Gato) // Panic aqui!
	fmt.Printf("Gato: %+v\n", cat)
}

// exemploTypeSwitch demonstra como usar type switch para diferentes tipos
func exemploTypeSwitch() {
	fmt.Println("\n[6] TYPE SWITCH (Tratamento de Múltiplos Tipos)")
	fmt.Println("==========================================================")


	animais := []Animal{
		Cachorro{Nome: "Rex"},
		Gato{Nome: "Mia"},
		Pato{Nome: "Donald"},
	}

	for _, animal := range animais {
		switch a := animal.(type) {
		case Cachorro:
			fmt.Printf("🐕 É um cachorro chamado %s: %s\n", a.Nome, a.Som())
		case Gato:
			fmt.Printf("🐱 É um gato chamado %s: %s\n", a.Nome, a.Som())
		case Pato:
			fmt.Printf("🦆 É um pato chamado %s: %s\n", a.Nome, a.Som())
		default:
			fmt.Printf("❓ Animal desconhecido: %s\n", a.Som())
		}
	}

	fmt.Println("\nNota: Type switch é útil quando você precisa tratar tipos específicos")
}

// exemploComposicao mostra como combinar interfaces
func exemploComposicao() {
	fmt.Println("\n[7] COMPOSIÇÃO DE INTERFACES")
	fmt.Println("==========================================================")


	// Pato implementa Animal, Voador e Nadador
	duck := Pato{Nome: "Donald"}

	// Como Animal
	var animal Animal = duck
	fmt.Printf("Como Animal: %s\n", animal.Som())

	// Como Voador
	var voador Voador = duck
	fmt.Printf("Como Voador: %s\n", voador.Voar())

	// Como Nadador
	var nadador Nadador = duck
	fmt.Printf("Como Nadador: %s\n", nadador.Nadar())

	// Como Ave (interface composta)
	var ave Ave = duck
	fmt.Printf("Como Ave (Animal+Voador): Som=%s, Voo=%s\n", ave.Som(), ave.Voar())

	fmt.Println("\nNota: Interfaces podem ser compostas incorporando outras interfaces")
}

// Tipos para caso prático
type Pagamento interface {
	Processar(valor float64) error
	Validar() error
}

type CartaoCredito struct {
	Numero   string
	Validade string
	CVV      string
}

type Boleto struct {
	Codigo     string
	Vencimento string
}

// Implementações para CartaoCredito
func (c CartaoCredito) Processar(valor float64) error {
	fmt.Printf("💳 Processando R$%.2f no cartão %s\n", valor, c.Numero[len(c.Numero)-4:])
	return nil
}

func (c CartaoCredito) Validar() error {
	if len(c.Numero) < 16 {
		return fmt.Errorf("número do cartão inválido")
	}
	return nil
}

// Implementações para Boleto
func (b Boleto) Processar(valor float64) error {
	fmt.Printf("📄 Gerando boleto no valor de R$%.2f\n", valor)
	return nil
}

func (b Boleto) Validar() error {
	if b.Codigo == "" {
		return fmt.Errorf("código do boleto vazio")
	}
	return nil
}

// exemploCasoPratico demonstra um caso prático com contexto real
func exemploCasoPratico() {
	fmt.Println("\n[8] CASO PRÁTICO - SISTEMA DE PAGAMENTOS")
	fmt.Println("==========================================================")


	// Função que processa qualquer forma de pagamento
	processarCompra := func(p Pagamento, valor float64) error {
		if err := p.Validar(); err != nil {
			return fmt.Errorf("validação falhou: %v", err)
		}
		return p.Processar(valor)
	}

	// Testando diferentes formas de pagamento
	pagamentos := []Pagamento{
		CartaoCredito{Numero: "1234567890123456", Validade: "12/25", CVV: "123"},
		Boleto{Codigo: "12345678901234567890123456789012345678901234", Vencimento: "2024-12-31"},
	}

	for i, pagamento := range pagamentos {
		fmt.Printf("\nProcessando pagamento %d:\n", i+1)
		if err := processarCompra(pagamento, 99.90); err != nil {
			fmt.Printf("❌ Erro: %v\n", err)
		} else {
			fmt.Println("✅ Pagamento processado com sucesso!")
		}
	}

	fmt.Println("\nNota: Interfaces permitem criar APIs flexíveis e extensíveis")
}
