

# LIBRARY

##	"bufio"
```go
bufio.NewReader(os.Stdin) 
```

##	"fmt"
```go
fmt.Print("Enter a grade: ")
fmt.Println("You have", 10-guesses, "guesses left.")
fmt.Printf("About one-third: %0.2f\n", 1.0/3.0)
{
    resultString := fmt.Sprintf("About one-third: %0.2f\n", 1.0/3.0) 
    fmt.Printf(resultString)
}
```
|   Параметр  |  Описание   |
|-----|-----|
|%f   |Число с плавающей точкой|
|%d   |Десятичное целое число|
|%s   |Строка|
|%v   |Произвольное значение (подходящий формат выбирается на основании типа передаваемого значения)|
|%#v  |Произвольное значение, отформатированное в том виде, в котором оно отображается в коде Go|
|%T   |Тип переданного значения (int, string и т. п.)|
|%t   |Логическое значение (true или false)|
|%%   |Знак процента (литерал)|

##	"os"
```go
file, err := os.Open("myfile.txt")
```
## type

## strconv
```go
bool, err := strconv.ParseBool("true")
grade, err := strconv.ParseFloat(input, 64)
guess, err := strconv.Atoi(input)
```
## strings
```go
string.TrimSpace(input)
```

## http
```go
response, err := http.Get("http://golang.org")
```

## log
```go
log.Fatal(err)
```

## 	"math/rand"
```go
target := rand.Intn(100) + 1
```

## 	"time"
```go
seconds := time.Now().Unix()
```
