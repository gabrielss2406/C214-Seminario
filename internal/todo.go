package todo

type ToDo struct {
	ID     int
	Tarefa string
	Feito  bool
}

type ToDoList struct {
	Itens []ToDo
}

func (t *ToDoList) Adicionar(tarefa string) ToDo {
	novaTarefa := ToDo{ID: len(t.Itens) + 1, Tarefa: tarefa, Feito: false}
	t.Itens = append(t.Itens, novaTarefa)
	return novaTarefa
}

func (t *ToDoList) Listar() []ToDo {
	return t.Itens
}

func (t *ToDoList) MarcarComoFeito(id int) bool {
	for i := range t.Itens {
		if t.Itens[i].ID == id {
			t.Itens[i].Feito = true
			return true
		}
	}
	return false
}

func (t *ToDoList) Remover(id int) bool {
	for i := range t.Itens {
		if t.Itens[i].ID == id {
			t.Itens = append(t.Itens[:i], t.Itens[i+1:]...)
			return true
		}
	}
	return false
}
