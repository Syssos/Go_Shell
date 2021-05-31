package cmds

import (
	"os"
	"fmt"
	"errors"
	"strings"
	"io/ioutil"

	"github.com/Syssos/Go_Shell/color"
)

// Structure used to call pond command
type Pond_cmd struct {
	Args []string
}

// Function called when user uses the "pond" command
func (p *Pond_cmd) Run() error{

	if len(p.Args) > 0 {
		if p.Args[0] == "create" {
			if len(p.Args) > 1 {	
				// Create new pond with name p.Args[1], ($ pond create p.Args[1])
				pond := Pond{Name: p.Args[1]}
				pond.CreateFolders()
				pond.CreateIndexHTML()
				pond.Create404HTML()
				pond.CreateStylesCSS()
				pond.MovePondIcon()
				pond.MoveGoServer()
				return nil
			} else {
				return errors.New("Cannot create pond without name, check usage with \"help pond\"")
			}
		}
	}

	return errors.New("pond method not found, check usage with \"help pond\"")
}

// Prints the usage message for the pond command
func (p Pond_cmd) Usage() {
	fmt.Println(color.Yellow + "\n\tpond - Creates a directory and basic files for a go web application\n\n\t\tpond create <pondName>\n" + color.Reset)
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
				fmt.Println(color.Red + "Couldn't create directory -", pth, "mkdir error" + color.Reset)
			}
			fmt.Println(color.Green + "Created local pond folder -", pth , "" + color.Reset)
		} else {
			fmt.Println(color.Red + "Couldn't create Pond folder -", pth, "already exsists" + color.Reset)
		}
	}
}

// Creates index template for app
func (pond *Pond) CreateIndexHTML() {
	
	if _, exsistsErr := os.Stat(pond.TempLocal + "/index.html"); os.IsNotExist(exsistsErr) {
		
		fmt.Println("\nCreating index.html file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/GofshTemplates/templates/index.html")
	    checkErr(err)
	    
	    // Replaces values in template while coping to reflect current pond
	    pageData := strings.Replace(string(dat), "{{title}}", pond.Name, -1)
	    
		fmt.Println("Saving index.html file content")
	    indErr := ioutil.WriteFile(pond.TempLocal + "/index.html", []byte(pageData), 0644)
	    checkErr(indErr)
		fmt.Println(color.Green + "index.html file created" + color.Reset)
	} else {
		fmt.Println(color.Red + "index.html already exsists" + color.Reset)
	}
}

// Generates 404 teplate for error page on app
func (pond *Pond) Create404HTML() {
	if _, exsistsErr := os.Stat(pond.TempLocal + "/404.html"); os.IsNotExist(exsistsErr) {
		
		fmt.Println("\nCreating 404.html file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/GofshTemplates/templates/404.html")
	    checkErr(err)
	    
	    pageData := strings.Replace(string(dat), "{{title}}", pond.Name, -1)
	    
		fmt.Println("Saving 404.html file content")
	    indErr := ioutil.WriteFile(pond.TempLocal + "/404.html", []byte(pageData), 0644)
	    checkErr(indErr)
		fmt.Println(color.Green + "404.html file created" + color.Reset)
	} else {
		fmt.Println(color.Red + "404.html already exsists" + color.Reset)
	}
}

// Copies CSS stylesheet to static/styles
func (pond *Pond) CreateStylesCSS() {
	if _, exsistsErr := os.Stat(pond.StylesLocal + "/main.css"); os.IsNotExist(exsistsErr) {
		
		fmt.Println("\nCreating main.css file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/GofshTemplates/static/styles/styles.css")
	    checkErr(err)

		fmt.Println("Saving main.css file content")
	    indErr := ioutil.WriteFile(pond.StylesLocal + "/main.css", []byte(dat), 0644)
	    checkErr(indErr)
		fmt.Println(color.Green + "main.css file created" + color.Reset)
	} else {
		fmt.Println(color.Red + "main.css already exsists" + color.Reset)
	}
}

// Copies image used in pond webapp to static/images folder
func (pond *Pond) MovePondIcon() {
	if _, exsistsErr := os.Stat(pond.ImagesLocal + "/Pond_Icon.ico"); os.IsNotExist(exsistsErr) {
		
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/GofshTemplates/static/images/Pond_Icon.ico")
	    checkErr(err)

	    err = ioutil.WriteFile(pond.ImagesLocal + "/Pond_Icon.ico", dat, 0644)
	    checkErr(err)
	    fmt.Println(color.Green + "\nPond_Icon.ico copied successfully from", cwd + "/GofshTemplates/Pond_Icon.ico" + color.Reset)
	} else {
		fmt.Println(color.Red + "Pond_Icon.ico already exsists" + color.Reset)
	}
}

// Copies go web server file to ponds "Local" directory
func (pond *Pond) MoveGoServer() {
	if _, exsistsErr := os.Stat(pond.Local + "/server.go"); os.IsNotExist(exsistsErr) {
	
		fmt.Println("\nCreating server.go file content")
		cwd, cwdErr := os.UserHomeDir()
		checkErr(cwdErr)
		dat, err := ioutil.ReadFile(cwd + "/GofshTemplates/server.go")
	    checkErr(err)
	    
	    // Use area for replacing variables in go file
	    // pageData = strings.Replace(pageData, "{{Images_Folder}}", pond.ImagesLocal, -1)
	    
		fmt.Println("Saving index.html file content")
	    indErr := ioutil.WriteFile(pond.Local + "/server.go", []byte(string(dat)), 0644)
	    checkErr(indErr)
		fmt.Println(color.Green + "server.go file created" + color.Reset)
	} else {
		fmt.Println(color.Red + "server.go already exsists" + color.Reset)
	}
}

// Checks for errors in unction returns
func checkErr(err error) {
	if err != nil {
		fmt.Println(color.Red + "", err, "" + color.Reset)
	}
}