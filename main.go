package main

import (
    "os"
    "fmt"
    "path"
    "time"
    "bytes"
    "bufio"
    "errors"
    "strings"
    "os/exec"
    "io/ioutil"
    "encoding/json"

    "github.com/komkom/toml"

    "github.com/Syssos/gofsh/src/color"
    "github.com/Syssos/gofsh/src/filelog"
)

var codycmds    = []string{"site", "pond", "r2h", "cd"}
var logger      = loggerFromFile()

func main() {
    
    if len(os.Args) > 1 {
        // None interactive mode
        check(nonInteractiveShell())
    } else {
        // Interactive mode
        logger.Greet()
        interactiveShell()
        logger.Salute()
    }
}

func runLinux(command string, args string) error {
    paths := getPathSlice()

    if command == "" {
        return errors.New("Missing command argument")
    }

    for _, path := range paths {
        if _, err := os.Stat(path+"/"+command); !os.IsNotExist(err) {
            if args != "" {
                output, err := exec.Command(path+"/"+command, args).Output()
                check(err)
                fmt.Println(color.Green + string(output) + color.Reset)
                return nil
            } else {
                output, err := exec.Command(path+"/"+command).Output()
                check(err)
                fmt.Println(color.Green + string(output) + color.Reset)
                return nil
            }
        }
    }

    return errors.New("Command not found")
}

func runCody(cmd string, args string) {

    if args != "" {
        output, err := exec.Command(filelog.GetHomeDir() + "/go/bin/" + cmd, args).Output()
        check(err)
        fmt.Println(string(output))
    } else {
        output, err := exec.Command(filelog.GetHomeDir() + "/go/bin/" + cmd).Output()
        check(err)
        fmt.Println(string(output))
    }
}

func interactiveShell() {
    
    for ;; {
        args := ""
        cwd   := filelog.GetCurrentDir()
        cuser := filelog.GetUser()
        
        // Printing what user see's before their input on command line
        fmt.Printf("%v - %v: ", color.Cyan + cuser + color.Reset, color.Blue + path.Base(cwd) + color.Reset)
        input := bufio.NewReader(os.Stdin)
        in, _ := input.ReadString('\n')
    
        parsed_input := createCmdSlice(in)
    
        if len(parsed_input) > 0 {
            if parsed_input[0] == "exit" {
        
                break
            } else if parsed_input[0] == "cd" {
                if len(parsed_input) > 1 {
                    check(Gofshcd(strings.Join(parsed_input[1:], " ")))
                } else {
                    fmt.Println("no directory to change to")
                }
            }else {
                if len(parsed_input) > 1 {
                    args = strings.Join(parsed_input[1:], " ")
                }
                for x, cmd := range codycmds {
                    if parsed_input[0] == cmd {
                        runCody(codycmds[x], args)
                        break
                    }
                }

                lerr := runLinux(parsed_input[0], args)
                check(lerr)
            }   
        } else {
                // Blank input was passed, restarting loop to collect input
        }
    }
}

func nonInteractiveShell() error{
    args := ""

    if len(os.Args) > 2 {
        args = strings.Join(os.Args[2:], " ")
    }

    for x, cmd := range codycmds {
        if os.Args[1] == cmd {
            runCody(codycmds[x], args)
            logger.Logmsg(fmt.Sprintf("%v, command ran from non-interactive mode", strings.Join(os.Args[1:], " ")))
            return nil
        }
    }

    lerr := runLinux(os.Args[1], args)
    check(lerr)
    logger.Logmsg(fmt.Sprintf("%v, command ran from non-interactive mode", strings.Join(os.Args[1:], " ")))

    return nil
}

func loggerFromFile() filelog.Flog {
    
    file, openErr := ioutil.ReadFile("/etc/gofsh/config/LogSettings.toml")
    if openErr != nil {
        fmt.Println(openErr)
    }

    doc := string(file)
    // Decodes toml to *json.Decoder
    dec := json.NewDecoder(toml.New(bytes.NewBufferString(doc)))
    
    st  := struct {
      Logger struct {
        Greeting string `json: "Greeting"`
        Salute string `json: "Salute"`
        LogFile string `json: "LogFile"`
        DtFormat string `json: "DtFormat"`
        DtTimeZone string `json: "DtTimeZone"`
        DtOffset int `json: "DtOffset"`
        Errormsg bool `json: "Errormsg"`
      } `json: "Logger"`
    }{}

    err := dec.Decode(&st)
    if err != nil {
      panic(err)
    }

    // Setting the error logs timestamp, timezone
    location := time.FixedZone(st.Logger.DtTimeZone, st.Logger.DtOffset*60*60)
    
    flog := filelog.Flog{ st.Logger.Greeting, st.Logger.Salute, st.Logger.LogFile, st.Logger.DtFormat, location, nil}
    return flog
}

func check(err error) {
    if err != nil {
        logger.Errormsg = err
        logger.Err()
        fmt.Println(err.Error())
    }
}

func Gofshcd(path string) error {
    if path != "" {
        os.Chdir(path)
        return nil
    } else {
        return errors.New("No directory to change to")
    }
}

// creates parsed slice from string based off spaces
func createCmdSlice(cmd string) []string {

    commands := []string{}

    words := strings.Fields(cmd)
    for _, word := range words {
        commands = append(commands, word)
    }

    return commands
}

func getPathSlice() []string {
    pathenvvar := os.Getenv("PATH")
    paths := []string{}
    path := []string{}

    for _, char := range []rune(pathenvvar) {
        if char == 58 {
            paths = append(paths, strings.Join(path, ""))
            path = []string{}
        } else {
            path = append(path, string(char))
        }
    }

    return(paths)
}
