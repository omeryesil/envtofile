# envtofile

Writes environment variables to a file in KEY=VALUE format. 

Go extension for VS Code runs in a separete shell session, therefore, it does not get environment variables from VS Code terminal. A simple use case can be:

1. You have a shell script to set environment variables for dev/local/test/ etc.. You need to run this script before running your go app
2. You launch the debug mode from VSCode, however, it wont get the environment variables from the step 1 as VS Code extensions (in this case it is Go extension) run in a separete terminal instance

In order to resolve this, we in this code, we are getting all the environment variables from the current terminal session, and writing them to a file. In this way we can easily set the file location in VS Code's launch.json file, which will make the env variables available for the Debug mode. 

Note that I might be wrong, may be there is a better/easier way to set pull environment variables for the Debug mode.

## Usage 

### Generating environment files 

This sample call will write the environment variables to local.env file 

```shell
go run main.go -f=local.env -x=GOPATH,GOBIN,PATH,LS_COLORS
```

- -f=local.env : File name which will store the environment variables
- -x=GOPATH,GOBIN,PATH,LS_COLORS : Environment variables names to be excluded 

### Using the environment files in your VsCode project 

In order to pass the environment files in your project, you can follow this steps:

1. [Debugging Go using VS Code](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code) 

2. In your .vscode/launch.json, enter envFile property that points to the generated environment variables file.
  Sample launch.json file 

  ```json
  {
	"version": "0.2.0",
	"configurations": [
		{
			"name": "Launch",
			"type": "go",
			"request": "launch",
			"mode": "auto",
			"program": "${fileDirname}",
            "envFile": "${workspaceRoot}/local.env",
            "args": [
                "ARG1=VALUE1",
                "ARG2=VALUE2"
            ]
		}
	]
  } 
  ``` 

  Now you should be able to select main.go, run debug (Debug>Start Debugging). This will pass all the environment variables defined in local.env file to your application 


