package todo

import (
	"fmt"
	"os"

	"github.com/jmutschler/kv"
)

type TodoStatus string

const (
	TodoStatusDone TodoStatus = "done"
)

type Todo struct {
	ID     string
	Status TodoStatus
}

type Catalog struct {
	store *kv.Store[Todo]
}

func Main() int {
	store, err := kv.OpenStore[Todo]("todo.db")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return 1
	}
	todos := List(store)
	return 0
}

func List(store *kv.Store[Todo]) []Todo {
	data := store.All()

	todos := make([]Todo, 0, len(data))
	for _, todo := range data {
		todos = append(todos, todo)
	}
	return todos
}

func (c *Catalog) NewTodo(todo Todo) {
	todo.ID = c.nextID()
	c.store.Set(todo.ID, todo)

}

func (c *Catalog) CompleteTodo(id string) bool {
	t, ok := c.store.Get(id)
	if !ok {
		fmt.Printf("No todo with ID %s\n", id)
		return false
	}
	t.Status = TodoStatusDone
}
