package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

type ModelMetadata struct {
	Name    string `json:"name"`
	Version string `json:"version"`
	Inputs  []struct {
		Name     string   `json:"name"`
		Datatype string   `json:"datatype"`
		Shape    []int    `json:"shape"`
	} `json:"inputs"`
	Outputs []struct {
		Name     string   `json:"name"`
		Datatype string   `json:"datatype"`
		Shape    []int    `json:"shape"`
	} `json:"outputs"`
}

func main() {
	
	modelName := os.Args[1]

	url := fmt.Sprintf("http://localhost:8000/v2/models/%s/versions/1", modelName)

	resp, err := http.Get(url)

	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return
	}

	var metadata ModelMetadata

	err = json.Unmarshal(body, &metadata)
	if err != nil {
		fmt.Printf("Error parsing response: %v\n", err)
		return
	}

	fmt.Println("Model Metadata:")

	fmt.Printf("Name: %s\n", metadata.Name)

	fmt.Printf("Version: %s\n", metadata.Version)
	
	fmt.Println("\nInputs:")

	for _, input := range metadata.Inputs {
		fmt.Printf("- Name: %s\n", input.Name)
		fmt.Printf("  Datatype: %s\n", input.Datatype)
		fmt.Printf("  Shape: %v\n", input.Shape)
	}
	
	fmt.Println("\nOutputs:")
	
	for _, output := range metadata.Outputs {
		fmt.Printf("- Name: %s\n", output.Name)
		fmt.Printf("  Datatype: %s\n", output.Datatype)
		fmt.Printf("  Shape: %v\n", output.Shape)
	}
}