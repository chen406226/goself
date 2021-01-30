package main

import (
    "fmt"
    "os/exec"
)

func FirstConnection()  {
    cmd := exec.Command("cmd.exe", "/c", "mstsc /v: 10.10.0.123 /console")
    cmd.Run()
    cmd = exec.Command("cmd.exe", "/c", "net use \\\\10.10.0.123\\ipc$ Nc@test /user:Administrator")
    cmd.Run()
}

func main() {
    FirstConnection()
    var SourceFolder = "E:\\other\\test\\test"
    //var stdout, stderr bytes.Buffer
    //cmd := exec.Command("cmd.exe", "/c", "echo 'hello' && echo 'hi'")
    //cmd.Stdout = &stdout
    //cmd.Stderr = &stderr
    //err := cmd.Run()
    //fmt.Println(err)
    //out := stdout.String() + stderr.String()
    //fmt.Printf(out)
    //var commands []*exec.Cmd
    //cmdLine := "cd E:\\other\\test\\test"
    //commands = append(commands,exec.Command("cmd.exe", "/c", "start " + cmdLine))
    //commands = append(commands,exec.Command("cmd.exe", "/c", "start " + "npm run build"))
    cmd := exec.Command("cmd.exe", "/c", "cd/d "+ SourceFolder +" && npm run build")

    //cmd.Run()
    //time.Sleep(time.Second * 2)
    //cmd = exec.Command("cmd.exe", "/c", "xcopy "+SourceFolder+"\\dist \\\\10.10.0.123\\D\\test /s/e/y")
    //cmd.Run()
    fmt.Println(cmd.Args)
    //cmd =
    //
    //cmd.Output()
    //cmd.StdoutPipe

    //for i:= 1;i < len(commands);i++{
    //    commands[i].Stdin, _ = commands[i-1].StdoutPipe()
    //}
    //commands[len(commands)-1].Stdout = os.Stdout
    //for i:=1;i<len(commands);i++{
    //
    //    err := commands[i].Start()
    //    if err != nil {
    //        panic(err)
    //    }
    //}
    //commands[0].Run()
    //for i:=1;i<len(commands);i++ {
    //
    //    err := commands[i].Wait()
    //    if err != nil {
    //        panic(err)
    //    }
    //}
    //cmd = exec.Command("cmd.exe", "/c", "start " + "call npm run build")
    //
    //err = cmd.Run()

    //fmt.Printf("%s, error:%v \n", cmdLine, err)
}


