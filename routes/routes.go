package routes

import (
	"encoding/json"
	"net/http"
	"strconv"
	todo "todolist/internal"

	"github.com/gorilla/mux"
)

var lista = todo.ToDoList{}

func AdicionarToDoHandler(w http.ResponseWriter, r *http.Request) {
	var novaTarefa todo.ToDo
	if err := json.NewDecoder(r.Body).Decode(&novaTarefa); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	lista.Adicionar(novaTarefa.Tarefa)
	w.WriteHeader(http.StatusCreated)
}

func ListarToDosHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(lista.Listar())
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func MarcarComoFeitoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if sucesso := lista.MarcarComoFeito(id); !sucesso {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func RemoverToDoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}
	if sucesso := lista.Remover(id); !sucesso {
		http.Error(w, "Tarefa não encontrada", http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusOK)
	}
}

func SetupRoutes() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/todos", AdicionarToDoHandler).Methods("POST")
	r.HandleFunc("/todos", ListarToDosHandler).Methods("GET")
	r.HandleFunc("/todos/{id:[0-9]+}", MarcarComoFeitoHandler).Methods("PUT")
	r.HandleFunc("/todos/{id:[0-9]+}", RemoverToDoHandler).Methods("DELETE")
	return r
}
