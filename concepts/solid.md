# 🏗️ Teoria do SOLID

O **SOLID** é um acrônimo criado por Robert C. Martin (Uncle Bob) que representa cinco princípios de design de software.  
O objetivo é criar sistemas **mais coesos, desacoplados, fáceis de manter e evoluir**.

---

## 1. S — Single Responsibility Principle (SRP)

**Princípio da Responsabilidade Única**

> Uma classe deve ter apenas um motivo para mudar.

- Cada módulo/componente deve assumir **uma única responsabilidade**.
- Evita que uma única classe se torne “faz-tudo”.
- **Benefícios:** manutenção mais simples, testes mais fáceis, menor acoplamento.

---

## 2. O — Open/Closed Principle (OCP)

**Princípio Aberto/Fechado**

> Entidades de software devem estar **abertas para extensão**, mas **fechadas para modificação**.

- O comportamento pode ser estendido sem alterar o código existente.
- Ex.: adicionar novas regras através de abstrações, sem reescrever o núcleo.
- **Benefícios:** evita regressões e reduz risco ao evoluir o sistema.

---

## 3. L — Liskov Substitution Principle (LSP)

**Princípio da Substituição de Liskov**

> Objetos de uma superclasse devem poder ser substituídos por objetos de suas subclasses **sem quebrar o sistema**.

- A hierarquia deve ser coerente: se B herda de A, então B deve se comportar como A.
- **Benefícios:** garante polimorfismo confiável, previsível e seguro.

---

## 4. I — Interface Segregation Principle (ISP)

**Princípio da Segregação de Interfaces**

> É melhor criar **interfaces específicas e pequenas** do que interfaces grandes e genéricas.

- Clientes não devem ser forçados a depender de métodos que não usam.
- **Benefícios:** reduz acoplamento e torna implementações mais enxutas.

---

## 5. D — Dependency Inversion Principle (DIP)

**Princípio da Inversão de Dependência**

> Módulos de alto nível não devem depender de módulos de baixo nível.  
> Ambos devem depender de **abstrações**.

- Detalhes devem depender de abstrações, não o contrário.
- Ex.: ao invés de depender de uma classe concreta de banco de dados, depender de uma interface de repositório.
- **Benefícios:** maior flexibilidade, testabilidade e independência de implementação.

---

# ✅ Conclusão

O **SOLID** fornece uma base para escrever código:

- **Mais modular**
- **Mais fácil de manter**
- **Mais testável**
- **Menos propenso a acoplamentos perigosos**

É um guia conceitual para projetar software de qualidade, independentemente da linguagem usada.
