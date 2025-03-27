package io

type Response struct {
	Status  bool
	Message string
	Data    []Detail
}

type Responsedetail struct {
	Status  bool
	Message string
	Data    []Detailstatus
}

type Detail struct {
	Id   string
	Info string
}

type Detailstatus struct {
	Username string
	Name     string
	Status   string
}
