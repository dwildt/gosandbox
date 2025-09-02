# Vibe Certificados - Certificate Generator API

## Descrição / Description

Sistema para geração de certificados em HTML ou PDF através de uma API REST.

## Funcionalidades / Features

- **Geração de certificados**: A partir de templates JSON configuráveis
- **Formatos de saída**: HTML e PDF
- **Entrada de dados**: 
  - Parâmetros únicos via API
  - Lote via arquivo CSV
- **Agrupamento**: Certificados agrupados por email da pessoa
- **Identificação única**: Cada certificado possui um UUID único
- **API REST**: Construída com gin-gonic
- **Documentação**: Swagger integrado

## Arquitetura / Architecture

```
vibe-certificados/
├── README.md
├── go.mod
├── main.go
├── api/
│   ├── handlers.go
│   ├── routes.go
│   └── swagger.go
├── models/
│   ├── certificate.go
│   └── template.go
├── services/
│   ├── certificate_service.go
│   ├── template_service.go
│   └── pdf_service.go
├── templates/
│   └── default.json
├── storage/
│   └── memory_storage.go
└── tests/
    ├── unit/
    └── integration/
```

## Endpoints da API / API Endpoints

### Certificados / Certificates
- `POST /api/certificates` - Gerar certificado único
- `POST /api/certificates/batch` - Gerar certificados em lote via CSV
- `GET /api/certificates/{id}.html` - Exportar certificado em HTML
- `GET /api/certificates/{id}.pdf` - Exportar certificado em PDF
- `GET /api/certificates/by-email/{email}` - Listar certificados por email

### Templates
- `GET /api/templates` - Listar templates disponíveis
- `POST /api/templates` - Criar novo template
- `GET /api/templates/{id}` - Obter template específico
- `PUT /api/templates/{id}` - Atualizar template
- `DELETE /api/templates/{id}` - Remover template

## Pré-requisitos / Prerequisites

- Go 1.24+ instalado
- Git para controle de versão

## Instalação e Execução / Installation and Execution

```bash
# Clonar o repositório
git clone <repository-url>
cd vibe-certificados

# Instalar dependências
go mod download

# Executar aplicação
go run main.go

# A aplicação estará disponível em:
# http://localhost:8080

# Executar testes
go test ./...

# Executar testes com cobertura
go test -cover ./...

# Build da aplicação
go build -o vibe-certificados main.go
./vibe-certificados
```

## Uso / Usage

### Geração de certificado único:
```bash
curl -X POST http://localhost:8080/api/certificates \
  -H "Content-Type: application/json" \
  -d '{
    "email": "user@example.com",
    "name": "João Silva", 
    "course": "Go Programming",
    "completion_date": "2024-01-15",
    "template_id": "default"
  }'
```

### Geração em lote via CSV:
```bash
curl -X POST http://localhost:8080/api/certificates/batch \
  -F "file=@certificates.csv"
```

**Formato do CSV:**
```csv
email,name,course,completion_date
user1@example.com,João Silva,Go Programming,2024-01-15
user2@example.com,Maria Santos,Web Development,2024-01-20
```

### Acessar certificado:
```bash
# HTML
http://localhost:8080/api/certificates/{uuid}.html

# PDF  
http://localhost:8080/api/certificates/{uuid}.pdf
```

## Templates JSON / JSON Templates

Os templates definem a estrutura e aparência dos certificados:

```json
{
  "id": "default",
  "name": "Template Padrão",
  "html_template": "<html><body><h1>Certificado</h1><p>Certificamos que {{.Name}} concluiu o curso {{.Course}} em {{.CompletionDate}}</p></body></html>",
  "fields": [
    {"name": "name", "type": "string", "required": true},
    {"name": "course", "type": "string", "required": true},
    {"name": "completion_date", "type": "date", "required": true}
  ]
}
```

## Testes / Tests

- **Unitários**: Framework padrão do Go (`testing`)
- **Funcionais**: Cypress (se aplicável)
- **Cobertura**: Validação de todos os endpoints e cenários

## Dependências / Dependencies

- **gin-gonic/gin** v1.10.1 - Framework web HTTP
- **google/uuid** v1.6.0 - Geração de identificadores únicos
- **jung-kurt/gofpdf** v1.16.2 - Geração de PDF nativo em Go

## Tecnologias / Technologies

- **Go 1.24+** - Linguagem principal
- **Gin-gonic** - Framework web REST API
- **UUID** - Identificação única de certificados
- **HTML/CSS** - Templates de certificados para web
- **gofpdf** - Geração nativa de PDF em Go (não requer dependências externas)
- **JSON** - Configuração de templates

## Funcionalidades Implementadas / Implemented Features

✅ **Geração de certificados únicos via API**
✅ **Geração em lote via upload de CSV**
✅ **Export para HTML com template personalizado**
✅ **Export para PDF com layout profissional**
✅ **Templates configuráveis via JSON**
✅ **Busca de certificados por email**
✅ **CRUD completo de templates**
✅ **Armazenamento em memória (para desenvolvimento)**
✅ **Testes unitários**
✅ **API REST documentada**

## Health Check

Verifique se a aplicação está funcionando:
```bash
curl http://localhost:8080/api/health
```

## Logs e Debugging

A aplicação usa o sistema de logs padrão do Gin. Para modo de produção:
```bash
export GIN_MODE=release
go run main.go
```