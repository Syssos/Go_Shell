# Color

This package has one intended purpose. Fast and easy access to multiple colors to use on the output of a program.

## Install
This package has the most use case outside of a shell enviornment. Making it a perfect package to utilize in other projects, while simple it gets the job done

To install this package use the go get command

```bash
go get github.com/Syssos/Go_Shell/color
```
This should download the package and install it in the `$GOPATH/src` directory.

## Usage
Once the package is retreived with go get we should be able to start using it in another package.

That starts with an import statment

```go
import "github.com/Syssos/Go_Shell/color"
```

Each "color" is of type string. To color text, append the text you would like to print to the color of your choice.

After the text and color are concated, we need to append the reset value to the end of the string. This will tell the compiler we are done using that color.

```go
str := color.Red + "This is a test" + color.Reset

fmt.Println(str)
```

Every color will require a reset after using, to automate this task you can create a "coloring" function for specific colors.

```go
func colorRed(str string) string {
	return color.Red + str + color.Reset
}

func colorBlue(str string) string {
	return color.Blue + str + color.Reset
}

str := "This " + colorRed("is a") + colorBlue(" test")

fmt.Println(str)
```
