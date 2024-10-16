package main

import "fmt"


type Todo struct {
	Id        int
	Title       string
	Description string
	Completed   bool
}

var todoList []Todo


func createTodo(title, description *string) string {
	fmt.Print("Enter a todo title: ")
	fmt.Scan(title)

	fmt.Print("Enter description for this todo: ")
	fmt.Scan(description)

	newId := len(todoList) + 1
	newTodo := Todo{ 
		Id: newId,
		Title: *title,
		Description: *description,
		Completed: false,
	}
	todoList = append(todoList, newTodo)

	return "Todo has been added successfully"
}

func getAllTodos() {
	if len(todoList) == 0 {
		fmt.Println("No todos found.")
		return
	}

	fmt.Println("Todo List:")
	for _, todo := range todoList {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Completed: %v\n",
			todo.Id, todo.Title, todo.Description, todo.Completed)
	}
}

func updateTodo(id *uint, title, description *string, completed *bool) {

	var userResponse uint
	
	fmt.Print("Enter todo ID: ")
	fmt.Scan(id)

	if (*id > uint(len(todoList))) {
		fmt.Println("Invalid Id, check Id and try again")
		return
	}

	for {

		fmt.Println("What do you want to update ?")
		fmt.Println(`
			1. title
			2. description
			3. todo_status
			4. exit
		`)

		fmt.Scan(&userResponse)

		if userResponse == 4 {
			break;
		}

		if userResponse == 1 {
			fmt.Print("Updated title: ")
			fmt.Scan(title)
			todoList[*id - 1] = Todo{Title: *title, Description: *description, Completed: *completed}
		} else if userResponse == 2 {
			fmt.Print("Updated description: ")
			fmt.Scan(description)
			todoList[*id - 1] = Todo{Description: *description, Completed: *completed, Title: *title}
		} else if userResponse == 3 {
			fmt.Print("Updated status: ")
			fmt.Scan(completed)
			todoList[*id - 1] = Todo{Completed: *completed,Title: *title, Description: *description}
			
		} else {
			fmt.Println("Invalid input, try again")
			continue
		}
		
	}
}

func deleteTodo(id *uint) {
	var userResponse uint
	
	
	for {
		fmt.Print("Enter todo ID: ")
		fmt.Scan(id)
	
		if (*id > uint(len(todoList))) {
			fmt.Println("Invalid Id, check Id and try again")
			return
		}

		// todoList = append(todoList[:i], todoList[i+1:]...)
		todoList = append(todoList[:*id], todoList[int(*id)+1])
		fmt.Printf("The todo with %v has been deleted successfully", *id)

		fmt.Println("Is that all ?")
		fmt.Println(`
			1. yes
			2. no
		`)

		fmt.Scan(&userResponse)

		if userResponse == 1 {
			break;
		} else if userResponse == 2 {
			continue
		} else {
			fmt.Println("Invalid input")
		}
	}

}

func collectUserResponse(userInput *uint) string {
	for {
		fmt.Print("Option: ")
		_, err:= fmt.Scan(userInput)
	
		if err != nil {
			fmt.Println("Error reading input, check input and try again")
			continue
		}
		break
	}
	result := fmt.Sprintf(`You selected %v`, *userInput)
	return result
}



func main() {
	var userResponse uint
	var title string
	var description string
	var id uint
	var completed bool

	for {
		fmt.Println("What action are you trying to perform ?")
		fmt.Println(`
			1. Add an item into todo list
			2. Get all todos
			3. Update an item from the todo list
			4. Delete an item from the todo list
			5. Quit
		`)

		fmt.Println(collectUserResponse(&userResponse))

		if userResponse == 5 {
			fmt.Println(`Quiting program...`)
			break
		}
	
		if userResponse == 1 {
			fmt.Println(createTodo(&title, &description))
		}
	
		if userResponse == 2 {
			getAllTodos()
		}

		if userResponse == 3 {
			updateTodo(&id, &title, &description, &completed)
		}

		if userResponse == 4 {
			deleteTodo(&id)
		}


		for i, value := range todoList {
			fmt.Println(i, value);
		}

	}


	
}