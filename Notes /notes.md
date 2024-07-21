Summary:

-  Golang is Compiled, can run directly without VM
-  run commmand - go run main.go
-  Get commands - go help
-  Types - Everything is a Type eg - String, Bool, Integer, Floating, Complex
-  Advance Types - Array, Slices, Maps, Structs, Pointers
-  %T - type of variable
-  If first letter is capital of a variable then it's a public variable
-  To search for any package in golang go to pkg.go.dev
-  ```
   //user input
   reader := bufio.NewReader(os.Stdin);
   ```

-  To build a executable file run - go build
-  Memory allocation and management happens automatically in golang
-  Garbage collections happens automatically
-  Pointer - use \* for creating a pointer and & for referencing/getting the address of the variable
-  removing a value from slice through index

   ```
   courses = append(courses[:index], courses[index+1:]...)

   ```

-
-  Steps
-  Create a folder
-  create main.go file
-  In cmd of folder run command `go mod init name`
-  run code - go run main.go


-install a dependency from terminal
 go get -u github.com/gorilla/mux

 #Mod Reference - https://go.dev/ref/mod

 -verify installed dependency from terminal
 go mod verify

 -list all dependencies in terminal
 go list -m all

 -tidy all library (expensive command)
 go mod tidy

 -check a dependency is used where 
 go mod why github.com/gorilla/mux

 -Create a file with dependencies for prod 
 go mod vendor

 -run code first from vendor and if dependency is not installed then from cache 
 go run -mod=vendor main.go  