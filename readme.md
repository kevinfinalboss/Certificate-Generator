# Certificate-Generator 🎓

Certificate-Generator é uma aplicação serverless que permite criar e recuperar certificados digitais. Construída com serviços AWS, incluindo API Gateway, Lambda, DynamoDB e S3, esta aplicação fornece uma solução escalável e eficiente para gerenciamento de certificados.

## Índice
- [Funcionalidades](#funcionalidades)
- [Arquitetura](#arquitetura)
- [Como Usar](#como-usar)
- [Endpoints da API](#endpoints-da-api)
- [Exemplos de Implementação](#exemplos-de-implementação)
  - [Go](#go)
  - [Python](#python)
  - [Node.js](#nodejs)
- [Implementação na sua própria Infraestrutura AWS](#implementação-na-sua-própria-infraestrutura-aws)
- [Contribuindo](#contribuindo)
- [Licença](#licença)

## Funcionalidades

- 🚀 Arquitetura serverless para alta escalabilidade
- 🔒 Criação e recuperação segura de certificados
- 📊 DynamoDB para armazenamento eficiente de dados
- 🗃️ S3 para armazenamento de templates
- 🌐 API Gateway para fácil acesso aos endpoints

## Arquitetura

O Certificate-Generator utiliza os seguintes serviços AWS:

- API Gateway: Gerencia requisições HTTP
- Lambda: Processa requisições e gera certificados
- DynamoDB: Armazena dados dos certificados
- S3: Armazena templates HTML para os certificados

## Como Usar

Para usar o Certificate-Generator, você pode fazer requisições HTTP para os endpoints da API Gateway. 

### Criar um Certificado

Faça uma requisição POST para o endpoint `/certificates` com os dados do certificado no corpo da requisição:

```bash
curl -X POST https://certificates.kevindev.com.br/certificates \
-H "Content-Type: application/json" \
-d '{
  "participant_name": "João Silva",
  "company_name": "Empresa XYZ",
  "course_name": "Fundamentos de AWS",
  "total_hours": "40",
  "start_date": "2024-01-01",
  "end_date": "2024-01-30",
  "director_name": "Maria Santos",
  "coordinator_name": "Carlos Oliveira",
  "certificate_id": "123456789",
  "issue_date": "2024-08-20"
}'
```

### Recuperar um Certificado

Faça uma requisição GET para o endpoint `/certificates` com o UUID do certificado como parâmetro de consulta:

```bash
curl https://certificates.kevindev.com.br/certificates?uuid=a06f686c-a240-4987-9228-ad8cb0f02471
```

## Endpoints da API

- `POST /certificates`: Cria um novo certificado
- `GET /certificates?uuid={uuid}`: Recupera um certificado existente pelo UUID

## Exemplos de Implementação

### Go

```go
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func createCertificate() {
	url := "https://certificates.kevindev.com.br/certificates"
	data := map[string]string{
		"participant_name": "João Silva",
		"company_name":     "Empresa XYZ",
		"course_name":      "Fundamentos de Go",
		"total_hours":      "40",
		"start_date":       "2024-01-01",
		"end_date":         "2024-01-30",
		"director_name":    "Maria Santos",
		"coordinator_name": "Carlos Oliveira",
		"certificate_id":   "123456789",
		"issue_date":       "2024-08-20",
	}

	jsonData, _ := json.Marshal(data)
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println("Erro ao criar certificado:", err)
		return
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("Resposta:", string(body))
}

func main() {
	createCertificate()
}
```

### Python

```python
import requests
import json

def create_certificate():
    url = "https://certificates.kevindev.com.br/certificates"
    data = {
        "participant_name": "João Silva",
        "company_name": "Empresa XYZ",
        "course_name": "Fundamentos de Python",
        "total_hours": "40",
        "start_date": "2024-01-01",
        "end_date": "2024-01-30",
        "director_name": "Maria Santos",
        "coordinator_name": "Carlos Oliveira",
        "certificate_id": "123456789",
        "issue_date": "2024-08-20"
    }
    
    response = requests.post(url, json=data)
    print("Status:", response.status_code)
    print("Response:", response.json())

if __name__ == "__main__":
    create_certificate()
```

### Node.js

```javascript
const axios = require('axios');

async function createCertificate() {
    const url = 'https://certificates.kevindev.com.br/certificates';
    const data = {
        participant_name: "João Silva",
        company_name: "Empresa XYZ",
        course_name: "Fundamentos de Node.js",
        total_hours: "40",
        start_date: "2024-01-01",
        end_date: "2024-01-30",
        director_name: "Maria Santos",
        coordinator_name: "Carlos Oliveira",
        certificate_id: "123456789",
        issue_date: "2024-08-20"
    };

    try {
        const response = await axios.post(url, data);
        console.log("Status:", response.status);
        console.log("Data:", response.data);
    } catch (error) {
        console.error("Erro ao criar certificado:", error);
    }
}

createCertificate();
```

## Implementação na sua própria Infraestrutura AWS

Para implementar o Certificate-Generator na sua própria infraestrutura AWS, siga estes passos:

1. Clone o repositório:
   ```
   git clone https://github.com/seu-usuario/certificate-generator.git
   cd certificate-generator
   ```

2. Configure as variáveis de ambiente no seu ambiente AWS:
   - `DYNAMODB_TABLE_NAME`: Nome da tabela DynamoDB para armazenar os certificados
   - `S3_BUCKET_NAME`: Nome do bucket S3 para armazenar os templates HTML
   - `S3_TEMPLATE_KEY`: Caminho do template HTML no bucket S3

3. Deploy da função Lambda:
   - Compile o código Go para Linux:
     ```
     GOOS=linux GOARCH=amd64 go build -o main main.go
     ```
   - Crie um arquivo ZIP com o executável:
     ```
     zip deployment.zip main
     ```
   - Faça upload do arquivo ZIP para o Lambda através do console AWS ou AWS CLI

4. Configure o API Gateway:
   - Crie uma nova API HTTP
   - Configure as rotas para POST e GET em `/certificates`
   - Integre as rotas com a função Lambda

5. Configure as permissões necessárias:
   - Crie uma role IAM para a função Lambda com permissões para acessar DynamoDB e S3
   - Configure as políticas de CORS no API Gateway

6. Teste a implementação usando os exemplos de código fornecidos acima

## Contribuindo

Contribuições são bem-vindas! Por favor, leia o arquivo CONTRIBUTING.md para detalhes sobre nosso código de conduta e o processo para enviar pull requests.

## Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE.md](LICENSE.md) para detalhes.