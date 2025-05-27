package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type recorded struct {
	Reply []byte `json:"reply"`
	Code  int    `json:"code"`
}

func errorMain() error {
	if len(os.Args) != 2 {
		return fmt.Errorf("usage: %s <data-path>", os.Args[0])
	}

	dataPath := os.Args[1]
	data, err := os.ReadFile(dataPath)
	if err != nil {
		return fmt.Errorf("reading data file %q: %w", dataPath, err)
	}
	var rec recorded
	if err := json.Unmarshal(data, &rec); err != nil {
		return fmt.Errorf("unmarshalling data from %q: %w", dataPath, err)
	}
	fmt.Printf("Code: %d\n", rec.Code)
	if len(rec.Reply) > 0 {
		fmt.Printf("Reply: %s\n", rec.Reply)
	} else {
		fmt.Println("Reply: <empty>")
	}

	return nil
}

func main() {
	if err := errorMain(); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
