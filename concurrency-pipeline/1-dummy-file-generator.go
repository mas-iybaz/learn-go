package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const TOTAL_FILE = 3000
const CONTENT_LENGTH = 5000

var tempPath = filepath.Join(os.Getenv("TEMP"), "concurrency-pattern-pipeline-temp")

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	log.Println("start generate file ...")
	start := time.Now()

	generateFiles()

	duration := time.Since(start)

	log.Println("Generate file finished in", duration.Seconds(), "seconds")
}

func randomString(length int) string {
	letters := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, length)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}

	return string(b)
}

func generateFiles() {
	os.RemoveAll(tempPath)
	os.MkdirAll(tempPath, os.ModePerm)

	for i := 0; i < TOTAL_FILE; i++ {
		filename := filepath.Join(tempPath, fmt.Sprintf("file-%d.txt", i))
		content := randomString(CONTENT_LENGTH)

		err := ioutil.WriteFile(filename, []byte(content), os.ModePerm)
		if err != nil {
			log.Println("Error when writing filename", filename)
		}

		if i%100 == 0 && i > 0 {
			log.Println(i, "files created")
		}
	}

	log.Printf("%d of total files created", TOTAL_FILE)
}
