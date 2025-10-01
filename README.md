Following this course for the project, handcoding everything instead of copy and paste.

https://www.jetbrains.com/guide/go/tutorials/cli-apps-go-cobra/creating_cli/

Run "go install" in terminal to use the exe globaly!
you shuold be able to type CLIgo.git with any command

EX. CLIgo.git add 4 4

addition of 4 and 4 equals 8.

Executed

If you don't want to type .git at the end here are simple instructions to fix that
1. Remove CLIgo.git.exe if still there
2. In go.mod remove .git in ln 1
3. In main.go remove .git in ln 6
4. type "go build"
5. type "go install" to make it global

Now you should be able to type "CLIgo add 4 4"

If needs updates just let me know!