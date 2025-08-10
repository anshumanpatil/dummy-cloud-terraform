package main

import "fmt"

type User struct {
	Email    string
	Password string
}

type Infrastructure struct {
	Name    string
	Type    string
	Details map[string]interface{}
}

type Database struct {
	Users           []*User
	Infrastructures []*Infrastructure
}

func NewDatabase() *Database {
	return &Database{
		Users:           []*User{},
		Infrastructures: []*Infrastructure{},
	}
}

func (db *Database) Flush() {
	if len(db.Users) > 5 {
		db.Users = []*User{}
	}

	if len(db.Infrastructures) > 35 {
		db.Infrastructures = []*Infrastructure{}
	}
	fmt.Println("Database flushed Users :- ", len(db.Users))
	fmt.Println("Database flushed Infrastructures :- ", len(db.Infrastructures))
}

func (db *Database) GetUser(email string) *User {
	for _, user := range db.Users {
		if user.Email == email {
			return user
		}
	}
	return nil
}

func (db *Database) AddUser(user *User) *User {
	db.Users = append(db.Users, user)
	return db.GetUser(user.Email)
}

func (db *Database) GetAllUsers() []*User {
	return db.Users
}

func (db *Database) AddInfrastructure(name string, typ string, details map[string]interface{}) *Infrastructure {
	infrastructure := &Infrastructure{
		Name:    name,
		Type:    typ,
		Details: details,
	}
	db.Infrastructures = append(db.Infrastructures, infrastructure)
	return infrastructure
}

func (db *Database) GetAllInfrastructures() []*Infrastructure {
	return db.Infrastructures
}

func (db *Database) GetInfrastructure(name string) *Infrastructure {
	for _, infra := range db.Infrastructures {
		if infra.Name == name {
			return infra
		}
	}
	return nil
}

func (db *Database) DeleteInfrastructure(name string) error {
	for i, infra := range db.Infrastructures {
		if infra.Name == name {
			db.Infrastructures = append(db.Infrastructures[:i], db.Infrastructures[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Infrastructure not found")
}

func (db *Database) UpdateInfrastructure(name string, typ string, details map[string]interface{}) *Infrastructure {
	infra := db.GetInfrastructure(name)
	if infra != nil {
		infra.Type = typ
		infra.Details = details
	}
	return infra
}
