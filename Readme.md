

# LIBRARY

##	"bufio"
```bufio.NewReader(os.Stdin) ```

##	"fmt"
```fmt.Print("Enter a grade: ")```
```fmt.Println("You have", 10-guesses, "guesses left.")```
##	"os"
```file, err := os.Open("myfile.txt")```
## type

## strconv
```bool, err := strconv.ParseBool("true")```
```grade, err := strconv.ParseFloat(input, 64)```
```guess, err := strconv.Atoi(input)```
## strings
```string.TrimSpace()```

## http
```response, err := http.Get("http://golang.org")```

## log
```log.Fatal(err)```

## 	"math/rand"
```target := rand.Intn(100) + 1```

## 	"time"
```seconds := time.Now().Unix()```
