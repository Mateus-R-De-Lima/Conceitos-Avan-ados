package exemplogoroutineecontext

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// GOROUTINES E CONCORRÊNCIA EM GO: GUIA ACADÊMICO
//
// Go foi projetado com concorrência em mente desde o início.
// Oferece primitivas simples mas poderosas para escrever programas concorrentes.
//
// Conceitos Fundamentais:
// 1. GOROUTINES: São threads leves gerenciadas pela runtime de Go.
//    - Muito mais leves que threads do SO (milhares podem rodar simultaneamente).
//    - Iniciam com apenas ~2KB de memória vs ~1MB para thread do SO.
//    - Não correspondem 1:1 com threads do SO; Go usa M:N scheduling.
//
// 2. CANAIS: Mecanismo de comunicação entre goroutines.
//    - Unbuffered (sem buffer): bloqueante, até que haja receptor.
//    - Buffered (com buffer): permite enviar N valores sem receptor imediato.
//    - Seguro: thread-safe, projetado para CSP (Communicating Sequential Processes).
//
// 3. CONTEXTO: Controla cancelamento, timeouts e valores compartilhados.
//    - Essencial para operações com timeout, cancelamento coordenado.
//    - Propaga informações em toda a cadeia de chamadas.
//
// 4. SYNC.WAITGROUP: Sincroniza múltiplas goroutines.
//    - Add(): adiciona contador de goroutines a esperar.
//    - Done(): decrementa contador.
//    - Wait(): bloqueia até contador chegar a zero.

func ExemploGoroutineEConcorrencia() {
	fmt.Println("=== GOROUTINES E CONCORRÊNCIA EM GO ===")
	fmt.Println("Demonstrando padrões de concorrência, canais e contexto.\n")

	// Exemplo 1: Sem goroutines (bloqueante/sequencial)
	fmt.Println("1. Execução Sequencial (Sem Goroutines - LENTO):")
	exemploSemGoroutine()

	// Exemplo 2: Com goroutines (concorrente)
	fmt.Println("\n2. Execução Concorrente (Com Goroutines - RÁPIDO):")
	exemploComGoroutine()

	// Exemplo 3: Com Contexto (controle de timeout e cancelamento)
	fmt.Println("\n3. Controle com Contexto (Timeout e Cancelamento):")
	exemploComContext()

	// Exemplo 4: Canais e Comunicação
	fmt.Println("\n4. Canais Unbuffered (Sincronização):")
	exemploCanalUnbuffered()

	// Exemplo 5: Canais com Buffer
	fmt.Println("\n5. Canais Buffered (Produtor-Consumidor):")
	exemploCanalBuffered()

	// Exemplo 6: Select para Multiplexação
	fmt.Println("\n6. Select para Múltiplos Canais:")
	exemploSelect()

	// Exemplo 7: Worker Pool Pattern
	fmt.Println("\n7. Worker Pool Pattern (Padrão Avançado):")
	exemploWorkerPool()

	// Exemplo 8: Fan-Out/Fan-In Pattern
	fmt.Println("\n8. Fan-Out/Fan-In Pattern (Padrão Avançado):")
	exemploFanOutFanIn()
}

func exemploSemGoroutine() {
	// Execução SEQUENCIAL: Cada requisição espera a anterior terminar.
	// Isso é INEFICIENTE para I/O-bound operations (como requisições HTTP).
	// Se cada requisição leva 1 segundo, 10 requisições levarão ~10 segundos.

	start := time.Now()

	// Simulando requisições com sleep em vez de HTTP real
	// (para evitar dependência de rede no exemplo)
	for i := range 5 {
		time.Sleep(1 * time.Second) // Simula latência de rede
		fmt.Printf("Requisição %d concluída\n", i+1)
	}

	totalTime := time.Since(start)
	fmt.Printf("Tempo total (sequencial): %.2f segundos (Ineficiente!)\n", totalTime.Seconds())
}

func exemploComGoroutine() {
	// Execução CONCORRENTE: Múltiplas requisições executam SIMULTANEAMENTE.
	// Isso é EFICIENTE para I/O-bound operations.
	// Se cada requisição leva 1 segundo, 5 requisições concorrentes levarão ~1 segundo!
	//
	// sync.WaitGroup: Garante que main() espere todas as goroutines terminarem.
	// - Add(n): Incrementa contador em n.
	// - Done(): Decrementa contador (defer garante que execute mesmo com erro).
	// - Wait(): Bloqueia até contador ser zero.

	start := time.Now()
	const numRequisicoes = 5

	var wg sync.WaitGroup
	wg.Add(numRequisicoes) // Diz ao WaitGroup: estarei gerando 5 goroutines

	for i := range numRequisicoes {
		go func(id int) { // Goroutine com parameter para evitar closure problems
			defer wg.Done() // Importante: Done() decrementa o contador

			time.Sleep(1 * time.Second) // Simula latência
			fmt.Printf("Goroutine %d concluída\n", id+1)
		}(i) // Passando 'i' como argumento (cópia do valor)
	}

	wg.Wait() // Espera todas as goroutines terminarem
	totalTime := time.Since(start)
	fmt.Printf("Tempo total (concorrente): %.2f segundos (Eficiente!)\n", totalTime.Seconds())
}

func exemploComContext() {
	// CONTEXTO: Controla cancelamento, timeouts e valores em goroutines.
	// context.WithTimeout(): Cria contexto que cancela após duração específica.
	// Isso é essencial para evitar goroutines "penduradas" indefinidamente.

	start := time.Now()
	const numRequisicoes = 5

	var wg sync.WaitGroup
	wg.Add(numRequisicoes)

	// Cria contexto que expira em 2 segundos
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel() // Sempre liberar recursos do contexto

	for i := range numRequisicoes {
		go func(id int, ctx context.Context) {
			defer wg.Done()

			// Aguarda contexto cancelar ou dorme
			select {
			case <-ctx.Done():
				fmt.Printf("Goroutine %d: Contexto expirou (timeout)\n", id+1)
				return
			case <-time.After(3 * time.Second):
				fmt.Printf("Goroutine %d: Operação completada\n", id+1)
			}
		}(i, ctx)
	}

	wg.Wait()
	totalTime := time.Since(start)
	fmt.Printf("Tempo total (com timeout): %.2f segundos\n", totalTime.Seconds())
}

// Exemplo 4: Canais Unbuffered
// Canais unbuffered são síncronos: envio bloqueia até haver receptor.
func exemploCanalUnbuffered() {
	// Canal de inteiros sem buffer
	canalNumeros := make(chan int) // Sem capacidade = unbuffered

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Enviando %d...\n", i)
			canalNumeros <- i // Bloqueia até haver receptor
			fmt.Printf("Valor %d recebido\n", i)
		}
		close(canalNumeros) // Sinaliza "sem mais mensagens"
	}()

	// Receptor (main thread)
	for valor := range canalNumeros {
		fmt.Printf("Main: recebi %d\n", valor)
		time.Sleep(100 * time.Millisecond)
	}
	fmt.Println("Canal fechado - comunicação encerrada")
}

// Exemplo 5: Canais com Buffer
// Canais buffered permitem envios sem receptor imediato (até atingir capacidade).
func exemploCanalBuffered() {
	// Canal com buffer de 2 (aceita 2 valores sem receptor)
	canalBuffer := make(chan string, 2)

	// Enviando sem bloquear (enquanto houver espaço no buffer)
	canalBuffer <- "Primeira mensagem"
	canalBuffer <- "Segunda mensagem"
	fmt.Println("Duas mensagens enfileiradas no buffer")

	// Agora o canal está cheio
	// Próximo envio bloquearia até haver espaço

	// Recebendo (libera espaço no buffer)
	fmt.Println("Primeira recebida:", <-canalBuffer)
	fmt.Println("Segunda recebida:", <-canalBuffer)

	// Agora temos espaço para enviar mais
	canalBuffer <- "Terceira mensagem"
	canalBuffer <- "Quarta mensagem"
	fmt.Println("Mais duas mensagens adicionadas")
}

// Exemplo 6: Select para Multiplexagem
// Select aguarda o primeiro canal ready (similar a select() em C).
func exemploSelect() {
	chanA := make(chan string)
	chanB := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		chanA <- "Mensagem de A"
	}()

	go func() {
		time.Sleep(500 * time.Millisecond)
		chanB <- "Mensagem de B"
	}()

	// Select espera qualquer um dos canais ficar ready
	for i := 0; i < 2; i++ {
		select {
		case msgA := <-chanA:
			fmt.Printf("Recebeu de A: %s\n", msgA)
		case msgB := <-chanB:
			fmt.Printf("Recebeu de B: %s\n", msgB)
		case <-time.After(2 * time.Second):
			fmt.Println("Timeout: nenhum canal respondeu")
		}
	}
}

// Exemplo 7: Worker Pool Pattern
// Padrão comum: N workers processam tarefas de uma fila.
func exemploWorkerPool() {
	const numWorkers = 3
	const numTarefas = 10

	// Canal de tarefas (buffered para não bloquear produtor)
	tarefas := make(chan int, numTarefas)
	resultados := make(chan string, numTarefas)

	var wg sync.WaitGroup

	// Criando workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go func(workerID int) {
			defer wg.Done()
			// Cada worker processa tarefas do canal
			for tarefa := range tarefas {
				resultado := fmt.Sprintf("Worker %d processou tarefa %d", workerID, tarefa)
				resultados <- resultado
				time.Sleep(100 * time.Millisecond) // Simula processamento
			}
		}(w)
	}

	// Produtor envia tarafas
	for tarefa := 1; tarefa <= numTarefas; tarefa++ {
		tarefas <- tarefa
	}
	close(tarefas) // Sinaliza: sem mais tarefas

	// Goroutine para coletar resultados
	go func() {
		wg.Wait()
		close(resultados)
	}()

	// Imprime resultados
	for resultado := range resultados {
		fmt.Println(resultado)
	}
}

// Exemplo 8: Fan-Out/Fan-In Pattern
// Fan-Out: distribui trabalho para múltiplas goroutines.
// Fan-In: coleta resultados de múltiplas goroutines em um canal.
func exemploFanOutFanIn() {
	// Fan-Out: distribuir números para múltiplas goroutines
	numeros := []int{1, 2, 3, 4, 5}
	canalResultados := make([]chan int, len(numeros))

	// Fan-Out: criar goroutine para cada número
	for i, num := range numeros {
		canalResultados[i] = make(chan int)
		go func(ch chan int, n int) {
			// Processa número (exemplo: ao quadrado)
			ch <- n * n
		}(canalResultados[i], num)
	}

	// Fan-In: coletar resultados de todos os canais
	canalMerge := make(chan int)
	var wg sync.WaitGroup

	for _, ch := range canalResultados {
		wg.Add(1)
		go func(c chan int) {
			defer wg.Done()
			canalMerge <- <-c // Recebe e envia para merge
		}(ch)
	}

	// Fechar merge após todas as goroutines terminarem
	go func() {
		wg.Wait()
		close(canalMerge)
	}()

	// Imprimir resultados
	fmt.Println("Resultados (números ao quadrado):")
	for resultado := range canalMerge {
		fmt.Printf("%d ", resultado)
	}
	fmt.Println()
}
