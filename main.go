package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
)

func home(res http.ResponseWriter, req *http.Request) {

}

// type Todo struct {
// 	Title   string
// 	Content string
// }

// var todos []Todo

// type PageVariables struct {
// 	PageTitle string
// 	PageTodos []Todo
// }

// func getTodos(w http.ResponseWriter, r *http.Request) {
// 	pageVariables := PageVariables{
// 		PageTitle: "Get Todos",
// 		PageTodos: todos,
// 	}

// 	t, err := template.ParseFiles("todos.html")

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		log.Print("Template parsing error: ", err)
// 	}

// 	err = t.Execute(w, pageVariables)
// 	_ = err
// }

// func addTodo(w http.ResponseWriter, r *http.Request) {
// 	err := r.ParseForm()
// 	fmt.Println("Add todo!")

// 	if err != nil {
// 		http.Error(w, err.Error(), http.StatusBadRequest)
// 		log.Print("Request parsing error: ", err)
// 	}

// 	todo := Todo{
// 		Title:   r.FormValue("title"),
// 		Content: r.FormValue("content"),
// 	}

// 	todos = append(todos, todo)
// 	log.Print(todos)
// 	http.Redirect(w, r, "/todos/", http.StatusSeeOther)
// }

type Contact struct {
	ID        int
	FirstName string
	LastName  string
	Phone     string
	Email     string
	Notes     string
}

type ContactsPageVariables struct {
	PageTitle string
	Contacts  []Contact
}

type ContactBook struct {
	Contacts []Contact
}

type Searchable interface {
	search(search string) []Searchable
}

var contactBook ContactBook = ContactBook{
	Contacts: []Contact{
		{ID: 1, FirstName: "Jane", LastName: "Doe", Phone: "987654321", Email: "janedoe@gmail.com"},
		{ID: 2, FirstName: "John", LastName: "Doe", Phone: "123456789", Email: "johndoe@gmail.com"},
	},
}

func (contactBook *ContactBook) search(search string) []Contact {
	if search == "" {
		return contactBook.Contacts
	}

	var result []Contact

	for _, contact := range contactBook.Contacts {
		if strings.Contains(contact.FirstName, search) || strings.Contains(contact.LastName, search) || strings.Contains(contact.Phone, search) || strings.Contains(contact.Email, search) || strings.Contains(contact.Notes, search) {
			result = append(result, contact)
		}
	}

	return result
}

func contactsPage(w http.ResponseWriter, r *http.Request) {
	search := r.URL.Query().Get("search")

	pageVariables := ContactsPageVariables{
		PageTitle: "Contacts",
		Contacts:  contactBook.search(search),
	}

	t, err := template.ParseFiles("contacts.html")

	fmt.Println(pageVariables)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Print("Template parsing error: ", err)
	}

	err = t.Execute(w, pageVariables)
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/contacts/", contactsPage)
	// http.HandleFunc("/todos/", getTodos)
	// http.HandleFunc("/add-todo/", addTodo)
	fmt.Println("Server is running on port :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
