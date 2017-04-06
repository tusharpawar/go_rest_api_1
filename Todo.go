package main 

import "time"

type Todo struct {
	Id int `json:"id"`
	Name string `json:"Name"`
	Completed bool `json:"completed"`
	Due time.Time `json:"due"`
}

type Todos []Todo