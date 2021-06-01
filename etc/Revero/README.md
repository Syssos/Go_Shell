# Revero

<p align="center">
  <img src="https://github.com/Syssos/Go_Shell/blob/main/etc/images/PondExampleApp.PNG" alt="Go Shell img"/>
</p>

This Directory started as a product of the `pond create` command. The only real change Revero offers from a default pond, is the string entered gets reversed. This README will document how that change was introduced.

# Overview
To start things off I would like to mention that we will be using the pond created in the commands example. If you need more information you can check it out [here](https://github.com/Syssos/Go_Shell/tree/main/cmds) then come back and continue.

## A ponds files

Ponds that get generated contain all of the basics a developer will need to start testing ideas for microservices. 

They will contain css and html pages for those in need of a quick way to test styling concepts.

	- `index.html`: Homepage for webapp, generate with the pond name in mind.
	- `form.html` : Page used for collection user input, handles "POST", and "GET" methods, Go template dependent
	- `404.html`  : Custom 404 error handling page for webapp

	- `static/styles/main.css` : Styles page used on all 3 html pages

They will contain the Go code that servers the files and handles routes.
	
	- `server.go` : Handles routes and server html files, accounts for static routes
	
and even have images for the app to be noticable from other tabs

	- `static/images/Pond_Icon.ico` : Image for the pond

## Modifing Server.go
This example will outline modifing the application to allow us to complete our objective of reversing a string. To get started we need to create a function that will complete this task.

```go
func Reverse(s string) string {
	rev := []rune(s)
	for x, y := 0, len(s)-1; x < y; x, y = x+1, y-1 {
		rev[x], rev[y] = rev[y], rev[x]
	}

	return string(rev)
}
```

With this function created we need to impliment it into the server file. To do that I will add it to the end of the file.

Now we can reverse the users input with just one function call, so lets go modify what gets returned to the user.

<p align="center">
  <img src="https://github.com/Syssos/Go_Shell/blob/main/etc/images/UpdateingOutput.PNG" alt="Go Shell img"/>
</p>

The image above outlines all the code we change inside of the "POST" method.

While much more complex goals can be acheived. The task of reverse a string is pretty straight forward.

With the changes made. Re-running the command should grant no errors and start the server on port 3000. When using the /form you should now see that the data entered is revered.

<p align="center">
  <img src="https://github.com/Syssos/Go_Shell/blob/main/etc/images/AppResult.PNG" alt="Go Shell img"/>
</p>