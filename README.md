# ToDo List em Go

Um simples aplicativo de ToDo List desenvolvido em Go. Este projeto permite adicionar tarefas, listar tarefas e marcar tarefas como concluídas.

## Tecnologias Utilizadas

- Go (Golang)
- Gorilla Mux (para roteamento)
- Testes utilizando o pacote `testing` do Go

## Instalação

1. **Certifique-se de que o Go está instalado**:
 Você pode verificar a instalação do Go executando `go version` no terminal.

2. **Clone o repositório**:
   ```bash
   git clone <URL do repositório>
   cd todolist
   ```

3. **Instale as dependências**:
   ```bash
   go mod tidy
   ```
## Rodando localmente

Já estando dentro da pasta do projeto, inicie o server com

```bash
  go run main.go
```
OBS: essa maneira de rodar não implementar hot reload, caso queira, configure o ```aix```.
## Endpoints da API
**Adicionar Tarefa**: POST /todos
```
{
  "tarefa": "Sua nova tarefa"
}
```
**Listar Tarefas**: GET /todos

**Marcar Tarefa como Feita**: PUT /todos/{id}

## Rodando os testes

Ainda não sei Ewel

