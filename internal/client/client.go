package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"
)

type EmployeeClient struct {
	addr string
}

func NewEmployeeClient(addr string) *EmployeeClient {
	return &EmployeeClient{addr: addr}
}

func (x *EmployeeClient) GetEmployeeById(id string) (*Employee, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api/employees/%v", x.addr, id))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	res := &Employee{}
	return res, json.Unmarshal(body, res)
}

type empErrTuple struct {
	emp *Employee
	err error
}

func (x *EmployeeClient) GetEmployeesByIds(ids []string) ([]*Employee, []error) {
	var wg sync.WaitGroup
	empChan := make(chan *empErrTuple)

	for _, id := range ids {
		wg.Add(1)
		go func(id string) {
			defer wg.Done()
			emp, err := x.GetEmployeeById(id)
			empChan <- &empErrTuple{emp: emp, err: err}
		}(id)
	}
	go func() {
		wg.Wait()
		close(empChan)
	}()
	emps := make([]*Employee, 0, len(ids))
	errs := make([]error, 0, len(ids))
	for tuple := range empChan {
		if tuple.err != nil {
			errs = append(errs, tuple.err)
		} else {
			emps = append(emps, tuple.emp)
		}
	}
	return emps, errs
}
