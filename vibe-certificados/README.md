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

## Instalação e Execução / Installation and Execution

```bash
# Instalar dependências
go mod download

# Executar aplicação
go run main.go

# Executar testes
go test ./...
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

## Tecnologias / Technologies

- **Go 1.20+**
- **Gin-gonic**: Framework web
- **UUID**: Identificação única
- **HTML/CSS**: Templates de certificados
- **PDF**: Geração via bibliotecas Go
- **Swagger**: Documentação da API