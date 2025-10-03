---

## **1. Estrutura de Projeto**

**P1:** Como você estrutura um projeto Go de médio porte?
**R:** Uso `go mod` para gerenciar dependências, separo `cmd/` para binários, `internal/` para código privado, `pkg/` para pacotes reutilizáveis, e `api/` para contratos ou handlers. Isso ajuda a manter clareza e escalabilidade.

---

## **2. Tipos e Interfaces**

**P2:** Qual a diferença entre `interface{}` e `any` no Go?
**R:** `any` é apenas um _alias_ para `interface{}`, introduzido no Go 1.18 para tornar o código genérico mais legível. Funcionalmente são idênticos.

**P3:** Como funciona _duck typing_ em Go?
**R:** Uma struct implementa uma interface implicitamente se possui todos os métodos exigidos por ela. Não é necessário declarar explicitamente que uma struct implementa a interface.

---

## **3. Concurrency**

**P4:** Explique a diferença entre _concurrency_ e _parallelism_ no Go.
**R:** _Concurrency_ é lidar com várias tarefas ao mesmo tempo de forma intercalada, enquanto _parallelism_ é executar várias tarefas simultaneamente em múltiplos núcleos. Goroutines são concorrentes, e podem rodar em paralelo dependendo de `GOMAXPROCS`.

**P5:** Para que serve o `sync.WaitGroup`?
**R:** Ele sincroniza a execução de múltiplas goroutines, permitindo que o programa espere todas terminarem antes de prosseguir.

---

## **4. Channels**

**P6:** Qual a diferença entre canal com _buffer_ e sem _buffer_?
**R:** Um canal sem buffer bloqueia o envio até que haja um receptor pronto. Um canal com buffer permite enfileirar valores até o limite do buffer sem bloquear o envio.

**P7:** O que acontece se você fechar um canal?
**R:** Os receptores ainda podem ler os valores que estavam no buffer. Leituras subsequentes retornam o _zero value_ do tipo e um `ok=false`, permitindo detectar que o canal foi fechado.

---

## **5. Context e Cancelamento**

**P8:** Para que serve o pacote `context`?
**R:** Ele propaga prazos (_deadlines_), cancelamentos e valores entre goroutines. Essencial para evitar goroutines "zumbis" e liberar recursos corretamente.

**P9:** Qual a diferença entre `context.Background()` e `context.TODO()`?
**R:** `Background()` é usado como contexto raiz em produção. `TODO()` é usado temporariamente quando ainda não se sabe qual contexto usar, útil durante desenvolvimento.

---

## **6. Erros**

**P10:** Como você lida com erros em Go?
**R:** Retorno erros como segundo valor (`result, err := ...`). Quando preciso de contexto, uso `fmt.Errorf("msg: %w", err)` para _wrapping_. Para erros bem definidos, implemento tipos de erro customizados.

**P11:** Quando faz sentido usar `panic`?
**R:** Apenas em situações irreversíveis ou de falha de programação, como índices inválidos ou inconsistências internas. Não deve ser usado para fluxo normal de erro.

---

## **7. Generics**

**P12:** O que mudou com a introdução de generics no Go 1.18?
**R:** Agora podemos escrever funções e tipos parametrizados com `type parameters`, reduzindo duplicação de código. Por exemplo, uma função genérica `Min[T constraints.Ordered](a, b T) T`.

**P13:** O que são _constraints_ em generics?
**R:** São restrições que definem quais tipos podem ser usados como parâmetro genérico, como `constraints.Ordered` ou interfaces customizadas.

---

## **8. Performance e Profiling**

**P14:** Como você faz profiling em Go?
**R:** Uso o pacote `net/http/pprof` ou `runtime/pprof` para capturar perfis de CPU, heap e blocos. Depois analiso com `go tool pprof` para encontrar _hotspots_.

**P15:** O que é _escape analysis_?
**R:** É a análise que o compilador faz para decidir se uma variável pode ficar na stack ou precisa ir para o heap. Se "escapa" para fora da função, vai para o heap.

---

## **9. Ferramentas e Testes**

**P16:** Como você escreve testes concorrentes de forma determinística?
**R:** Uso `t.Parallel()` em testes que podem rodar em paralelo, e bibliotecas como `testing/synctest` para controlar execução de goroutines de forma reprodutível.

**P17:** Qual a diferença entre `go test` e `go test -race`?
**R:** `-race` habilita o _race detector_, que detecta condições de corrida em variáveis acessadas concorrentemente sem sincronização adequada.

---

## **10. Boas Práticas e Deploy**

**P18:** Como você garante que seu binário Go seja pequeno para produção?
**R:** Compilo com `go build -ldflags="-s -w"` para remover símbolos de debug, e uso ferramentas como `upx` se necessário para compressão.

**P19:** Como você lida com variáveis de ambiente no Go?
**R:** Uso `os.Getenv` ou bibliotecas como `envconfig`/`viper` para carregar configs. Também valido os valores no startup para falhar rápido em caso de erro.

**P20:** Por que Go é uma boa escolha para microsserviços?
**R:** É rápido, possui binários estáticos, excelente suporte a concorrência, baixa utilização de memória e uma curva de aprendizado simples, facilitando manutenção e deploy.

---

Quer que eu monte isso em **formato de simulado de entrevista** (onde você responde e depois confere a resposta certa) ou em um **PDF de estudo** igual a um cheatsheet?

Primitivos: bool, int, float, string, rune, complex.

Objetos (compostos): array, slice, map, chan, struct, interface, func, pointer.
