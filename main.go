/* This is a rewrite of rxfetch in golang */ 
package main 

import (
  "fmt"
  "strconv"
  "strings"
  "os/exec"
  "github.com/go-ini/ini"
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

func getDistroName(configfile string) map[string]string {
  cfg, err := ini.Load(configfile)
  if err != nil {
    fmt.Printf("Fail to read file: ", err)
  }

  ConfigParams := make(map[string]string)
  ConfigParams["PRETTY_NAME"] = cfg.Section("").Key("PRETTY_NAME").String()

  return ConfigParams
}

func runCommand(cmd string, args string){
  run := exec.Command(cmd, args)
  out, err := run.CombinedOutput()

  if err != nil {
    fmt.Println(fmt.Sprint(err) + ": " + string(out))
    return
  }
  fmt.Println(string(out[:]))
}

func getPackages(){
  package_managers := [...]string{"pacman", "emerge", "apt", "xbps-install", "apk", "port", "nix", "dnf", "rpm", "pkg"}

  for _, pm := range package_managers {
    out, err := exec.Command("which", pm).Output()

    // gpt code here
    if err != nil {
			if exitError, ok := err.(*exec.ExitError); ok {
				if exitError.ExitCode() == 1 {
					continue 
				}
			}
			fmt.Printf("Error checking for %s: %s\n", pm, err)
			continue
		}

    if strings.TrimSpace(string(out)) == "/usr/bin/"+pm {
      switch pm {
        case "pacman":
          runCommand("pacman", "-Q")
        default:
          fmt.Println("gobrrr")
      }
    }
  }
}

func main() {
  runCommand("uname", "-a")
  OSInfo := getDistroName("/etc/os-release")
  OSRelease := OSInfo["PRETTY_NAME"]
  fmt.Print(OSRelease, "\n")

  getPackages()
}
