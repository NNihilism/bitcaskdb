package main

import (
	"fmt"
	"os"

	"github.com/NNihilism/bitcaskdb"
	"github.com/NNihilism/bitcaskdb/options"
)

func main() {
	path := "D:" + string(os.PathSeparator) + "test"
	opts := options.DefaultOptions(path)
	_, err := bitcaskdb.Open(opts)
	if err != nil {
		fmt.Printf("open bitcask err: %v", err)
		return
	}

}
