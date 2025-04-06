package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

var m map[int64]string

// timer returns a function that prints the name argument and
// the elapsed time between the call to timer and the call to
// the returned function. The returned function is intended to
// be used in a defer statement:
//
//	defer timer("sum")()
func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	defer timer("main")()
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
