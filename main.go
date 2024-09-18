/* This is a rewrite of rxfetch in golang */ 
package main 

import (
  "fmt"
  "strconv"
  "strings"
  "os/exec"
  "github.com/benhoyt/goawk/interp"
)

var (
	// Colors and font options via ANSI escape codes
	Reset     = "\033[0m"
	Black     = "\033[30m"
	Red       = "\033[31m"
	Green     = "\033[32m"
	Yellow    = "\033[33m"
	Blue      = "\033[34m"
	Magenta   = "\033[35m"
	Cyan      = "\033[36m"
	Gray      = "\033[37m"
	White     = "\033[97m"
	Bold      = "\033[1m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Invert    = "\033[7m"
)

func Color(input interface{}, color ...string) string {
	var s string
	c := ""
	for i := range color {
		c = c + color[i]
	}
	switch v := input.(type) {
	case int:
		s = c + strconv.Itoa(v) + Reset
	case bool:
		s = c + strconv.FormatBool(v) + Reset
	case []string:
		s = c + strings.Join(v, ", ") + Reset
	case string:
		s = c + v + Reset
	default:
		fmt.Printf("Unsupported type provided to Color func - %T\n", v)
	}
	return s
}

// This function will define the further use of commands as the kernel type will make changes in the type of OS. 
func getKernel(cmd string, args string){
  out, err := exec.Command(cmd).Output()
  if err != nil {
    fmt.Printf("%s", err)
  }
  fmt.Println(string(out[:]))
}

func getDistroName(){
  out:= exec.Command("lsb_release", "-d")

  err := interp.Exec("$0 { print $1 }", " ", out, nil)
  if err != nil {
    fmt.Println(err)
    return
  }

  fmt.Println(string(out[:]))
}

func getPackages(cmd string, args string){
  out, err := exec.Command(cmd, args).Output()

  package_managers := [...]string{"pacman", "emerge", "apt", "xbps-install", "apk", "port", "nix", "dnf", "rpm", "pkg"}

  if err != nil {
    fmt.Printf("%s", err)
  }

  fmt.Println(string(out[:]))
}

func main() {
  getDistroName()
}
