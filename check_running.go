package main

import (
    "time"
    "bytes"
    "fmt"
    "os/exec"
    "strings"
    "./Config"
)

func getProcess(processName string) {
    cmd := exec.Command("/bin/sh", "-c", "ps aux | grep \"" + processName + "\" | grep -vc \"grep\"")
    cmd.Stdin = strings.NewReader("some input")
    var out bytes.Buffer
    cmd.Stdout = &out
    err := cmd.Run()

    if err != nil {
        startProcess(processName)
    }

    return
}

func startProcess(processName string) {
    start := exec.Command("/bin/sh", "-c", "sudo /etc/init.d/" + processName + " restart")

    start.Stdin = strings.NewReader("some input")
    var subOut bytes.Buffer
    start.Stdout = &subOut
    startErr := start.Run()

    if startErr != nil {
        fmt.Println("Error: " + startErr.Error())
    }

    return
}

func main() {
    for {
        config := Config.LoadConfig()

        for _, daemon := range config.Daemons {
            getProcess(daemon)
        }

        time.Sleep(time.Duration(config.Interval) * time.Minute)
    }
}
