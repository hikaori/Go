# Basic of GO

## How to create go project

1. create "main.go" file
2. create a new module. to do so run below command.

   ```
   $go mod init <projectName>
   e.g) go mod init go-basic
   ```

   after run above command, automatically "go.mod" file should be created.

3. write a code.
   Write below code inside main.go file.

   ```
    package main

    import "fmt"

    func main(){
        fmt.Println("Hello World")
    }
   ```

4. Run file.
   ```
   $go run <fileName>
   e.g) go run main.go
   ```
   in terminal, "Hello World" should printed out.
