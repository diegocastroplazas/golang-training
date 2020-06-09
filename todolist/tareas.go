package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
}

func (t *taskList) imprimirLista() {
	for _, tarea := range t.tasks {
		fmt.Println("Nombre: ", tarea.nombre)
		fmt.Println("Descripción: ", tarea.descripcion)
	}
}

func (t *taskList) imprimirListaCompletados() {
	for _, tarea := range t.tasks {
		if tarea.completado == true {
			fmt.Println("Nombre: ", tarea.nombre)
			fmt.Println("Descripción: ", tarea.descripcion)
		}
	}
}

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompleta() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {
	t := &task{
		nombre:      "Hola",
		descripcion: "Quiubo",
		completado:  true,
	}
	t1 := &task{
		nombre:      "Hola2",
		descripcion: "Python",
	}
	t2 := &task{
		nombre:      "Hola3",
		descripcion: "Golang",
	}

	t3 := &task{
		nombre:      "Hola3",
		descripcion: "Esta es una nueva tarea",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2, t,
		},
	}
	fmt.Println(lista.tasks[0])

	lista.agregarALista(t3)
	fmt.Println(len(lista.tasks))
	lista.imprimirLista()
	fmt.Println("Tareas completadas: ")
	lista.imprimirListaCompletados()

	mapaTareas := make(map[string]*taskList)
	mapaTareas["Diego"] = lista

	fmt.Println("Tareas de Diego: ")
	mapaTareas["Diego"].imprimirLista()
}
