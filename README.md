# Golang BFA (Backend For Frontend)

Um projeto Backend For Frontend (BFF) em Golang utilizando o framework Fiber v3. Este projeto demonstra uma arquitetura escalável e testável para APIs modernas em Go, seguindo boas práticas da comunidade.

## Pré-requisitos

- Go 1.21 ou superior
- Git

## Instalação

1. Clone o repositório:
   ```bash
   git clone <url-do-repositorio>
   cd golang-bfa
   ```

2. Instale as dependências:
   ```bash
   go mod tidy
   ```

## Execução

Para iniciar o servidor:

```bash
go run cmd/api/main.go
```

O servidor será iniciado na porta 8080 (ou na porta definida pela variável de ambiente `PORT`).

## Testes

Para executar os testes unitários:

```bash
go test ./tests/...
```

## Estrutura do Projeto

```
golang-bfa/
├── cmd/
│   └── api/
│       └── main.go                 # Ponto de entrada da aplicação
├── internal/
│   ├── config/
│   │   └── config.go               # Configurações da aplicação
│   ├── server/
│   │   └── fiber.go                # Inicialização do servidor Fiber
│   ├── routes/
│   │   └── routes.go               # Definição das rotas
│   ├── handlers/
│   │   ├── health_handler.go       # Handler para endpoint de saúde
│   │   └── user_handler.go         # Handler para usuários
│   ├── services/
│   │   └── user_service.go         # Lógica de negócio para usuários
│   ├── repositories/
│   │   └── user_repository.go      # Acesso a dados (simulado)
│   ├── models/
│   │   └── user.go                 # Modelos de domínio
│   ├── dto/
│   │   ├── user_request.go         # DTOs de request
│   │   └── user_response.go        # DTOs de response
│   └── middlewares/
│       ├── logger.go               # Middleware de logging
│       └── error_handler.go        # Middleware de tratamento de erros
├── pkg/
│   └── logger/
│       └── logger.go               # Logger estruturado reutilizável
└── tests/
    ├── handlers/
    │   └── user_handler_test.go    # Testes para handlers
    └── services/
        └── user_service_test.go    # Testes para services
```

## Endpoints da API

### Health Check
- **GET** `/health`
  - Retorna o status da aplicação.

### Usuários
- **GET** `/users/:id`
  - Retorna um usuário pelo ID.
- **POST** `/users`
  - Cria um novo usuário.

### Exemplos de Uso

#### Criar um usuário
```bash
curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{"name": "João Silva", "email": "joao@example.com"}'
```

#### Obter um usuário
```bash
curl http://localhost:8080/users/1
```

#### Health Check
```bash
curl http://localhost:8080/health
```

## Arquitetura

O projeto segue uma arquitetura em camadas:

- **Handlers**: Recebem requisições HTTP, fazem parsing e validação, chamam services.
- **Services**: Contêm regras de negócio, independentes de HTTP.
- **Repositories**: Responsáveis por acesso a dados ou APIs externas.
- **Models**: Estruturas de domínio.
- **DTOs**: Objetos de request/response da API.

Utiliza injeção de dependência manual para facilitar testes e manutenção.

## Tecnologias Utilizadas

- [Fiber v3](https://github.com/gofiber/fiber) - Framework web para Go
- [Testify](https://github.com/stretchr/testify) - Framework de testes
- [slog](https://pkg.go.dev/log/slog) - Logger estruturado do Go

## Desenvolvimento

Para contribuir:

1. Faça um fork do projeto
2. Crie uma branch para sua feature (`git checkout -b feature/nova-feature`)
3. Commit suas mudanças (`git commit -am 'Adiciona nova feature'`)
4. Push para a branch (`git push origin feature/nova-feature`)
5. Abra um Pull Request

## Licença

Este projeto está sob a licença MIT. Veja o arquivo LICENSE para mais detalhes.