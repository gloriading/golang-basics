# Golang Learning

## Official website starters - Part 2

> https://golang.org/doc/tutorial/create-module

* `go mod init example.com/greetings`
  ```
  The go mod init command creates a go.mod file to track your code's dependencies. So far, the file includes only the name of your module and the Go version your code supports. But as you add dependencies, the go.mod file will list the versions your code depends on. This keeps builds reproducible and gives you direct control over which module versions to use.
  ```

* `touch greetings.go`

* Function:

  ```
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
## Official website starters - Part 2

> https://golang.org/doc/tutorial/handle-errors

* impoert `errors` package

  ```
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

  ```
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