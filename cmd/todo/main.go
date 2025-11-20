package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID			int       `json:"id"`
	Title		string    `json:"title"`
	Done		bool      `json:"done"`
	CreatedAt	time.Time `json:"created_at"`
}

type Tasks []Task

const dataFile = "../../tasks.json"

func loadTasks() Tasks {
	var tasks Tasks
	data, err := os.ReadFile(dataFile)
	if err != nil {
		return tasks
	}
	json.Unmarshal(data, &tasks)
	return tasks
}

func saveTasks(tasks Tasks) {
	data, _ := json.MarshalIndent(tasks, "", "  ")
	os.WriteFile(dataFile, data, 0644)
}

func addTask(title string) {
	tasks := loadTasks()
	newTask := Task{
		ID:        	len(tasks) + 1,
		Title:    	strings.TrimSpace(title),
		Done:	 	false,
		CreatedAt:	time.Now(),
	}
	tasks = append(tasks, newTask)
	saveTasks(tasks)
	fmt.Println("Todo berhasil ditambahkan!")
}

func listTasks() {
	tasks := loadTasks()
	if len(tasks) == 0 {
		fmt.Println("Yeay! Belum ada todo, saatnya santai")
		return
	}
	fmt.Println("Daftar Todo :")
	fmt.Println(strings.Repeat("=", 30))
	for _, t := range tasks {
		status := "[ ]"
		if t.Done {
			status = "[x]"
		}
		fmt.Printf("%d. %s %s\n", t.ID, status, t.Title)
	}
	fmt.Print(strings.Repeat("=", 30))
}

func doneTask(id int) {
	tasks := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			saveTasks(tasks)
			fmt.Println("Todo berhasil ditandai selesai!")
			return
		}
	}
	fmt.Println("Todo dengan ID tersebut tidak ditemukan.")
}

func deleteTask(id int) {
	tasks := loadTasks()
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			for j := i; j < len(tasks); j++ {
				tasks[j].ID = j + 1
			}
			saveTasks(tasks)
			fmt.Println("Todo berhasil dihapus!")
			return
		}
	}
	fmt.Println("Todo dengan ID tersebut tidak ditemukan.")
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(`
CLI Todo List (Go)
Cara pakai:
  go run . add "Belajar Go 1 jam"
  go run . list
  go run . done 1
  go run . delete 2
		`)
		return
	}

	cmd := os.Args[1]

	switch cmd {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Error: tulis judul todo!")
			return
		}
		title := strings.Join(os.Args[2:], " ")
		addTask(title)

	case "list":
		listTasks()

	case "done":
		if len(os.Args) != 3 {
			fmt.Println("Error: go run . done <nomor>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Nomor harus angka!")
			return
		}
		doneTask(id)

	case "delete":
		if len(os.Args) != 3 {
			fmt.Println("Error: go run . delete <nomor>")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Nomor harus angka!")
			return
		}
		deleteTask(id)

	default:
		fmt.Println("Perintah salah. Pilihan: add | list | done | delete")
	}
}