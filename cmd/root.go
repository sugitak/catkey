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

var open_file = bufio.NewScanner(os.Stdin)
var rootCmd = &cobra.Command{
	Use:   "catkey",
	Short: "catkey reads ssl/tls keys",
	Long:  "catkey reads ssl/tls keys",
	Run: func(cmd *cobra.Command, args []string) {
		// case stdin
		if len(args) != 0 {
			fd, _ := os.Open(args[0])
			open_file = bufio.NewScanner(fd)
		}
		headline := first_line()

		switch {
		case headline == "-----BEGIN RSA PRIVATE KEY-----\n":
			text := headline + rest_lines()

			fmt.Println("$ openssl rsa -noout -text")
			command("openssl rsa -noout -text", text)
		case headline == "-----BEGIN EC PRIVATE KEY-----\n" || headline == "-----BEGIN EC PARAMETERS-----\n":
			text := headline + rest_lines()

			fmt.Println("$ openssl ec -noout -text")
			command("openssl ec -noout -text", text)
		case headline == "-----BEGIN CERTIFICATE REQUEST-----\n":
			text := headline + rest_lines()

			fmt.Println("$ openssl req -noout -text")
			command("openssl req -noout -text", text)
		case headline == "-----BEGIN CERTIFICATE-----\n":
			text := headline + rest_lines()

			fmt.Println("$ openssl x509 -noout -text")
			command("openssl req -noout -text", text)
		default:
			fmt.Println("Unknown Format! Maybe not implemented yet.")
		}
	},
}

func first_line() string {
	open_file.Scan()
	return open_file.Text() + "\n"
}

func rest_lines() string {
	text := ""
	for open_file.Scan() {
		text += open_file.Text() + "\n"
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
