package Domain

import "time"

type Task struct {
	ID          string    
	Title       string    
	Description string    
	Status      string    
	DueDate     time.Time 
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