package cmd

import (
	"bufio"
	"fmt"
	"github.com/spf13/cobra"
	"io"
	"os"
	"os/exec"
	"strings"
)

var rootCmd = &cobra.Command{
	Use:   "catkey",
	Short: "catkey reads ssl/tls keys",
	Long:  "catkey reads ssl/tls keys",
	Run: func(cmd *cobra.Command, args []string) {
		// case stdin
		headline := stdin_first()

		switch {
		case headline == "-----BEGIN RSA PRIVATE KEY-----\n":
			text := headline + stdin_rest()

			fmt.Println("$ openssl rsa -noout -text")
			command("openssl rsa -noout -text", text)
		default:
			fmt.Println("Unknown Format! Maybe not implemented yet.")
		}
	},
}

var stdin = bufio.NewScanner(os.Stdin)

func stdin_first() string {
	stdin.Scan()
	return stdin.Text() + "\n"
}

func stdin_rest() string {
	text := ""
	for stdin.Scan() {
		text += stdin.Text() + "\n"
	}
	return text
}

func command(commandline string, input string) {
	splitted := strings.Split(commandline, " ")
	command := splitted[0]
	opts := splitted[1:]
	cmd := exec.Command(command, opts...)

	cmd_stdin, _ := cmd.StdinPipe()
	io.WriteString(cmd_stdin, input)
	cmd_stdin.Close()

	out, _ := cmd.Output()
	fmt.Printf("Result: %s", out)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
