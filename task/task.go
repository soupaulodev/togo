package task

import (
	"errors"
	"sync"
	"time"
)

// Status representation
type Status string

const (
	StatusPending Status = "pending"
	StatusDone    Status = "done"
)

// Task representation
type Task struct {
	ID int
	Title string
	Description string
	Status Status
	Priority uint8
	CreatedAt time.Time
	UpdateAt time.Time
	DueDate time.Time
}

// Task List in memory (protected by a mutex)
type TaskList struct {
	mu sync.Mutex
	tasks map[int]*Task
	nextId int
}

// TaskList contructor
func NewTaskList() *TaskList {
	return &TaskList{
		tasks: make(map[int]*Task),
		nextId: 1,
	}
}

// Add a new task to the list
func (tl *TaskList) AddTask(title, description string, priority uint8, dueDate time.Time) (*Task, error) {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	if title == "" {
		return nil, errors.New("title cannot be empty")
	}

	now := time.Now()

	if !dueDate.IsZero() && dueDate.Before(now) {
        return nil, errors.New("due date cannot be in the past")
    }

	task := &Task{
		ID: tl.nextId,
		Title: title,
		Description: description,
		Status: StatusPending,
		Priority: priority,
		CreatedAt: now,
		UpdateAt: now,
		DueDate: dueDate,
	}

	tl.tasks[tl.nextId] = task
	tl.nextId++
	return task, nil
}

// Get all tasks
func (tl *TaskList) ListTasks() []*Task {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	tasks := []*Task{}
	for _, task := range tl.tasks {
		tasks = append(tasks, task)
	}
	return tasks
}

// Update status of a task
func (tl *TaskList) UpdateTaskStatus(id int, status Status) error {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	task, exists := tl.tasks[id]

	if !exists {
		return errors.New("task not found")
	}

	task.Status = status
	task.UpdateAt = time.Now()
	return nil
}

// Remove a task
func (tl *TaskList) DeleteTask(id int) error {
	tl.mu.Lock()
	defer tl.mu.Unlock()

	if _, exists := tl.tasks[id]; !exists {
		return errors.New("task not found")
	}

	delete(tl.tasks, id)
	return nil
}