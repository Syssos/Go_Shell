# Filelog

Filelog is a tool designed to quickly and easily log an incident to a file. This was original created to log any errors as well as a start/end timestime.

## Usage

This package was designed to keep logging lightweight and easy to impliment multiple log files. 

To start off create an instance of the ``` Flog ``` struct, using the F_init(). This will generate a Flog instance with preset defaults. 

```go 
LogOne := filelog.F_init() 
```

If you would like to use your own settings use the following format

```go
LogOne := filelog.Flog{GreetingString, SaluteString, LogFilePath, nil}
```
**Note the last value is nil because it will hold an err if one is present**

Once the instance is created, to log a file, the Log method can be utilized.

### Logging strings (Timestamp not yet supported)

```go
LogOne.Log("Message to log")
```
### Logging errors (Logs with timestamp)

```go
LogOne.Errormsg = errors.New("Error to log")
``` 

### Logging errors from function return

```go
res, err := somefun(a, b)
if err != nil {
	LogOne.Errormsg = err
	LogOne.Err()
}
```
The code above will log the error message to the log file as well as print the error to the screen using a Println function.