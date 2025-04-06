package main

import (
	"fmt"
	"os"
	"path/filepath"
)

var m map[int64]string

func main() {
	d := `C:\Users\nolan\Desktop\shows`
	m = make(map[int64]string)

	err := filepath.Walk(d, func(path string, f os.FileInfo, err error) error {
		stats, statErr := os.Stat(path)

		if statErr != nil {
			panic(statErr)
		}

		ext := filepath.Ext(path)
		if ext != ".mkv" {
			return nil
		}

		val, ok := m[stats.Size()]

		if ok {
			fmt.Printf("Collision in Sizes Val: %s Size: %d Val: %s\n", val, stats.Size(), path)
			return nil
		}

		m[stats.Size()] = path

		//		fmt.Printf("Base: %s Ext: %s Size: %d\n", stats.Name(), ext, stats.Size())
		return err
	})
	if err != nil {
		panic(err)
	}
}
