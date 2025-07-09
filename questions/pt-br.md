## 🤖 Perguntas genéricas 🤖

#### 🔎 Teste o código no [playground](https://go.dev/play/). 🔎

#### 1. O que é Go? Explique suas principais características.

> **Resposta:**
> Go é uma linguagem de código aberto, fortemente tipada e compilada, escrita para construir software concorrente e escalável. A linguagem foi inventada no Google por Rob Pike, Ken Thomson e Robert Griesemer. É estaticamente tipada, compilada e conhecida por sua simplicidade, eficiência e suporte à concorrência. Os principais recursos incluem coleta de lixo, uma biblioteca padrão rica, gorrotinas para concorrência e ferramentas integradas para formatação de código e testes.

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

> Aqui está um exemplo básico:

```go
package main
import "fmt"
func main() {
	f := foo()
	fmt.Println(f)
}
func foo() (r string) {
	defer func() {
		r = r + "after"
	}()
	return "before "
}
```

#### 5. Explique como Go lida com erros.

> **Resposta:**
> Go lida com erros usando um tipo de erro simples e explícito. Funções que podem resultar em um erro retornam um valor de erro adicional. O tratamento de erros é feito verificando esse valor retornado. A biblioteca padrão fornece o pacote errors para criar e manipular erros. Importante ressaltar que [erros são tratados como valores.](https://go.dev/blog/errors-are-values)

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
>
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

#### 13. O que é um array em Go?

> **Resposta:**
> Um array em Go é uma coleção de elementos do mesmo tipo, com um tamanho fixo definido no momento da declaração.

> **Explicação:**
> Os arrays têm um tamanho fixo, o que significa que o número de elementos é especificado quando o array é declarado e não pode ser alterado. Por exemplo, var arr [5]int declara um array de inteiros com cinco elementos. O tipo do array inclui seu tamanho, então [5]int e [10]int são tipos diferentes.

#### 14. O que é um slice em Go?

> **Resposta:**
> Um slice em Go é uma sequência de elementos que é mais flexível que um array, pois seu tamanho pode crescer ou diminuir dinamicamente.

> **Explicação:**
> Um slice é uma referência a um segmento de um array, permitindo que seu comprimento e capacidade sejam ajustados conforme necessário. Slices são usados com mais frequência do que arrays em Go devido à sua flexibilidade. Por exemplo, var slice []int declara um slice de inteiros. Slices podem ser redimensionados usando funções como append.

#### 15. Qual é a principal diferença entre um array e um slice em Go?

> **Resposta:**
> A principal diferença é que arrays têm um tamanho fixo, enquanto slices têm um tamanho dinâmico que pode crescer ou diminuir conforme necessário.

#### 16. O que é `go vet` em Go?

**Resposta:**

> `go vet` é uma ferramenta que analisa o código-fonte em Go para detectar construções suspeitas ou erros comuns que não são detectados pelo compilador.
> Ao executar `go vet`, você pode identificar e corrigir esses problemas antes de compilar o código, ajudando a evitar bugs e comportamentos inesperados.

#### 17. O que é `go fmt` em Go?

**Resposta:**

> `go fmt` reformata o código para seguir as convenções de formatação do Go, tornando o código mais legível e consistente. Isso inclui:
>
> - Indentação correta.
> - Espaçamento adequado.
> - Quebra de linhas longas.
>   Usar `go fmt` ajuda a manter um estilo de código uniforme em projetos, facilitando a leitura e manutenção do código por diferentes desenvolvedores.

#### 18. O que é o pacote `Context` em Go?

**Resposta:**

> O pacote context é usado para gerenciar prazos, sinais de cancelamento e outros valores escopados a pedidos em limites de API e entre processos. Ele é normalmente usado para controlar e gerenciar o ciclo de vida de uma solicitação, permitindo desligamentos suaves, tempos limite e propagação de sinais de cancelamento. O contexto é geralmente passado como o primeiro parâmetro para funções que executam operações de longa duração, garantindo que possam ser terminadas, se necessário.

#### 19. O que é uma `Struct` em Go?

> Em Go, uma `struct` é um tipo de dado composto que agrupa variáveis sob um único nome. Essas variáveis, conhecidas como campos, podem ter diferentes tipos e são acessadas usando notação de ponto. Structs são usadas para criar estruturas de dados complexas e são centrais para a programação orientada a objetos em Go.
> Aqui está um exemplo básico:

```go
package main
import "fmt"
// Define uma struct chamada Person
type Person struct {
    Name  string
    Age   int
    Email string
}
func main() {
    // Cria uma instância da struct Person
    p := Person{
        Name:  "Alice",
        Age:   30,
        Email: "alice@example.com",
    }
    // Acessa os campos da struct
    fmt.Println("Name:", p.Name)
    fmt.Println("Age:", p.Age)
    fmt.Println("Email:", p.Email)
}
```

> **Explicação:**
>
> - `Person` é um tipo de struct com três campos: `Name` (string), `Age` (int) e `Email` (string).
> - Uma instância de `Person` é criada e inicializada com valores para cada campo.
> - Os campos são acessados usando a notação de ponto, por exemplo, `p.Name`.

## ❗❗ Perguntas difíceis ❗❗

#### 1. Qual é a saída do seguinte programa Go e por quê?

```
import "fmt"
package main
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")
}
```

> **Resposta:** \
> hello \
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

> **Resposta:** [0 0 0 0 0 1 2 3]

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

> **Resposta:** > _Erros comuns de desempenho incluem:_
>
> - Uso excessivo de goroutines levando a alto consumo de memória.
> - Uso ineficiente de canais causando bloqueios desnecessários.
> - Grandes alocações de memória e pausas frequentes do coletor de lixo.
> - Uso subótimo das funções da biblioteca padrão.

> _Estes podem ser evitados:_
>
> - Limitando o número de goroutines.
> - Usando canais bufferizados adequadamente.
> - Perfilando a aplicação para identificar e otimizar pontos críticos.
> - Aproveitando estruturas de dados e algoritmos eficientes.

#### 8. Explique a diferença entre um receptor de ponteiro e um

> **Resposta:**
> Um receptor de ponteiro permite que métodos modifiquem o valor do receptor e evitem copiá-lo. É usado quando o método precisa mutar o receptor ou quando o receptor é uma struct grande para evitar sobrecarga de desempenho de copiar. Um receptor de valor é usado quando o método não precisa modificar o receptor, e copiar é barato.

#### 9. O que é um canal sem buffer em Go?

> **Resposta:**
> Um canal sem buffer é um tipo de canal que não tem capacidade para armazenar valores. Quando um valor é enviado para um canal sem buffer, a goroutine que envia é bloqueada até que outra goroutine receba o valor do canal. Da mesma forma, uma goroutine que recebe é bloqueada até que um valor seja enviado para o canal.

> **Explicação:**
> Canais sem buffer impõem uma sincronização estrita entre as goroutines de envio e recebimento. Isso significa que as operações de envio e recebimento acontecem simultaneamente, garantindo que ambas as goroutines coordenem de perto.

#### 10. O que é um canal com buffer em Go?

> **Resposta:**
> Um canal com buffer é um tipo de canal que tem capacidade para armazenar um número específico de valores. Quando um valor é enviado para um canal com buffer, a goroutine que envia só é bloqueada se o canal estiver cheio. Da mesma forma, a goroutine que recebe só é bloqueada se o canal estiver vazio.

> **Explicação:**
> Canais com buffer permitem comunicação assíncrona entre goroutines. Isso significa que um remetente pode enviar valores para o canal sem esperar por um receptor, até a capacidade do canal. Da mesma forma, um receptor pode receber valores do canal sem esperar por um remetente, desde que haja valores no buffer.

#### 11. O que é e como são usados os Generics em Go?

**Resposta:**

> Generics são uma funcionalidade em linguagens de programação que permitem escrever código flexível e reutilizável, possibilitando que funções, estruturas de dados e tipos operem com diferentes tipos de dados sem sacrificar a segurança de tipos. No Go eles permitem a criação de funções e tipos que podem operar com qualquer tipo especificado.

**Exemplo**

```go
package main
import "fmt"
func Max[T any](values []T) T {
    var max T
    for _, v := range values {
        if v > max {
            max = v
        }
    }
    return max
}

func main() {
    ints := []int{1, 2, 3, 4, 5}
    floats := []float64{1.1, 2.2, 3.3, 4.4, 5.5}

    fmt.Println("Max int:", Max(ints))
    fmt.Println("Max float:", Max(floats))
}
```

> **Explicação:**

> `Max` é uma função genérica que pode operar em slices de qualquer tipo `T`.
> O parâmetro de tipo `T` é especificado entre colchetes `[]`.
> A restrição `any` indica que `T` pode ser de qualquer tipo.

#### 12. Quais são os principais benefícios dos generics?

> **Resposta:**

> Os principais benefícios dos generics incluem:
>
> - Reusabilidade: Permitem escrever código mais reutilizável sem duplicar para cada tipo.
> - Segurança de Tipos: Garantem a correção dos tipos em tempo de compilação.
> - Manutenibilidade: Simplificam a manutenção ao reduzir a duplicação de código.

#### 13. O que são `Métodos` em Go?

> **Resposta:**

> Em Go, métodos são funções associadas a um tipo específico, permitindo definir comportamentos para seus tipos personalizados (structs). Métodos utilizam um argumento de receptor especial que aparece entre a palavra-chave `func` e o nome do método.
> Aqui está um exemplo:

```go
package main
import "fmt"
// Define uma struct chamada Person
type Person struct {
    Name string
    Age  int
}
// Define um método na struct Person
func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}
func main() {
    // Cria uma instância da struct Person
    p := Person{
        Name: "Alice",
        Age:  30,
    }
    // Chama o método Greet na instância Person
    p.Greet()
}
```

> **Explicação:**

> - `Person` é um tipo de struct com campos `Name` e `Age`.
> - `Greet` é um método com um receptor do tipo `Person`. O receptor `(p Person)` permite que `Greet` acesse os campos da struct `Person`.
> - O método é chamado usando a notação de ponto, por exemplo, `p.Greet()`.
>   Métodos em Go podem ter receptores de valor ou de ponteiro:
>
> 1. **_Receptores de Valor_**: O método opera em uma cópia do receptor. Mudanças feitas no receptor dentro do método não afetam o valor original. Isso é usado quando o método não precisa modificar o receptor ou quando o receptor é pequeno e barato de copiar.
> 2. **_Receptores de Ponteiro_**: O método opera no receptor original. Mudanças feitas no receptor dentro do método afetam o valor original. Isso é usado quando o método precisa modificar o receptor ou quando o receptor é grande e copiá-lo seria ineficiente.
>    Exemplo com um receptor de ponteiro:

```go
package main
import "fmt"
// Define uma struct chamada Person
type Person struct {
    Name string
    Age  int
}
// Define um método com um receptor de ponteiro
func (p *Person) HaveBirthday() {
    p.Age++
}
func main() {
    p := Person{
        Name: "Alice",
        Age:  30,
    }
    // Chama o método HaveBirthday na instância Person
    p.HaveBirthday()
    fmt.Println("After birthday, Age:", p.Age)  // Outputs: After birthday, Age: 31
}
```

> Neste exemplo, `HaveBirthday` usa um receptor de ponteiro `(p *Person)`, permitindo que ele modifique o campo `Age` da struct `Person`.

#### 14. Explique os usos do pacote `sync/atomic`.

> O pacote `atomic` fornece operações de memória atômica de baixo nível, que são cruciais para a implementação de algoritmos de sincronização. No entanto, essas funções devem ser usadas com extremo cuidado. Na maioria dos casos, é melhor usar canais ou outras ferramentas de sincronização fornecidas pelo pacote `sync`. O princípio a ser seguido é: compartilhe memória comunicando; não comunique compartilhando memória.
> Aqui está um exemplo:

```go
package main
import (
	"fmt"
	"sync/atomic"
	"time"
)
func main() {
	var x atomic.Int32
	x.Store(1)
	go func() {
		x.Add(1)
	}()
	time.Sleep(1 * time.Second)
	fmt.Println(x.Load())
    // 2
}
```

### 15. Defina tipos first-class citizen em Golang.

**Explicação:**

> Em Golang, os tipos first-class citizen são estruturas de dados embutidas que têm suporte completo na linguagem. Isso significa que você pode::
>
> 1.Atribuí-los a variáveis
> 2.Passá-los como argumentos para funções
> 3.Retorná-los de funções
> 4.Armazená-los em outros contêineres
>
> Tipos comuns de contêineres de primeira classe em Go:
>
> - Slices: []int, []string, etc.
> - Maps: map[string]int, map[int]string, etc.
> - Arrays: [5]int, [3]string, etc.
> - Channels: chan int, chan string, etc.
