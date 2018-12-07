package main

import (
	"time"
)

type Output struct {
	name        string
	division    string
	status      string
	currentTime time.Time
}

type Division struct {
	id   int
	name string
}

type Employee struct {
	id        int
	name      string
	division  string
	status    string
	createdAt *time.Time
}
