package Domain

import "time"

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	DueDate     time.Time `json:"due_date"`
}

// Constructor
func NewTask(id, title, description string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      "pending",
	}
}

// Update
func (t *Task) Update(title, description, status string) {
	if title != "" {
		t.Title = title
	}
	if description != "" {
		t.Description = description
	}
	if status != "" {
		t.Status = status
	}
}

// ------------------- TASK -------------------
// type Task struct {
// 	id          string
// 	title       string
// 	description string
// 	status      string
// }

// func NewTask(id, title, description string) *Task {
// 	return &Task{
// 		id:          id,
// 		title:       title,
// 		description: description,
// 		status:      "pending",
// 	}
// }

// Getters
// func (t *Task) ID() string          { return t.id }
// func (t *Task) Title() string       { return t.title }
// func (t *Task) Description() string { return t.description }
// func (t *Task) Status() string      { return t.status }

// // Setters
// func (t *Task) SetID(id string) { t.id = id }

// Update
// func (t *Task) Update(title, description, status string) {
// 	if title != "" {
// 		t.title = title
// 	}
// 	if description != "" {
// 		t.description = description
// 	}
// 	if status != "" {
// 		t.status = status
// 	}
// }

// ------------------- USER -------------------
type User struct {
	id       string
	username string
	password string
	role     string
}

func NewUser(id, username, password, role string) *User {
	return &User{
		id:       id,
		username: username,
		password: password,
		role:     role,
	}
}

// Getters
func (u *User) ID() string       { return u.id }
func (u *User) Username() string { return u.username }
func (u *User) Password() string { return u.password }
func (u *User) Role() string     { return u.role }

// Setters
func (u *User) SetID(id string)       { u.id = id }
func (u *User) SetPassword(pw string) { u.password = pw }
func (u *User) SetRole(role string)   { u.role = role }