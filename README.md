# API de Estudos em Go

Esta é uma API simples desenvolvida em Go, projetada para praticar o desenvolvimento de APIs RESTful.

## Visão Geral

A API consulta uma tabela chamada "clientes" em um banco de dados, onde cada cliente possui os campos "id", "nome" e "idade".

## Como Executar

1. Certifique-se de ter o Go instalado em seu sistema. Você pode baixá-lo em [golang.org](https://golang.org/).

2. Clone este repositório: git clone https://github.com/gabrielcamurcab/desafio-apirestful-go-dio

3. Navegue até o diretório do projeto: cd desafio-apirestful-go-dio

4. Execute o seguinte comando para iniciar o servidor: go run main.go


Agora a API estará sendo executada em `http://localhost:3000`.

## Rotas Disponíveis

- **GET /client**: Retorna todos os clientes da tabela "clientes".
- **GET /client/{id}**: Retorna um cliente específico com o ID fornecido.
- **POST /client**: Cria um novo cliente com base nos dados fornecidos.
- **DELETE /client/{id}**: Exclui um cliente específico com o ID fornecido.
- **PUT /client/{id}**: Atualiza os dados de um cliente específico com o ID fornecido.

Lembre-se de substituir `{id}` pelo ID real do cliente ao fazer requisições para as rotas que exigem um ID.
