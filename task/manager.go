package task

import (
	"fmt"
	"time"
)

// TaskManager representation
type TaskManager struct {
	tl *TaskList
}

// TaskManager contructor
func NewTaskManager() *TaskManager {
	return &TaskManager{
		tl: NewTaskList(),
	}
}

const errorMessage = "Error: "

// Create Task
func (tm *TaskManager) CreateTask(title, description string, priority uint8, dueDate time.Time) {
	task, err := tm.tl.AddTask(title, description, priority, dueDate)
	if err != nil {
		fmt.Println(errorMessage, err)
		return
	}
	fmt.Printf("Task created: %s | Due Date: %s\n",
        task.Title,
        task.DueDate.Format("02/01/2006 15:04"))
}

// List and show tasks
func (tm *TaskManager) ShowTask() {
	tasks := tm.tl.ListTasks()
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
	}
	for _, task := range tasks {
		fmt.Printf(
			"ID: %d | Title: %s | Status: %s | Priority: %d | Due Date: %s\n\n\n",
			task.ID,
			task.Title,
			task.Status,
			task.Priority,
            task.DueDate.Format("02/01/2006 15:04"),
		)
	}
}

// Update Status
func (tm *TaskManager) UpdateTask(id int, status Status) {
	err := tm.tl.UpdateTaskStatus(id, status)
	if err != nil {
		fmt.Println(errorMessage, err)
		return
	}
	fmt.Println("Task updated successfully.")
}

// Task deletion
func (tm *TaskManager) RemoveTask(id int) {
	err := tm.tl.DeleteTask(id)
	if err != nil {
		fmt.Println(errorMessage, err)
	}
	fmt.Println("Task removed successfully.")
}