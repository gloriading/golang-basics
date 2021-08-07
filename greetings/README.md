# Golang Learning

## Official website starters - Part 2

> https://golang.org/doc/tutorial/create-module

* `go mod init example.com/greetings`
  ```
  The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.
  ```

* `touch greetings.go`

* Function:

  ```go
  func Hello(name string) string {
    message := fmt.Sprintf("Hi %v. Welcome!", name)
    return message
  }

  ```

    * In Go, a function whose name starts with a capital letter can be called by a function _not_ in the same package. This is known in Go as an `exported name`.

    * In Go, the `:=` operator is a shortcut for declaring and initializing a variable in one line (Go uses the value on the right to determine the variable's type).

* Go to another module (Hello) and call Greetings.Hello
    - For production use, you’d publish the example.com/greetings module from its repository (with a module path that reflected its published location), where Go tools could find it to download it. For now, because you haven't published the module yet, you need to adapt the example.com/hello module so it can find the example.com/greetings code on your local file system.

    - To do that, use the `go mod edit command` to edit the example.com/hello module to redirect Go tools from its module path (where the module isn't) to the local directory (where it is).

    ```go
      // in Hello module
      go mod tidy
      go mod edit -replace example.com/greetings=../greetings
    ```
## Official website starters - Part 3

> https://golang.org/doc/tutorial/handle-errors

* impoert `errors` package

  ```go
  import (
      "errors"
      "fmt"
  )

  // Any Go function can return multiple values.

  func Hello(name string) (string, error) {
      if name == "" {
          return "", errors.New("empty name")
      }

      message := fmt.Sprintf("Hi, %v. Welcome!", name)
      return message, nil
  }
  ```

* Handle error

  ```go
  package main

  import (
      "fmt"
      "log"

      "example.com/greetings"
  )

  func main() {
      // Set properties of the predefined Logger, including
      // the log entry prefix and a flag to disable printing
      // the time, source file, and line number.
      log.SetPrefix("greetings: ")
      log.SetFlags(0)

      message, err := greetings.Hello("")
      if err != nil {
          log.Fatal(err)
          // exit the program.
      }

      fmt.Println(message)
  }
  ```

## Official website starters - Part 4

> https://golang.org/doc/tutorial/random-greeting

* Import `math/rand` and `time` packages

* Add an init function to seed the rand package with the current time. Go executes init functions automatically at program startup, after global variables have been initialized.

  ```go
  func init() {
    rand.Seed(time.Now().UnixNano())
  }
  ```
* Create private function

  *  A `slice` is like an array, except that its size changes dynamically as you add and remove items

  * function starts with a lower-cased letter meaning it's a private function - not accesible to other modules

  ```go
  func randomFormat() string {
    formats := []string{
      "Good morning, %v!!",
      "Nice to meet you %v!!",
      "Hi %v, your hair looks great today!!",
    }
    return formats[rand.Intn(len(formats))]
  }

  ```

## Official website starters - Part 5

> https://golang.org/doc/tutorial/greetings-multiple-people


* Hellos returns a map that associates each of the named people with a greeting message.

  ```go
  func Hellos(names []string) (map[string] string, error) {
    // In Go, you initialize a map with the following syntax: make(map[key-type]value-type)
    messages := make(map[string]string)

    // In this for loop, range returns two values: the index of the current item in the loop and a copy of the item's value
    for _, name := range names {
      message, err := Hello(name)
      if err != nil {
        return nil, err
      }
      messages[name] = message
    }

    return messages, nil
  }
  ```
