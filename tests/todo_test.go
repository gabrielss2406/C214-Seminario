package tests

import (
	"testing"
	todo "todolist/internal" // Ajuste o caminho de import conforme o nome do seu módulo

	"github.com/stretchr/testify/assert"
)

// TestAdicionar verifica se uma tarefa é adicionada corretamente
func TestAdicionarNativo(t *testing.T) {
	lista := todo.ToDoList{}

	tarefa := lista.Adicionar("Estudar Go")

	if len(lista.Itens) != 1 {
		t.Errorf("Esperado 1 tarefa na lista, mas obteve %d", len(lista.Itens))
	}

	if tarefa.Tarefa != "Estudar Go" {
		t.Errorf("Esperado 'Estudar Go', obtido '%s'", tarefa.Tarefa)
	}
}

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

func TestAdicionarNativoNegativo(t *testing.T) {
	lista := todo.ToDoList{}

	// Adiciona uma tarefa
	lista.Adicionar("Estudar Go")

	// Verifica se a tarefa adicionada está incorreta (caso negativo simulado)
	assert.NotEqual(t, 0, len(lista.Itens), "A lista não está vazia")
	assert.NotEqual(t, "Tarefa Fake", lista.Itens[0].Tarefa, "O nome da tarefa não é 'Tarefa Fake'")
}

func TestAdicionarNegativo(t *testing.T) {
	lista := todo.ToDoList{}
	tarefa := lista.Adicionar("Estudar Go")

	// Verifica que a lista contém uma tarefa e não está vazia (caso negativo simulado)
	assert.NotEqual(t, 0, len(lista.Itens), "A lista contém uma tarefa")
	assert.NotEqual(t, "Outra Tarefa Fake", tarefa.Tarefa, "A tarefa adicionada não tem o nome 'Outra Tarefa Fake'")
}

func TestListarNegativo(t *testing.T) {
	lista := todo.ToDoList{}
	lista.Adicionar("Estudar Go")
	lista.Adicionar("Praticar testes")
	itens := lista.Listar()

	// Verifica que a lista contém 2 itens, conforme esperado no caso negativo
	assert.NotEqual(t, 0, len(itens), "A lista não está vazia")
	assert.NotEqual(t, 3, len(itens), "A lista não contém 3 tarefas")
}
