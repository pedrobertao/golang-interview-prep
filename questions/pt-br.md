## 🤖 Perguntas gerais 🤖

#### 1. O que é Go? Explique suas principais características.

> **Resposta:**
> Go, também conhecido como Golang, é uma linguagem de programação de código aberto desenvolvida pelo Google. É uma linguagem compilada, de tipagem estática, conhecida por sua simplicidade, eficiência e suporte à concorrência. Principais características incluem coleta de lixo, biblioteca padrão rica, goroutines para concorrência e ferramentas integradas para formatação de código e testes.

#### 2. Como você declara uma variável em Go?

> **Resposta:**
> Você pode declarar uma variável usando a palavra-chave var ou o operador de declaração curta :=. Por exemplo:

```
var x int
x = 10 // ou
y := 20
```

#### 3. O que são Goroutines e como elas diferem de threads?

> **Resposta:**
> Goroutines são threads leves gerenciadas pelo runtime do Go. Elas são mais eficientes do que threads do sistema operacional, permitindo que milhares delas sejam executadas simultaneamente. Goroutines usam uma técnica de multiplexação para alcançar concorrência, enquanto threads são gerenciadas pelo sistema operacional e consomem mais recursos.

#### 4. Qual é o propósito da instrução defer em Go?

> **Resposta:**
> A instrução defer é usada para garantir que uma chamada de função seja executada mais tarde na execução do programa, geralmente para fins de limpeza. Funções adiadas são executadas em ordem LIFO (Last In, First Out) logo antes da função circundante retornar.

#### 5. Explique como Go lida com erros.

> **Resposta:**
> Go lida com erros usando um tipo de erro simples e explícito. Funções que podem resultar em um erro retornam um valor de erro adicional. O tratamento de erros é feito verificando esse valor retornado. A biblioteca padrão fornece o pacote errors para criar e manipular erros.

#### 6. O que é uma interface em Go e como é usada?

> **Resposta:**
> Uma interface em Go é um tipo que especifica um conjunto de assinaturas de métodos. Um tipo implementa uma interface implementando seus métodos. Interfaces são usadas para definir comportamento e alcançar polimorfismo. Por exemplo:
```
type Reader interface {
    Read(p []byte) (n int, err error)
}
```

#### 7. Descreva o gerenciamento de pacotes do Go e como as dependências são tratadas.

> **Resposta:**
> Go usa módulos para gerenciar dependências, definidos por um arquivo go.mod. O comando go fornece várias ferramentas para gerenciar dependências, como go get para baixar pacotes e go mod para gerenciar módulos e dependências.

#### 8. O que são canais em Go e como são usados?

> **Resposta:**
> Canais são conduítes tipados através dos quais você pode enviar e receber valores com o operador de canal, <-. Eles são usados para sincronizar goroutines e comunicar dados entre elas. Canais podem ser bufferizados ou não bufferizados.

#### 9. Explique como o coletor de lixo funciona em Go.

> **Resposta:**
> Go usa um coletor de lixo concorrente, do tipo mark-and-sweep. Ele é executado em segundo plano, identificando e coletando objetos de memória não utilizados, enquanto minimiza as pausas na execução do programa. O coletor de lixo usa múltiplas estratégias para equilibrar entre eficiência de memória e latência de aplicação.

#### 10. Quais são as melhores práticas para escrever programas concorrentes em Go?

> **Resposta:**
> *As melhores práticas incluem:*
> - Evitar compartilhamento de memória comunicando-se através de canais.
> - Usar goroutines de maneira eficiente e evitar criar muitas.
> - Lidar adequadamente com a sincronização usando canais, mutexes e outros primitivos de sincronização.
> - Evitar condições de corrida protegendo o acesso a dados compartilhados.

#### 11. Como Go implementa slices e quais são os mecanismos subjacentes?

> **Resposta:**
> Slices em Go são visões flexíveis e de tamanho dinâmico dos elementos de um array. Elas são descritas por uma estrutura contendo um ponteiro para o array, o comprimento e a capacidade. Operações em slices envolvem a manipulação desses campos e podem envolver realocação de arrays subjacentes ao redimensionar.

#### 12. O que é uma instrução select em Go e como é usada?

> **Resposta:**
> A instrução select permite que uma goroutine espere em múltiplas operações de comunicação. Ela bloqueia até que um de seus casos possa prosseguir, então executa esse caso. É usada para lidar com múltiplos canais e evitar bloqueio em um único canal.

## ❗ Perguntas difíceis ❗

#### 1. Qual é a saída do seguinte programa Go e por quê?

```
import "fmt"
package main
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

> **Resposta:**
> hello
> world

> **Explicação:**
> A instrução defer adia a execução de fmt.Println("world") até que a função circundante (main) retorne, então "hello" é impresso primeiro, seguido por "world".

#### 2. Como você previne uma condição de corrida no seguinte código Go?
```
package main
import (
    "fmt"
    "time"
)
func main() {
    var x int
    go func() {
        x++
    }()
    time.Sleep(1 * time.Second)
    fmt.Println(x)
}
```

> **Resposta:**
> Para prevenir uma condição de corrida, você pode usar um primitivo de sincronização como um mutex ou um canal. Aqui está uma maneira usando um mutex:
```
package main
import (
    "fmt"
    "sync"
    "time"
)
func main() {
    var x int
    var mu sync.Mutex
    go func() {
        mu.Lock()
        x++
        mu.Unlock()
    }()
    time.Sleep(1 * time.Second)
    mu.Lock()
    fmt.Println(x)
    mu.Unlock()
}
```

#### 3. Qual será a saída do seguinte programa? Explique por quê.
```
package main
import "fmt"
func main() {
    s := make([]int, 5)
    s = append(s, 1, 2, 3)
    fmt.Println(s)
}
```

> **Resposta:**
> [0 0 0 0 0 1 2 3]

> **Explicação:**
> O slice s é inicialmente criado com um comprimento de 5, contendo cinco zeros. Quando append é chamado, ele adiciona três elementos ao slice, resultando em um comprimento total de 8.

#### 4. Como Go lida com interfaces nil? Qual é uma armadilha comum associada a elas?

> **Resposta:**
> Uma interface nil em Go é um valor de interface que não contém nem um valor nem um tipo. Uma armadilha comum é verificar apenas o valor para nil, não considerando que uma interface pode ser não-nil enquanto contém um valor concreto nil. Isso pode levar a comportamentos inesperados ao comparar interfaces com nil.

#### 5. Descreva o modelo de memória de Go e suas implicações para programação concorrente.

> **Resposta:**
> O modelo de memória de Go define o comportamento do acesso à memória em programas concorrentes. Ele garante que operações de leitura e escrita em variáveis sejam atômicas e fornece regras para a visibilidade de escritas para outras goroutines. Este modelo exige sincronização adequada para garantir acesso concorrente seguro a variáveis compartilhadas, tipicamente usando canais ou primitivas sync.

#### 6. Como você implementaria um pool de trabalhadores em Go?

> **Resposta:**
> Um pool de trabalhadores pode ser implementado usando goroutines e canais. Aqui está um exemplo:
```
package main
import (
    "fmt"
    "sync"
)
func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
    defer wg.Done()
    for j := range jobs {
        fmt.Printf("Worker %d started job %d\n", id, j)
        results <- j * 2
    }
}
func main() {
    const numJobs = 5
    jobs := make(chan int, numJobs)
    results := make(chan int, numJobs)

    var wg sync.WaitGroup
    for w := 1; w <= 3; w++ {
        wg.Add(1)
        go worker(w, jobs, results, &wg)
    }

    for j := 1; j <= numJobs; j++ {
        jobs <- j
    }
    close(jobs)

    wg.Wait()
    close(results)

    for r := range results {
        fmt.Println("Result:", r)
    }
}
```

#### 7. Quais são alguns erros comuns de desempenho em Go e como evitá-los?

> **Resposta:**
> *Erros comuns de desempenho incluem:*
> - Uso excessivo de goroutines levando a alto consumo de memória.
> - Uso ineficiente de canais causando bloqueios desnecessários.
> - Grandes alocações de memória e pausas frequentes do coletor de lixo.
> - Uso subótimo das funções da biblioteca padrão.

> *Estes podem ser evitados:*
> - Limitando o número de goroutines.
> - Usando canais bufferizados adequadamente.
> - Perfilando a aplicação para identificar e otimizar pontos críticos.
> - Aproveitando estruturas de dados e algoritmos eficientes.

#### 8. Explique a diferença entre um receptor de ponteiro e um

 receptor de valor em métodos Go. Quando você usaria cada um?

> **Resposta:**
> Um receptor de ponteiro permite que métodos modifiquem o valor do receptor e evitem copiá-lo. É usado quando o método precisa mutar o receptor ou quando o receptor é uma struct grande para evitar sobrecarga de desempenho de copiar. Um receptor de valor é usado quando o método não precisa modificar o receptor, e copiar é barato.