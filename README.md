## ⚽️ 🥅 goals 
goals is a CLI tool that encourages you to set goals for youself and enables you to track these goals in the terminal. It can also be used as a habit tracker or a todo list.

## ⚡️ Quick start

First, [download](https://golang.org/dl/) and install **Go** if you haven't already.

Installation of `goals` is done by using the [`go install`](https://golang.org/cmd/go/#hdr-Compile_and_install_packages_and_dependencies) command. 

```bash
go install github.com/yuanb01/goals@latest
```
The binary should be installed to your `$GOPATH/bin` by default. To be able to call the `goals` binary from any directory in your terminal, you will need to add the directory where the binary is located to your PATH environment variable. For ex., I ran the following:

```
export PATH=$PATH:~/go/bin
```

That's all you need to know to start! 🎉

## ⚙️ Usage Documentation
```
Here is a sample usage of this CLI app:

$ goals
⚽️ 🥅 goals is a goals tracker right in your terminal!

Usage:
  goals [command]

Available Commands:
  add <goal-name> [repeat]  Add a new goal to your goals list with optional [repeat] param
  delete <goal-name>        Delete a goal from your goals list by its name
  delete <goal-number>      Delete a goal from your goals list by its number in the goals list
  do <goal-name>            Mark a goal on your goals list as complete by its name
  do <goal-number>          Mark a goal on your goals list as complete by its number in the goals list
  list                      List all of your goals

Use "goals [command] --help" for more information about a command.

$ goals add french 5
Added "french" to your goals list with repeat 5x!

$ goals add gym
Added "gym" to your goals list!

$ goals list
You have the following goals:
1. french 5
2. gym

$ goals do gym
Yay! You have completed your "gym" goal! 🎉

$ goals list
You have the following goals:
1. french 5

$ goals do 1
You are making progress towards your "french" goal! You need to do this goal 4 more times 🔁

$ goals do french
You are making progress towards your "french" goal! You need to do this goal 3 more times 🔁

$ goals list
You have the following goals:
1. french 3

$ goals delete french
Deleted goal "french"

$ goals list
Your goals list is empty! Why not add a goal? 📝 🥅
```

## ⚠️ License

`goals` is licensed under the [MIT License](https://github.com/yuanb01/goals/blob/main/LICENSE).