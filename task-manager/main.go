package main

import (
	"errors"
	"fmt"
)

type Task struct {
	Id          int
	Description string
}

type TaskManager struct {
	Tasks []Task
}

func main() {
	var operation int
	manager := &TaskManager{}

	for {
		fmt.Println("\nQual operação deseja realizar? \n 1 - Adicionar Task \n 2 - Remover Task \n 3 - Listar Tasks \n 4 - Encerrar programa \n ")
		fmt.Scanln(&operation)

		if operation == 1 {
			fmt.Println("\nAdicionando tarefa")
			if err := manager.AddTask(); err != nil {
				fmt.Println(err)
				break
			}
		}

		if operation == 2 {
			fmt.Println("\nRemovendo task")
			if err := manager.RemoveTask(); err != nil {
				fmt.Println(err)
			}
		}
		if operation == 3 {
			fmt.Println("\nListando tarefas existentes")
			if err := manager.ListTask(); err != nil {
				fmt.Println(err)
			}
		}
		if operation == 4 {
			fmt.Println("\nEncerrando programa!")
			break
		}

		continue
	}
}

func (tm *TaskManager) AddTask() error {
	var desc string
	var id int
	fmt.Println("Insira o ID da task: ")
	fmt.Scanln(&id)
	fmt.Println("\nInsira a descrição da tarefa: ")
	fmt.Scanln(&desc)

	for _, task := range tm.Tasks {
		if task.Id == id {
			return errors.New("\nErro: ID de tarefa já existente\n")
		}
	}
	task := Task{Id: id, Description: desc}
	tm.Tasks = append(tm.Tasks, task)
	fmt.Println("\nTask adicionada com sucesso!\n")

	return nil
}

func (tm *TaskManager) RemoveTask() error {
	var id int

	fmt.Println("\nInsira o ID da task que deseja remover: ")
	fmt.Scanln(&id)

	index := -1
	for i, task := range tm.Tasks {
		if task.Id == id {
			index = i
			break
		}
	}

	if index == -1 {
		return errors.New("\nErro: tarefa não encontrada! Tente novamente")
	}

	// Remove o elemento do slice
	tm.Tasks = append(tm.Tasks[:index], tm.Tasks[index+1:]...)
	fmt.Println("Task removida com sucesso! \n")

	return nil

}

func (tm *TaskManager) ListTask() error {
	if tm.Tasks == nil { // if len(tmTasks == 0)
		return errors.New("\nErro: Sem tarefas existentes\n")
	}
	fmt.Println(tm.Tasks, "\n")
	return nil
}
