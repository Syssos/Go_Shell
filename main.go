package main

import (
    "os"
    "fmt"
    "path"
    "time"
    "bufio"
    "errors"
    "strings"
    "os/exec"

    "github.com/Syssos/gofsh/src/color"
    "github.com/Syssos/gofsh/src/filelog"
)

const linuxpath = "/usr/bin/"
var linuxcmds   = []string{"touch", "cd", "pwd", "ls"}
var codycmds    = []string{"site", "pond", "r2h"}
var logger      = filelog.Flog{"User started Shell Interactive", "User ended Shell Interactive", "/tmp/GoShellLogfile.txt", "01-02-2006 15:04:05", time.FixedZone("UTC -7", -7*60*60), nil}

func main() {
    // output, err := exec.Command("/home/cody/go/bin/commands/site", "https://syssos.app").Output()
    
    if len(os.Args) > 1 {
        // None interactive mode
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

func check(err error) {
    if err != nil {
        logger.Errormsg = err
        logger.Err()
        fmt.Println(err.Error())
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

