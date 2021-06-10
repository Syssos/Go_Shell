# Go Shell commands

<p align="center">
  <img src="https://github.com/Syssos/Go_Shell/blob/main/etc/images/Shell.PNG" alt="Go Shell img"/>
</p>

# Table of Contents
- [Core Commands](#core-commands)
	* [help](#help)
	* [exit](#exit)
	* [ls](#ls)
	* [pwd](#pwd)
	* [cd](#cd)
- [site](#site)
- [hostme](#hostme)
- [pond](#pond)
	* [Pond Example](#pond-example)

# Core Commands

Each command structure has a Usage() method which is called when the "help" command is ran. To use this feature, enter ```help``` followed by the command you need help with.

## Help
The help command is designed to work with every command the loop is able to run.

```
$ help ls
```

## Exit
This exits the shell properly and allows for logging of the user quitting.

```
$ exit
```
## ls
This command should work just about the same as it does on a native linux system, minus flag functionality. While it takes a location to list it cannot accept other flags at the moment.

```
$ ls
$ ls ../
$ ls Go_Shell/
```
## pwd
This command will print the working directory. It's a pretty straight forward command and doesn't take arguments.

```
$ pwd
```
## cd
This command changes the working directory much like cd in linux, like ls this will take a location to change to but will not accept flags.

```
$ cd ../
$ cd Go_Shell/
```

# site

The site command is a command that allows for the user to see a status code for a specific url. 

The purpose of this is mainly as a check to ensure the web application or program we are working with is returning an "OK" status.

```
$ site https://github.com/Syssos/Go_Shell
```

After a url is entered the site command will "remember" the url, if at any point in that shell instance you want the status of that site, the word status can be used to indicate it.

```
$ site status
```
the output for both of these commands should be

```
200 - Site https://github.com/Syssos/Go_Shell is active
```

# hostme

The host me command will spin up a small http server to temporarily host a single html file. 

To use this command, enter the command name hostme, then the name of the file you would like to host

```
$ hostme <path/to/htmlfile.html>
```

This will start a server on "localhost" over port 3000.

To view the host html file, navigate to your browser of choice and go to the following url
```
localhost:3000
```
or
```
127.0.0.1:3000
```
This should bring us to the host html file

# pond

The pond command is a simple web application environment creation tool. Meaning it will generate all of the basic files and folders needed to start working on a web application written in go.

Using the "pond create" command will create a "pond" and generate all the files and folders we'll need for a basic web application.

<p align="center">
  <img src="https://github.com/Syssos/Go_Shell/blob/main/settings/images/PondCreate.PNG" alt="Go Shell img" width="675px" height="600"/>
</p>

The pond will be created with the pond name in mind, the templates will be generated to use the name of the pond

Once the pond is created we can use CTRL+C to exit the Go Shell and cd into the newly created pond.
**Note ponds are meant to be test environments used outside of the Go Shell**

Once inside the new pond's directory run the `server.go` file with Go.

```bash
go run server.go
```
The server should start right up and tell us it is now listening on port 3000

```
Starting server on localhost:3000
```

When you see this message you should now be able to open a browser of your choice, and navigate to localhost port 3000. When the page loads, the following should be seen.

```
http://localhost:3000
```
or
```
http://127.0.0.1:3000
```

## Pond Example

Ponds are meant to give a developer the means to start up a basic web application so they can get to testing services as fast as possible.

An example of how this is achieved was created and stored within the [etc/Revero](https://github.com/Syssos/Go_Shell/tree/main/etc/Revero) directory.
