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
	return ToDo{ID: len(t.Itens) + 1, Tarefa: tarefa, Feito: false}
}

func (t *ToDoList) Listar() []ToDo {
	return t.Itens
}

func (t *ToDoList) MarcarComoFeito(id int) bool {
	return true
}
