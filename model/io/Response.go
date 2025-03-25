package io

type Response struct {
	Status  bool
	Message string
	Data    []Detail
}

type Detail struct {
	Id   string
	Info string
}
