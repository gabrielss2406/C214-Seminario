package tests

import (
	"testing"
	todo "todolist/internal" // Ajuste o caminho de import conforme o nome do seu módulo

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type ToDo struct {
	ID     int
	Tarefa string
	Feito  bool
}

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

// Mock da estrutura ToDoList
type MockToDoList struct {
	mock.Mock
}

func (m *MockToDoList) Listar() []ToDo {
	args := m.Called()
	return args.Get(0).([]ToDo)
}

// Teste com mock
func TestListarComMock(t *testing.T) {
	// Cria o mock
	mockLista := new(MockToDoList)

	// Definindo comportamento esperado para o mock do método Listar
	mockLista.On("Listar").Return([]ToDo{
		{ID: 1, Tarefa: "Estudar Go", Feito: false},
		{ID: 2, Tarefa: "Praticar testes", Feito: false},
	})

	// Obtém a lista de tarefas
	itens := mockLista.Listar()

	// Verifica se o mock foi chamado corretamente
	mockLista.AssertExpectations(t)

	// Verifica se a lista contém 2 tarefas
	assert.Equal(t, 2, len(itens), "Esperado 2 tarefas na lista")
	assert.Equal(t, "Estudar Go", itens[0].Tarefa, "Esperado tarefa 'Estudar Go'")
	assert.Equal(t, "Praticar testes", itens[1].Tarefa, "Esperado tarefa 'Praticar testes'")
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

func TestMarcarComoFeitoNegativo(t *testing.T) {
	lista := todo.ToDoList{}
	tarefa := lista.Adicionar("Fazer exercício")
	// Como a tarefa não foi feita, o status 'feito' tem que estar desmarcado
	assert.False(t, tarefa.Feito, "O status 'Feito' não está marcado")
}

func TestRemoverNegativo(t *testing.T) {
	lista := todo.ToDoList{}
	tarefa := lista.Adicionar("Limpar o código")
	// Tenta remover uma tarefa com ID inexistente
	sucesso := lista.Remover(tarefa.ID + 1)
	// Verifica que a tarefa não foi removida
	assert.False(t, sucesso, "Esperado falha ao tentar remover a tarefa")
	assert.Equal(t, 1, len(lista.Itens), "Esperado 1 tarefa na lista após tentativa de remoção incorreta")
}
