package main

import (
	"fmt"
	"os"

	"github.com/NNihilism/bitcaskdb/bitcask"
	"github.com/NNihilism/bitcaskdb/options"
)

func main() {
	path := "D:" + string(os.PathSeparator) + "test"
	opts := options.DefaultOptions(path)
	_, err := bitcask.Open(opts)
	if err != nil {
		fmt.Printf("open bitcask err: %v", err)
		return
	}

}
