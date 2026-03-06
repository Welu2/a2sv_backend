package data

import (
	"errors"
	"task_manager/models"
)

var tasks = []models.Task{}

func GetAllTasks() []models.Task {
	return tasks
}

func GetTaskByID(id string) (models.Task, error) {
	for _, t := range tasks {
		if t.ID == id {
			return t, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func CreateTask(task models.Task) models.Task {
	tasks = append(tasks, task)
	return task
}

func UpdateTask(id string, updatedTask models.Task) (models.Task, error) {
	for i, t := range tasks {
		if t.ID == id {
			updatedTask.ID = id
			tasks[i] = updatedTask
			return updatedTask, nil
		}
	}
	return models.Task{}, errors.New("task not found")
}

func DeleteTask(id string) error {
	for i, t := range tasks {
		if t.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("task not found")
}
