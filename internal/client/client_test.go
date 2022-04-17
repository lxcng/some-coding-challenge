package client

import (
	"fmt"
	"testing"
)

const (
	addr = "https://employees-api.vercel.app"

	id0 = "6eb0ff94-5d64-4175-95fa-c5a30c923e28"
	id1 = "0432526b-6bce-4b13-afec-5a6754f70500"
	id2 = "abfbd243-c699-4d82-a9c9-f92e5936ee10"
)

func TestGetEmp(t *testing.T) {
	cl := NewEmployeeClient(addr)
	res, err := cl.GetEmployeeById(id0)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(res)
}

func TestGetEmps(t *testing.T) {
	cl := NewEmployeeClient(addr)
	res, errs := cl.GetEmployeesByIds([]string{id0, id1, id2})
	for _, err := range errs {
		t.Error(err)
	}
	for _, r := range res {
		fmt.Println(r)
	}
}
