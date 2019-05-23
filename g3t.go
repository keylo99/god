package main

import (
    "bufio"
    "fmt"
	"os"
	"os/exec"
	"bytes"
    "strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(execCmd("ls .", false))
    for {
        fmt.Print("> ")
        // Read the keyboad input.
        input, err := reader.ReadString('\n')
        if err != nil {
            fmt.Fprintln(os.Stderr, err)
		}
		// Handle the execution of the input.
        if retval := execCmd(input, true); len(retval) > 0 {
			fmt.Fprintln(os.Stderr, retval)
        }
    }
}

var outb, errb bytes.Buffer

func execCmd(input string, git_cmd bool) string {
	// Add the 'git' command
	if git_cmd {
		input = "git " + input
	}
    // Remove the newline character.
    input = strings.TrimSuffix(input, "\n")

    // Split the input separate the command and the arguments.
    args := strings.Split(input, " ")

    // Check for built-in commands.
    switch args[1] {
    case "exit":
        os.Exit(0)
    }

    // Prepare the command to execute.
    cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	if git_cmd{
    	cmd.Stderr = os.Stderr
		cmd.Stdout = os.Stdout
	}else{
		cmd.Stdout = &outb
		cmd.Stderr = &errb
	}
    // Execute the command and return the error.
	err := cmd.Run()
	if err != nil{
		return err.Error()
	}
	if !git_cmd{
		return outb.String()
	}
	return ""
}