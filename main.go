package main

// Initialize the application
func main() {
	tds := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&tds)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&tds)
	storage.Save(tds)
}