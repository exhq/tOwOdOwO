package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		print("bruh wtf you aint got the args mann wtf\n")
		os.Exit(0)
	}
	args := os.Args[1:]
	if args[0] == "-a" {
		if _, err := os.Stat(gethome() + "/data"); err != nil {
			f, _ := os.Create(gethome() + "/data")

			_, _ = f.WriteString(args[1] + "\n")

			print("done")
			f.Close()
		} else {
			f, err := os.OpenFile(gethome()+"/data", os.O_APPEND|os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println(err)
				return
			}
			_, err = fmt.Fprintln(f, args[1])
			if err != nil {
				fmt.Println(err)
				f.Close()
				return
			}
			err = f.Close()
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println("file appended successfully")
		}
	} else if args[0] == "-r" {
		content, _ := ioutil.ReadFile(gethome() + "/data")

		for i, b := range strings.Split(string(content), "\n") {
			if b != "" {
				print(i+1, " - ", b, "\n")
			}
		}
	} else if args[0] == "-d" {
		content, _ := ioutil.ReadFile(gethome() + "/data")
		bruh := string(content)
		f, _ := os.Create(gethome() + "/data")
		for i, b := range strings.Split(bruh, "\n") {
			if b != "" {
				e, _ := strconv.Atoi(args[1])
				if i+1 != e {
					_, _ = f.WriteString(b + "\n")
				}
			}
		}
	}
}
func gethome() string {
	return (os.Getenv("HOME"))
}
