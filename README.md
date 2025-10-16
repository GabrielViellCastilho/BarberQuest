# BarberQuest

## Descrição

Sistema de agendamento online para barbearias, permitindo gerenciar clientes, horários e serviços de forma prática e intuitiva.

> ⚠️ Observação: O foco principal do BarberQuest é o desenvolvimento de APIs em Go, gestão de dados e lógica de negócio. O front-end serve apenas como interface de teste e demonstração das funcionalidades.

## Tecnologias

* **Back-end:** Go
* **Banco de dados:** PostgreSQL
* **Containerização:** Docker
* **Front-end:** HTML, CSS, JavaScript

## Funcionalidades

* Criar, editar e cancelar agendamentos
* Visualizar horários disponíveis
* Sistema de login para clientes e administradores
* Cadastro de usuários e barbeiros
* Gerenciar horários diferenciados para feriados
* Criar e gerenciar serviços oferecidos

## Configuração do ambiente

1. Copie o arquivo `.env.template` para `.env`:

   ```bash
   cp .env.template .env
   ```
2. Preencha as variáveis de ambiente conforme necessário:

    * **Banco de dados (PostgreSQL)**
      `PGUSER`, `PGPASSWORD`, `PGHOST`, `PGPORT`, `PGDATABASE`
    * **E-mail (SMTP)**
      `SMTP_EMAIL`, `SMTP_PASSWORD`
    * **Autenticação e segurança**
      `JWT_SECRET_KEY`
    * **Logs**
      `LOG_OUTPUT`, `LOG_LEVEL`
    * **Informações da aplicação**
      `URL`, `ADMIN_EMAIL`, `ADMIN_PASSWORD`, `ADMIN_NAME`, `ADMIN_PHONE`

## Como rodar o projeto

1. Clone o repositório:

   ```bash
   git clone https://github.com/GabrielViellCastilho/BarberQuest.git
   cd BarberQuest
   ```
2. Inicie os containers com Docker Compose:

   ```bash
   docker compose up -d
   ```
3. Baixe as dependências do Go:

   ```bash
   go mod download
   ```
4. Execute o servidor:

   ```bash
   go run main.go
   ```

