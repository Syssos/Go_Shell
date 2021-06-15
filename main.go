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

const linuxpath = "/usr/bin/"
var linuxcmds   = []string{"touch", "pwd", "ls"}
var codycmds    = []string{"site", "pond", "r2h"}
var logger      = loggerFromFile()

func main() {
    
    if len(os.Args) > 1 {
        // None interactive mode
        nonInteractiveShell()
    } else {
        // Interactive mode
        logger.Greet()
        interactiveShell()
        logger.Salute()
    }
}

func runLinux(cmd string, args string) {
    if args != "" {
        output, err := exec.Command(linuxpath+cmd, args).Output()
        check(err)
        fmt.Println(color.Green + string(output) + color.Reset)
    } else {
        output, err := exec.Command(linuxpath+cmd).Output()
        check(err)
        fmt.Println(color.Green + string(output) + color.Reset)
    }
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

// creates parsed slice from string based off spaces
func createCmdSlice(cmd string) []string {

    commands := []string{}

    words := strings.Fields(cmd)
    for _, word := range words {
        commands = append(commands, word)
    }

    return commands
}

func interactiveShell() {
    
    for ;; {
        args := ""
        handled := false
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
            } else {
                if len(parsed_input) > 1 {
                    args = strings.Join(parsed_input[1:], " ")
                }
                for x, cmd := range codycmds {
                    if parsed_input[0] == cmd {
                        runCody(codycmds[x], args)
                        handled = true
                    }
                }

                for x, cmd := range linuxcmds {
                    if parsed_input[0] == cmd {
                        runLinux(linuxcmds[x], args)
                        handled = true
                    }
                }

                if handled == true {
                    handled = false
                } else {
                    logger.Errormsg = errors.New("No valid command found")
                    logger.Err()
                }
            }   
        } else {
                // Blank input was passed, restarting loop to collect input
        }
    }
}

func nonInteractiveShell() {
    args := ""
    handled := false

    if len(os.Args) > 2 {
        args = strings.Join(os.Args[2:], " ")
    }

    for x, cmd := range codycmds {
        if os.Args[1] == cmd {
            runCody(codycmds[x], args)
            logger.Logmsg(fmt.Sprintf("%v, command ran from non-interactive mode", strings.Join(os.Args[1:], " ")))
            handled = true
        }
    }

    for x, cmd := range linuxcmds {
        if os.Args[1] == cmd {
            runLinux(linuxcmds[x], args)
            logger.Logmsg(fmt.Sprintf("%v, command ran from non-interactive mode", strings.Join(os.Args[1:], " ")))
            handled = true
        }
    }

    if handled == true {
        handled = false
    } else {
        logger.Errormsg = errors.New("No valid command found")
        logger.Err()
    }
}

func loggerFromFile() filelog.Flog {

    cwd, cwdErr := os.UserHomeDir()
    if cwdErr != nil {
        fmt.Println(cwdErr)
    }
    file, openErr := ioutil.ReadFile(cwd + "/.gofsh/config/LogSettings.toml")
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