# Go Shell

This is a simple shell built in go designed to give you basic functionality such as changing directory and listing files.

## Main

This file will call the Loop function from the cmds package. This loop is what is responsible for all of the commands. The plans to re-use this code mean I need all of the code in one package, more or less, that I can grab and use in another project.

## Logging

This program will log errors to a file called "logfile.txt". It will be placed within the working directory upon running the program. This is temporary as I figure out where the most convient place for this file to live on the system is.

## Travis

Very unhappy with this, .org worked the first time I used it no issues. the .com is the worst thing I've ever seen.