# 🏦 GoBank

> API bancária desenvolvida em **Go (Golang)** com foco em estudo de desenvolvimento backend, arquitetura, autenticação, sessões, PostgreSQL e boas práticas para construção de APIs REST.

O **GoBank** é um projeto de estudo criado para aprofundar conhecimentos em desenvolvimento de aplicações backend utilizando Go, explorando conceitos comuns em sistemas financeiros e aplicações que trabalham com dados sensíveis.

> ⚠️ **Projeto para fins educacionais.** Não deve ser utilizado para operações financeiras reais.

---

## 🚀 Objetivo

O principal objetivo do GoBank é servir como um laboratório prático para estudar e aplicar conceitos importantes do ecossistema Go.

Durante o desenvolvimento, o projeto busca trabalhar conceitos como:

* Desenvolvimento de APIs REST
* Organização de projetos em Go
* Injeção de dependências
* PostgreSQL
* `pgx` e `pgxpool`
* Gerenciamento de sessões
* Cookies HTTP
* Autenticação
* Hash de senhas
* UUID
* Context
* Middlewares
* Docker
* Variáveis de ambiente
* Separação de responsabilidades
* Tratamento de erros
* Modelagem de dados
* Regras de negócio

---

## 🛠️ Tecnologias utilizadas

| Tecnologia            | Utilização                             |
| --------------------- | -------------------------------------- |
| 🐹 **Go**             | Linguagem principal                    |
| 🌐 **Chi**            | Router HTTP                            |
| 🐘 **PostgreSQL**     | Banco de dados                         |
| 🔌 **pgx**            | Driver/cliente PostgreSQL              |
| 🔐 **SCS**            | Gerenciamento de sessões               |
| 🍪 **pgxstore**       | Persistência das sessões no PostgreSQL |
| 🆔 **UUID**           | Identificação de entidades             |
| 🔒 **x/crypto**       | Recursos criptográficos                |
| 🐳 **Docker Compose** | Ambiente do PostgreSQL                 |
| ⚙️ **godotenv**       | Carregamento das variáveis de ambiente |

As dependências principais podem ser conferidas no `go.mod` do projeto.

---

## 📐 Estrutura do projeto

O projeto utiliza uma organização baseada na separação de responsabilidades:

```text
GoBank/
│
├── cmd/
│   └── api/
│       └── main.go
│
├── internal/
│   ├── api/
│   │
│   ├── services/
│   │   ├── user/
│   │   ├── pessoa_fisica/
│   │   └── pessoa_juridica/
│   │
│   └── ...
│
├── .env.example
├── .gitignore
├── docker-compose.yml
├── go.mod
├── go.sum
└── README.md
```

A aplicação é inicializada através de `cmd/api/main.go`, onde são configurados o PostgreSQL, pool de conexões, sessões, serviços e rotas HTTP.

---

## 🏗️ Arquitetura

O projeto busca aplicar uma separação entre:

```text
HTTP Request
      │
      ▼
   Router
      │
      ▼
   Handler
      │
      ▼
   Service
      │
      ▼
 PostgreSQL
```

### Router

Responsável por receber as requisições HTTP e direcioná-las para os endpoints correspondentes.

O projeto utiliza o **Chi Router**.

### Handler / API

Responsável por:

* Receber a requisição;
* Validar os dados de entrada;
* Chamar os serviços;
* Montar a resposta HTTP;
* Definir status codes.

### Service

Concentra as regras de negócio da aplicação.

Exemplos:

```text
UserService
PessoaFisicaService
PessoaJuridicaService
```

### Repository / Database

Responsável pela comunicação com o PostgreSQL através do `pgx`.

---

# 🔐 Autenticação e sessões

O GoBank utiliza sessões HTTP através do pacote **SCS**.

A sessão é persistida no PostgreSQL utilizando:

```text
scs
   │
   └── pgxstore
           │
           ▼
       PostgreSQL
```

A aplicação também configura propriedades de segurança para o cookie da sessão, incluindo:

* `HttpOnly`
* `SameSite=Lax`
* Tempo de expiração da sessão

Atualmente, a sessão possui lifetime configurado de **24 horas**.

---

# 🐘 Banco de dados

O projeto utiliza **PostgreSQL** como banco de dados principal.

A aplicação utiliza `pgxpool` para gerenciamento do pool de conexões.

Exemplo conceitual da configuração:

```text
Go API
  │
  ▼
pgxpool
  │
  ▼
PostgreSQL
```

O banco pode ser executado utilizando Docker Compose.

---

# 🐳 Executando o PostgreSQL

O projeto possui um `docker-compose.yml` configurado para executar o PostgreSQL.

O serviço utiliza:

```yaml
image: postgres:latest
```

e cria um volume persistente para os dados do banco.

Execute:

```bash
docker compose up -d
```

Para verificar o container:

```bash
docker ps
```

Para parar o ambiente:

```bash
docker compose down
```

---

# ⚙️ Configuração

Crie um arquivo `.env` baseado no exemplo:

```bash
cp .env.example .env
```

Configure:

```env
GOBANK_DATABASE_PORT=5432
GOBANK_DATABASE_HOST="localhost"
GOBANK_DATABASE_USER="pgUser"
GOBANK_DATABASE_PASSWORD="pgPassword"
GOBANK_DATABASE_NAME="GOBANK"
```

Essas variáveis são utilizadas pela aplicação para montar a conexão com o PostgreSQL.

> 🔒 Nunca versione o arquivo `.env` contendo credenciais reais.

---

# ▶️ Executando a aplicação

Depois de iniciar o PostgreSQL:

```bash
go mod download
```

Execute a API:

```bash
go run ./cmd/api
```

A aplicação será iniciada na porta:

```text
8080
```

Acesse:

```text
http://localhost:8080
```

---

# 📚 Conceitos estudados

Este projeto está sendo utilizado como laboratório para estudar principalmente:

### Go

* Structs
* Interfaces
* Pointers
* Packages
* Goroutines
* Context
* Error handling
* Dependency Injection
* `net/http`
* JSON
* Middleware

### Backend

* API REST
* HTTP methods
* Status codes
* Cookies
* Sessões
* Autenticação
* Autorização
* Validação
* Tratamento de erros

### PostgreSQL

* Modelagem relacional
* Queries
* Constraints
* Foreign Keys
* Transactions
* Índices
* Connection Pool

### Segurança

* Hash de senha
* Sessões
* Cookies `HttpOnly`
* `SameSite`
* Variáveis de ambiente
* Proteção de credenciais

### DevOps

* Docker
* Docker Compose
* Configuração por ambiente

---

# 🧪 Próximos estudos

O GoBank será evoluído gradualmente para explorar conceitos mais avançados.

### 🔹 Autenticação

* [ ] Login
* [ ] Logout
* [ ] Middleware de autenticação
* [ ] Controle de sessão
* [ ] Expiração de sessão
* [ ] Recuperação de senha

### 🔹 Usuários

* [ ] Cadastro
* [ ] Atualização
* [ ] Exclusão
* [ ] Validação de dados
* [ ] Perfil do usuário

### 🔹 Contas bancárias

* [ ] Criar conta
* [ ] Consultar conta
* [ ] Consultar saldo
* [ ] Bloquear conta
* [ ] Encerrar conta

### 🔹 Transações

* [ ] Depósito
* [ ] Saque
* [ ] Transferência
* [ ] Histórico de transações
* [ ] Validação de saldo
* [ ] Transações PostgreSQL
* [ ] Controle de concorrência

### 🔹 API

* [ ] Padronização de responses
* [ ] Padronização de erros
* [ ] Middleware de logs
* [ ] Middleware de recovery
* [ ] Request ID
* [ ] Paginação
* [ ] Validação

### 🔹 Testes

* [ ] Unit Tests
* [ ] Integration Tests
* [ ] Testes dos services
* [ ] Testes dos handlers
* [ ] Testes de banco
* [ ] Testes de concorrência

### 🔹 Infraestrutura

* [ ] Dockerizar a API
* [ ] Docker Compose completo
* [ ] Health Check
* [ ] CI/CD
* [ ] GitHub Actions

---

# 📊 Roadmap

```text
[✓] Estrutura inicial do projeto
[✓] Configuração PostgreSQL
[✓] Docker Compose
[✓] Pool de conexões
[✓] Router HTTP
[✓] Services
[✓] Sessões
[ ] Autenticação completa
[ ] Middleware de autenticação
[ ] Contas bancárias
[ ] Transações
[ ] Transferências
[ ] Testes automatizados
[ ] Observabilidade
[ ] CI/CD
```

---

# 🎯 Objetivo de aprendizado

O GoBank não tem como objetivo apenas implementar um CRUD.

A proposta é utilizar um domínio conhecido — **sistema bancário** — para estudar problemas reais de backend, como:

```text
                    ┌──────────────┐
                    │    Cliente   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   HTTP API   │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │    Router    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   Handler    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │   Service    │
                    └──────┬───────┘
                           │
                           ▼
                    ┌──────────────┐
                    │  PostgreSQL  │
                    └──────────────┘
```

A ideia é evoluir o projeto junto com o aprendizado de Go, adicionando gradualmente problemas mais próximos dos encontrados em aplicações backend profissionais.

---

# 📖 Status

🚧 **Em desenvolvimento**

Este projeto faz parte dos meus estudos de **Go, Backend e Engenharia de Software**.

Novas funcionalidades e melhorias serão adicionadas conforme novos conceitos forem estudados.

---

## 👨‍💻 Autor

**Mateus R. de Lima**

GitHub:

👉 [Mateus-R-De-Lima](https://github.com/Mateus-R-De-Lima)

Repositório:

👉 [GoBank](https://github.com/Mateus-R-De-Lima/GoBank)

---

## 📄 Licença

Este projeto é destinado a fins educacionais e de estudo.
