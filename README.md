# envtofile

Go extension for VS Code runs in a separete shell session, therefore, it does not get environment variables from VS Code terminal. A simple use case can be:
1. You have a shell script to set environment variables for dev/local/test/ etc.. You need to run this script before running your go app
2. You launch the debug mode from VSCode, however, it wont get the environment variables from the step 1 as VS Code extensions (in this case it is Go extension) run in a separete terminal instance

I order to resolve this, we in this code, we are getting all the environment variables from the current terminal session, and writing them to a file. In this way we can easily set the file location in VS Code's launch.json file, which will make the env variables available for the Debug mode. 

Note that I might be wrong, may be there is a better/easier way to set pull environment variables for the Debug mode.


