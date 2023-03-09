package main

import (
	"encoding/json"
	"fmt"
)

type Fruits struct {
	List []string `json:"results"`
}

func (f *Fruits) MarshalJSON() ([]byte, error) {
	resp := struct {
		Results      []string `json:"results"`
		TotalResults int      `json:"total_results"`
	}{
		Results:      f.List,
		TotalResults: len(f.List),
	}

	return json.Marshal(resp)
}

func main() {
	fruits := Fruits{
		List: []string{"banana", "orange", "apple", "grapes"},
	}

	data, err := json.Marshal(&fruits)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(data))
}
