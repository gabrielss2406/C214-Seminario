package tests

import (
	"testing"
	todo "todolist/internal" // Ajuste o caminho de import conforme o nome do seu módulo

	"github.com/stretchr/testify/assert"
)

// TestAdicionar verifica se uma tarefa é adicionada corretamente
func TestAdicionar(t *testing.T) {
	lista := todo.ToDoList{}

	// Adiciona uma nova tarefa
	tarefa := lista.Adicionar("Estudar Go")

	// Verifica que a lista contém 1 item
	assert.Equal(t, 1, len(lista.Itens), "Esperado 1 tarefa na lista")

	// Verifica que a tarefa adicionada tem o nome correto
	assert.Equal(t, "Estudar Go", tarefa.Tarefa, "Esperado 'Estudar Go', obtido '%s'", tarefa.Tarefa)
}

// TestListar verifica se a lista de tarefas está correta
func TestListar(t *testing.T) {
	lista := todo.ToDoList{}
	lista.Adicionar("Estudar Go")
	lista.Adicionar("Praticar testes")

	// Obtém a lista de tarefas
	itens := lista.Listar()

	// Verifica que a lista contém 2 itens
	assert.Equal(t, 2, len(itens), "Esperado 2 tarefas na lista")
}

// TestMarcarComoFeito verifica se uma tarefa é marcada como feita corretamente
func TestMarcarComoFeito(t *testing.T) {
	lista := todo.ToDoList{}
	tarefa := lista.Adicionar("Fazer exercício")

	// Marca a tarefa como feita
	sucesso := lista.MarcarComoFeito(tarefa.ID)

	// Verifica que a tarefa foi marcada como feita
	assert.True(t, sucesso, "Esperado sucesso ao marcar a tarefa como feita")
	assert.True(t, lista.Itens[0].Feito, "Esperado que a tarefa estivesse marcada como feita")
}

// TestRemover verifica se uma tarefa é removida corretamente
func TestRemover(t *testing.T) {
	lista := todo.ToDoList{}
	tarefa := lista.Adicionar("Limpar o código")

	// Remove a tarefa
	sucesso := lista.Remover(tarefa.ID)

	// Verifica que a tarefa foi removida com sucesso
	assert.True(t, sucesso, "Esperado sucesso ao remover a tarefa")
	assert.Equal(t, 0, len(lista.Itens), "Esperado 0 tarefas na lista após remoção")
}
