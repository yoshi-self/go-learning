Basic directory structure

## run / build
* `go run main.go app.go`
  * Needs to specify both files
* `go build`
  * Automatically names the binary after parent directory
  * Build \*.go in current directory and also resolve packages

## test
* Filenames for testing must ends \_test.go
* Function names must be `TestXxx`
`go test -v`
`go test -v ./animals`

