package main

import (
	"os"
	"fmt"
	"strings"
	"io/ioutil"
)

func main() {
	if len(os.Args) > 1 {
		pond := Pond{Name: os.Args[1]}
		pond.CreateFolders()
		pond.CreateIndexHTML()
		pond.CreateFormHTML()
		pond.Create404HTML()
		pond.CreateStylesCSS()
		pond.MovePondIcon()
		pond.MoveGoServer()

	} else {
		fmt.Println("Invalid number of arguments")
	}

}

// Structure for pond, stores all data needed to keep track of file locations
type Pond struct {
	Name        string
	Local       string
	TempLocal   string
	StaticLocal string
	StylesLocal string
	ImagesLocal string
	JsLocal     string
}

// Creates a folder for all the files needed by the webapp
func (pond *Pond) CreateFolders() {

	cwd, err := os.Getwd()
	checkErr(err)

	var allPaths []string

	ppath := "/" + pond.Name
	pond.Local = cwd + ppath
	allPaths = append(allPaths, pond.Local)

	pond.TempLocal = pond.Local + "/templates"
	allPaths = append(allPaths, pond.TempLocal)

	pond.StaticLocal = pond.Local + "/static"
	allPaths = append(allPaths, pond.StaticLocal)

	pond.StylesLocal = pond.StaticLocal + "/styles"
	allPaths = append(allPaths, pond.StylesLocal)

	pond.ImagesLocal = pond.StaticLocal + "/images"
	allPaths = append(allPaths, pond.ImagesLocal)
	
	pond.JsLocal = pond.StaticLocal + "/js"
	allPaths = append(allPaths, pond.JsLocal)

	// Creating folder for each path above
	for _, pth := range allPaths {
		if _, err := os.Stat(pth); os.IsNotExist(err) {
			
			fmt.Println("Creating: ", pth)
			errDir := os.MkdirAll(pth, 0755)
			if errDir != nil {
				fmt.Println("Couldn't create directory -", pth, "mkdir error")
			}
			fmt.Println("Created local pond folder -", pth)
		} else {
			fmt.Println("Couldn't create Pond folder -", pth, "already exsists")
		}
	}
}

// Creates index template for app
func (pond *Pond) CreateIndexHTML() {
	
	if _, exsistsErr := os.Stat(pond.TempLocal + "/index.html"); os.IsNotExist(exsistsErr) {
		
		fmt.Println("\nCreating index.html file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/.gofsh/pond_templates/templates/index.html")
	    checkErr(err)
	    
	    // Replaces values in template while coping to reflect current pond
	    pageData := strings.Replace(string(dat), "{{title}}", pond.Name, -1)
	    
		fmt.Println("Saving index.html file content")
	    indErr := ioutil.WriteFile(pond.TempLocal + "/index.html", []byte(pageData), 0644)
	    checkErr(indErr)
		fmt.Println("index.html file created")
	} else {
		fmt.Println("index.html already exsists")
	}
}

// Generates 404 teplate for error page on app
func (pond *Pond) Create404HTML() {
	if _, exsistsErr := os.Stat(pond.TempLocal + "/404.html"); os.IsNotExist(exsistsErr) {
		
		fmt.Println("\nCreating 404.html file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/.gofsh/pond_templates/templates/404.html")
	    checkErr(err)
	    
	    pageData := strings.Replace(string(dat), "{{title}}", pond.Name, -1)
	    
		fmt.Println("Saving 404.html file content")
	    indErr := ioutil.WriteFile(pond.TempLocal + "/404.html", []byte(pageData), 0644)
	    checkErr(indErr)
		fmt.Println("404.html file created")
	} else {
		fmt.Println("404.html already exsists")
	}
}

// Generate form template page for app
func (pond *Pond) CreateFormHTML() {
	if _, exsistsErr := os.Stat(pond.TempLocal + "/form.html"); os.IsNotExist(exsistsErr) {
		
		fmt.Println("\nCreating form.html file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/.gofsh/pond_templates/templates/form.html")
	    checkErr(err)
	    
	    pageData := strings.Replace(string(dat), "{{title}}", pond.Name, -1)
	    
		fmt.Println("Saving form.html file content")
	    indErr := ioutil.WriteFile(pond.TempLocal + "/form.html", []byte(pageData), 0644)
	    checkErr(indErr)
		fmt.Println("form.html file created")
	} else {
		fmt.Println("form.html already exsists")
	}
}

// Copies CSS stylesheet to static/styles
func (pond *Pond) CreateStylesCSS() {
	if _, exsistsErr := os.Stat(pond.StylesLocal + "/main.css"); os.IsNotExist(exsistsErr) {
		
		fmt.Println("\nCreating main.css file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/.gofsh/pond_templates/static/styles/styles.css")
	    checkErr(err)

		fmt.Println("Saving main.css file content")
	    indErr := ioutil.WriteFile(pond.StylesLocal + "/main.css", []byte(dat), 0644)
	    checkErr(indErr)
		fmt.Println("main.css file created")
	} else {
		fmt.Println("main.css already exsists")
	}
}

// Copies image used in pond webapp to static/images folder
func (pond *Pond) MovePondIcon() {
	if _, exsistsErr := os.Stat(pond.ImagesLocal + "/Pond_Icon.ico"); os.IsNotExist(exsistsErr) {
		
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/.gofsh/pond_templates/static/images/Pond_Icon.ico")
	    checkErr(err)

	    err = ioutil.WriteFile(pond.ImagesLocal + "/Pond_Icon.ico", dat, 0644)
	    checkErr(err)
	    fmt.Println("\nPond_Icon.ico copied successfully from", cwd + "/.gofsh/pond_templates/Pond_Icon.ico")
	} else {
		fmt.Println("Pond_Icon.ico already exsists")
	}
}

// Copies go web server file to ponds "Local" directory
func (pond *Pond) MoveGoServer() {
	if _, exsistsErr := os.Stat(pond.Local + "/server.go"); os.IsNotExist(exsistsErr) {
	
		fmt.Println("\nCreating server.go file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/.gofsh/pond_templates/server.go")
	    checkErr(err)
	    
	    // Use area for replacing variables in go file
	    // pageData = strings.Replace(pageData, "{{Images_Folder}}", pond.ImagesLocal, -1)
	    
		fmt.Println("Saving index.html file content")
	    indErr := ioutil.WriteFile(pond.Local + "/server.go", []byte(string(dat)), 0644)
	    checkErr(indErr)
		fmt.Println("server.go file created")
	} else {
		fmt.Println("server.go already exsists")
	}
}

// Checks for errors in unction returns
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}