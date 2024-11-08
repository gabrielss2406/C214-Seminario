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

	tarefa := lista.Adicionar("Estudar Go")

	assert.Equal(t, 1, len(lista.Itens), "Esperado 1 tarefa na lista")

	assert.Equal(t, "Estudar Go", tarefa.Tarefa, "Esperado 'Estudar Go', obtido '%s'", tarefa.Tarefa)
}

// TestListar verifica se a lista de tarefas está correta
func TestListar(t *testing.T) {
	lista := todo.ToDoList{}
	lista.Adicionar("Estudar Go")
	lista.Adicionar("Praticar testes")

	itens := lista.Listar()

	assert.Equal(t, 2, len(itens), "Esperado 2 tarefas na lista")
}

// TestMarcarComoFeito verifica se uma tarefa é marcada como feita corretamente
func TestMarcarComoFeito(t *testing.T) {
	lista := todo.ToDoList{}
	tarefa := lista.Adicionar("Fazer exercício")

	sucesso := lista.MarcarComoFeito(tarefa.ID)

	assert.True(t, sucesso, "Esperado sucesso ao marcar a tarefa como feita")
	assert.True(t, lista.Itens[0].Feito, "Esperado que a tarefa estivesse marcada como feita")
}

// TestRemover verifica se uma tarefa é removida corretamente
func TestRemover(t *testing.T) {
	lista := todo.ToDoList{}
	tarefa := lista.Adicionar("Limpar o código")

	sucesso := lista.Remover(tarefa.ID)

	assert.True(t, sucesso, "Esperado sucesso ao remover a tarefa")
	assert.Equal(t, 0, len(lista.Itens), "Esperado 0 tarefas na lista após remoção")
}
