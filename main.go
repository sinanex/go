// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"
// )

// func main() {
// 	fmt.Println(add(18, 34))

// 	var age int = 19
// 	fmt.Println(age)

// 	var name string = "Sinan"
// 	fmt.Println(message(name))

// 	nameprint()
// }

// func add(a int, b int) string {
// 	sum := a + b
// 	return fmt.Sprintf("Result: %d + %d = %d", a, b, sum)
// }

// func message(name string) string {
// 	return fmt.Sprintf("Hello %s", name)
// }

// func nameprint() {
// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Println("Enter your name:")

// 	name, _ := reader.ReadString('\n')
// 	name = strings.TrimSpace(name)

// 	fmt.Printf("Name is %s\n", name)
// }

package main

import (
	"fmt"
	"net/http"

	"encoding/json"
	"sync"
)

type Task struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

var (
	tasks  = []Task{}
	nextID = 1
	mutex  = &sync.Mutex{}
)

func main() {
	http.HandleFunc("/tasks", handleTask)
	fmt.Println("server running")
	http.ListenAndServe(":8000", nil)
}

func handleTask(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet {
		json.NewEncoder(w).Encode(tasks)
		return
	}
	if r.Method == http.MethodPost {
		var task Task
		err := json.NewDecoder(r.Body).Decode(&task)
		if err != nil {
			http.Error(w, "Invalid data", http.StatusBadRequest)
			return
		}

		mutex.Lock()
		task.ID = nextID
		nextID++
		tasks = append(tasks, task)
		mutex.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(task)
		return
	}

	http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)

}
