package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type envVars string
type excludedEnvVars []string

type fields struct {
	fileName         string
	excludeVariables excludedEnvVars
}

// sample call
// go run main.go -f=local.env -x=GOPATH,GOBIN,PATH,LS_COLORS
func main() {

	arguments := getCommandLineArguments(os.Args[1:])
	envVariables := getEnvsInVsCodeFormat(arguments.excludeVariables)

	envVariables.writeToFile(arguments.fileName)
}

func getCommandLineArguments(args []string) fields {
	var arguments fields

	for _, v := range args {
		//Sample v value : -o=env.test
		keyAndValue := strings.Split(v, "=")

		switch keyAndValue[0] {
		case "-f":
			arguments.fileName = keyAndValue[1]
		case "-x":
			arguments.excludeVariables = (excludedEnvVars)(strings.Split(keyAndValue[1], ","))
		case "--help":
			fmt.Println("Sample usage: go run main.go -f=testfile -x=ENV1,ENV2")
			fmt.Println("-f : Output file name")
			fmt.Println("-x : Comma separeted environment names which will be excluded")
		default:
			log.Fatal("Undefined argument :", keyAndValue[0])
		}
	}

	return arguments
}

//Sample return
//ENV1="VALUE1"
//ENV2="VALUE"
func getEnvsInVsCodeFormat(excludes excludedEnvVars) envVars {
	var variables envVars

	for _, e := range os.Environ() {
		pair := strings.SplitN(e, "=", 2)

		if !excludes.Contains(pair[0]) {
			variables += (envVars)(pair[0] + "=" + pair[1] + "\n")
		}
	}
	return variables
}

func (s envVars) writeToFile(fileName string) bool {
	dataInBytes := []byte(s)

	err := ioutil.WriteFile(fileName, dataInBytes, 0644)

	if err != nil {
		log.Fatal(err)
		return false
	}
	return true
}

func (envs excludedEnvVars) Contains(x string) bool {
	for _, n := range envs {
		if x == n {
			return true
		}
	}
	return false
}
