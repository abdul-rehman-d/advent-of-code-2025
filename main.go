package main

import (
	"advent-of-code-2025/day1"
	"advent-of-code-2025/day2"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"slices"
	"strconv"
	"strings"
	"time"
)

func getChallengeInput(i int) string {
	arrByte, err := os.ReadFile(fmt.Sprintf("./day%d/input.txt", i))
	if err != nil {
		panic(err)
	}
	return string(arrByte)
}

func getMainFileContents(day int, part string) string {
	return fmt.Sprintf(`package day%d

func Part%s(data string) int {
	return 0
}
`, day, part)
}

func getTestFileContents(day int) string {
	return fmt.Sprintf(`package day%d

import (
	"testing"
)

const (
	data = ""
)

func TestPartA(t *testing.T) {
	expected := 0
	result := PartA(data)

	if expected != result {
		t.Fatalf("\nExpected = %%d\nResult = %%d\n", expected, result)
	}
}


func TestPartB(t *testing.T) {
	expected := 0
	result := PartB(data)

	if expected != result {
		t.Fatalf("\nExpected = %%d\nResult = %%d\n", expected, result)
	}
}
`, day)
}

func solve(day int) {

	partA := []func(string) int{
		day1.PartA,
		day2.PartA,
	}
	partB := []func(string) int{
		day1.PartB,
		day2.PartB,
	}

	if day > len(partA) {
		log.Fatalf("Day cannot be more than %d\n", len(partA))
	}

	input := getChallengeInput(day)
	startTime := time.Now()
	partASolution := partA[day-1](input)
	partATimeTaken := time.Now().Sub(startTime).Milliseconds()
	startTime = time.Now()
	partBSolution := partB[day-1](input)
	partBTimeTaken := time.Now().Sub(startTime).Milliseconds()

	fmt.Printf("Part A Answer => %d\n", partASolution)
	fmt.Printf("Part B Answer => %d\n", partBSolution)
	fmt.Printf("Part A Time => %dms\n", partATimeTaken)
	fmt.Printf("Part B Time => %dms\n", partBTimeTaken)
}

func generate(day int) error {
	folderName := fmt.Sprintf("./day%d", day)

	if s, err := os.Stat(folderName); !os.IsNotExist(err) && s != nil {
		return fmt.Errorf("already exists")
	}

	err := os.Mkdir(folderName, 0755)
	if err != nil {
		return err
	}

	var f *os.File

	f, err = os.Create(fmt.Sprintf("%s/a.go", folderName))
	if err != nil {
		return err
	}
	f.WriteString(getMainFileContents(day, "A"))

	f, err = os.Create(fmt.Sprintf("%s/b.go", folderName))
	if err != nil {
		return err
	}
	f.WriteString(getMainFileContents(day, "B"))

	_, err = os.Create(fmt.Sprintf("%s/input.txt", folderName))
	if err != nil {
		return err
	}

	_, err = os.Create(fmt.Sprintf("%s/questionA.md", folderName))
	if err != nil {
		return err
	}
	_, err = os.Create(fmt.Sprintf("%s/questionB.md", folderName))
	if err != nil {
		return err
	}

	f, err = os.Create(fmt.Sprintf("%s/day%d_test.go", folderName, day))
	if err != nil {
		return err
	}
	f.WriteString(getTestFileContents(day))

	f, err = os.OpenFile("main.go", os.O_RDWR, 0644)
	if err != nil {
		return err
	}
	defer f.Close()
	bytes, err := io.ReadAll(f)
	if err != nil {
		return err
	}

	lines := strings.Split(string(bytes), "\n")

	foundAt := -1
	for lineNum, line := range lines {
		if strings.Contains(line, "import") {
			foundAt = lineNum
			break
		}
	}

	if foundAt < 0 {
		return fmt.Errorf("why not")
	}

	column := strings.LastIndex(lines[foundAt+1], "/day")

	if column < 0 {
		return fmt.Errorf("why not 2")
	}

	toInsert := []string{
		lines[foundAt+1][:column] + fmt.Sprintf("/day%d\"", day),
	}
	lines = slices.Concat(lines[:foundAt+2], toInsert, lines[foundAt+2:])

	for lineNum := foundAt; lineNum < len(lines); lineNum++ {
		if strings.Contains(lines[lineNum], "partA :=") {
			foundAt = lineNum
			break
		}
	}

	toInsert = []string{fmt.Sprintf("day%d.PartA,", day)}
	for lineNum := foundAt; lineNum < len(lines); lineNum++ {
		if strings.Contains(lines[lineNum], "}") {
			foundAt = lineNum
			lines = slices.Concat(lines[:lineNum], toInsert, lines[lineNum:])
			break
		}
	}

	toInsert = []string{fmt.Sprintf("day%d.PartB,", day)}
	for lineNum := foundAt; lineNum < len(lines); lineNum++ {
		if strings.Contains(lines[lineNum], "partB := ") {
			foundAt = lineNum
			break
		}
	}
	for lineNum := foundAt; lineNum < len(lines); lineNum++ {
		if strings.Contains(lines[lineNum], "}") {
			lines = slices.Concat(lines[:lineNum], toInsert, lines[lineNum:])
			break
		}
	}

	f.Seek(0, 0)
	f.WriteString(strings.Join(lines, "\n"))

	cmd := exec.Command("gopls", "format", "-w", "main.go")
	if err = cmd.Run(); err != nil {
		return err
	}

	return nil
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println(`Usage ./aoc [cmd] [day]
Available commands:-
- solve
- generate`)
		os.Exit(1)
	}

	day, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Printf("Error parsing day %s\n", os.Args[2])
		os.Exit(1)
	}

	if day <= 0 {
		fmt.Println("Day cannot be less than 1")
		os.Exit(1)
	}

	switch os.Args[1] {
	case "solve":
		solve(day)
	case "generate":
		err := generate(day)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	default:
		fmt.Println("Invalid command.")
		os.Exit(1)
	}
}
