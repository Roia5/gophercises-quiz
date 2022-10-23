package main

import (
    "bufio"
    "encoding/csv"
    "flag"
    "fmt"
    "log"
    "math/rand"
    "os"
    "strings"
    "time"
)

var totalCorrect int = 0
var total int = 0

func handleError(err error) {
    if err != nil {
        log.Fatal(err)
    }
}
func handleQuestion(record []string) int {
    fmt.Println(record[0])
    scanner := bufio.NewScanner(os.Stdin)
    fmt.Print("Enter your answer: ")
    scanner.Scan()
    text := scanner.Text()
    text = strings.TrimSpace(text)
    text = strings.ToLower(text)
    if text != record[1] {
        return 0
    }
    return 1
}
func finalMessage() {
    fmt.Printf("Number of questions: %v\n", total)
    fmt.Printf("Correct answers: %v\n", totalCorrect)
}
func main() {
    filenameFlag := flag.String("filename", "problems.csv", "csv file")
    timerFlag := flag.Int("timer", 30, "timeout in seconds")
    shuffleFlag := flag.Bool("shuffle", false, "should the questions be shuffled?")
    flag.Parse()
    if !strings.HasSuffix(*filenameFlag, ".csv") {
        fmt.Println("Please use a .csv file")
        return
    }
    f, err := os.Open(*filenameFlag)
    handleError(err)
    defer f.Close()
    fmt.Print("Press any key to start the timer...")
    scanner := bufio.NewScanner(os.Stdin)
    scanner.Scan()
    timer1 := time.NewTimer(time.Duration(*timerFlag) * time.Second)
    go func() {
        <-timer1.C
        fmt.Println("\n=========")
        fmt.Println("Timeout!")
        finalMessage()
        os.Exit(0)
    }()
    reader := csv.NewReader(f)
    questions, err := reader.ReadAll()
    handleError(err)
    if *shuffleFlag {
        rand.Seed(time.Now().UnixNano())
        rand.Shuffle(len(questions), func(i, j int) { questions[i], questions[j] = questions[j], questions[i]})
    }
    for _,question := range questions {
        total++
        totalCorrect += handleQuestion(question)
    }
    finalMessage()
}
