# Golang Quickstart

## Contents

A quickstart for Go-lang with all the necessary code and comments

| Name                                                   | File                                                 |
| ------------------------------------------------------ | ---------------------------------------------------- |
| 1. Variables                                           | [variables_functions.go](src/variables_functions.go) |
| 2. Functions (Arguments, Signatures and Named returns) | [variables_functions.go](src/variables_functions.go) |
| 3. Simple Datatypes                                    | [types.go](src/types.go)                             |
| 4. Complex Datatypes                                   | [types.go](src/types.go)                             |
| 5. File Handling                                       | [write_file.go](src/write_file.go)                   |
| 6. Flow Control (Conditionals)                         | [control_loops.go](src/control_loops.go)             |
| 7. Loops                                               | [control_loops.go](src/control_loops.go)             |
| 8. HTTP Requests                                       | [http_req.go](src/http_req.go)                       |

---

## Execution

- Open [main.go](main.go) and remove the comment on the function required
- `go run main.go`

## Go Modules

- Go Modules allows us to restructure our code easily
- For initializing a new project - `go mod init <project-name>`
- For installing dependencies - `go get -u <dependency>`
- `package <folder-name>` inside the file should match the parent folder
- inside the root folder, make a file main.go with `package main` to run the functions
- Names of **Exported functions** have to start in uppercase.
