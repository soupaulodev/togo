package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title	   	string
	Completed	bool
	CreatedAt	time.Time
	CompletedAt	*time.Time
}

type Todos []Todo

func (tds *Todos) add(title string) {
	todo := Todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*tds = append(*tds, todo)
}

func (tds *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*tds) {
		err := errors.New("Invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (tds *Todos) delete(index int) error {
	t := *tds

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*tds = append(t[:index], t[index+1:]...)
	return nil
}

func (tds *Todos) toggle(index int) error {
	t := *tds

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isCompleted := t[index].Completed
	
	if !isCompleted {
		completionTime := time.Now()
		t[index].CompletedAt = &completionTime
	}
	t[index].Completed = !isCompleted
	return nil
}

func (tds *Todos) edit(index int, title string) error {
	t := *tds

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Title = title
	return nil
}

func (tds *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for i, t := range *tds {
		completed := "No"
		completedAt := ""

		if t.Completed {
			completed = "Yes"
			if t.CompletedAt != nil {
				completedAt = t.CompletedAt.Format("2006-01-02 15:04:05")
			}
		}

		table.AddRow(strconv.Itoa(i), t.Title, completed, t.CreatedAt.Format("2006-01-02 15:04:05"), completedAt)
	}

	table.Render()
}