package main

import (
	"bufio"
	"fmt"
	"github.com/captncraig/gpiet/machine"
	"github.com/captncraig/gpiet/micropiet"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatalln(err)
	}
	for _, f := range files {
		if !strings.HasSuffix(f.Name(), ".mp") {
			continue
		}
		runTest(f.Name())
	}
	fmt.Println("--------------")
}

func runTest(name string) {
	fmt.Println("--------------")
	fmt.Print(name, "...  ")

	vm := machine.NewMachine()

	file, err := os.Open(name)
	if err != nil {
		log.Fatalln(err)
	}
	defer func() {
		x := recover()
		if x != nil {
			log.Fatalln("TEST FAILED: ", x)
		}
	}()
	scanner := bufio.NewScanner(file)
	text := ""
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "@@") {
			prog, err := micropiet.Parse(text)
			if err != nil {
				log.Fatalln(err)
			}
			prog.Run(vm)
			compareStacks(vm.GetStack(), line[3:])
			text = ""
		} else {
			text = text + line + "\n"
		}
	}
	fmt.Println("PASS")
}

func compareStacks(stack []int64, expected string) {
	actual := fmt.Sprint(stack)
	actual = strings.Trim(actual, "[] ")
	expected = strings.TrimSpace(expected)
	if actual != expected {
		fmt.Println("\nStack doesn't match!!!")
		fmt.Println("Expected: ", expected)
		fmt.Println("Actual: ", actual)
		log.Fatalln("TEST FAILED")
	}
}
