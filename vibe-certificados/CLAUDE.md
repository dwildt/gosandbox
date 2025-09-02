# Vibe Certificados - Instruções para Claude

## Visão Geral do Projeto

Este é um projeto Go que implementa uma API REST para geração de certificados em formato HTML e PDF. O sistema permite criar certificados únicos ou em lote através de arquivos CSV.

## Arquitetura do Projeto

```
vibe-certificados/
├── main.go                    # Entrada principal da aplicação
├── go.mod                     # Dependências do Go
├── api/
│   ├── handlers.go           # Handlers HTTP (controladores)
│   └── routes.go             # Configuração de rotas
├── models/
│   ├── certificate.go        # Modelo de certificado
│   └── template.go           # Modelo de template + requests
├── services/
│   ├── certificate_service.go # Lógica de negócio de certificados
│   ├── template_service.go    # Lógica de templates
│   └── pdf_service.go        # Geração de PDF (usando gofpdf)
├── storage/
│   └── memory_storage.go     # Armazenamento em memória
└── tests/
    └── services/             # Testes unitários
```

## Tecnologias e Dependências

### Dependências Principais
- **gin-gonic/gin** v1.10.1 - Framework web para API REST
- **google/uuid** v1.6.0 - Geração de identificadores únicos
- **jung-kurt/gofpdf** v1.16.2 - Geração nativa de PDF

### Versão do Go
- Go 1.24+ (definido no go.mod)

## Como Executar o Projeto

### Desenvolvimento Local
```bash
# Instalar dependências
go mod download

# Executar aplicação
go run main.go

# Aplicação disponível em http://localhost:8080
```

### Comandos de Teste
```bash
# Executar todos os testes
go test ./...

# Executar testes com cobertura
go test -cover ./...

# Executar testes de um pacote específico
go test ./services/
```

### Build de Produção
```bash
# Build executável
go build -o vibe-certificados main.go

# Executar binário
./vibe-certificados
```

## Endpoints da API

### Certificados
- `POST /api/certificates` - Criar certificado único
- `POST /api/certificates/batch` - Criar certificados em lote (CSV)
- `GET /api/certificates/{id}.html` - Baixar certificado em HTML
- `GET /api/certificates/{id}.pdf` - Baixar certificado em PDF
- `GET /api/certificates/by-email/{email}` - Listar certificados por email

### Templates
- `GET /api/templates` - Listar todos os templates
- `POST /api/templates` - Criar novo template
- `GET /api/templates/{id}` - Obter template específico
- `PUT /api/templates/{id}` - Atualizar template
- `DELETE /api/templates/{id}` - Deletar template

### Health Check
- `GET /api/health` - Verificar status da aplicação

## Estrutura de Dados

### Certificado (Certificate)
```go
type Certificate struct {
    ID             string            `json:"id"`
    Email          string            `json:"email"`
    Name           string            `json:"name"`
    Course         string            `json:"course"`
    CompletionDate time.Time         `json:"completion_date"`
    TemplateID     string            `json:"template_id"`
    CreatedAt      time.Time         `json:"created_at"`
    Data           map[string]string `json:"data,omitempty"`
}
```

### Requisição de Certificado (CertificateRequest)
```go
type CertificateRequest struct {
    Email          string            `json:"email" binding:"required"`
    Name           string            `json:"name" binding:"required"`
    Course         string            `json:"course" binding:"required"`
    CompletionDate string            `json:"completion_date" binding:"required"`
    TemplateID     string            `json:"template_id"`
    Data           map[string]string `json:"data,omitempty"`
}
```

## Funcionalidades Implementadas

### Geração de PDF
- Implementada usando a biblioteca `gofpdf` (Go nativo)
- Layout profissional em formato paisagem A4
- Não requer dependências externas do sistema
- Fontes: Arial com tamanhos variados para hierarquia visual

### Geração de HTML  
- Templates configuráveis via JSON
- Usa Go templates para renderização
- Suporte a campos dinâmicos

### Armazenamento
- **Atual**: Em memória (adequado para desenvolvimento/testes)
- **Futuro**: Pode ser expandido para banco de dados

## Testes

### Estrutura de Testes
- Testes unitários em `tests/services/`
- Cobertura de cenários positivos e negativos
- Validação de inputs e outputs

### Executando Testes
```bash
go test ./tests/services/ -v
```

## Desenvolvimento e Manutenção

### Adicionando Novas Funcionalidades
1. **Modelos**: Adicionar em `models/`
2. **Lógica de negócio**: Implementar em `services/`
3. **Endpoints**: Adicionar handlers em `api/handlers.go` e rotas em `api/routes.go`
4. **Testes**: Criar testes correspondentes em `tests/`

### Padrões de Código
- Seguir conventions Go padrão
- Usar validação com tags binding do Gin
- Tratamento de erros consistente
- Logging adequado para debugging

### Debugging
- Usar logs do Gin (modo debug por padrão)
- Para produção: `export GIN_MODE=release`
- Health check sempre disponível em `/api/health`

## Comandos Úteis para Claude

```bash
# Verificar status da aplicação
curl http://localhost:8080/api/health

# Criar certificado de teste
curl -X POST http://localhost:8080/api/certificates \
  -H "Content-Type: application/json" \
  -d '{"email":"test@example.com","name":"Test User","course":"Test Course","completion_date":"2024-01-15","template_id":"default"}'

# Testar PDF (substitua {id} pelo ID retornado)
curl -o test.pdf http://localhost:8080/api/certificates/{id}.pdf

# Listar templates
curl http://localhost:8080/api/templates

# Executar testes
go test ./... -v

# Verificar dependências
go mod tidy && go mod download
```

## Limitações Atuais

- Armazenamento apenas em memória (dados perdidos ao reiniciar)
- Templates básicos (podem ser expandidos)
- Sem autenticação/autorização
- Sem rate limiting

## Próximos Passos Sugeridos

1. Implementar persistência em banco de dados
2. Adicionar sistema de autenticação
3. Melhorar templates com mais opções de design
4. Adicionar validações mais robustas
5. Implementar cache para melhor performance
6. Adicionar métricas e monitoramento