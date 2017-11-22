package main

import (
    "fmt"
    "time"
    "os/exec"
)

var(
    currentExercise = 1
    allExercises =  5
)

func main() {
    exercise := make(chan bool)
    shortRest := make(chan bool)
    longRest := make(chan bool)
    done := make(chan bool)

    go circuit(exercise)
    go circuitService(exercise, shortRest, longRest, done)

    <-done
}

func circuit(exerciseChan chan bool) {
    beginExercise()
    time.Sleep(time.Minute*3)
    endExercise()
    exerciseChan <- true
}

func shortRest(restChan chan bool) {
    beginShortRest()
    time.Sleep(time.Minute*1)
    endShortRest()
    restChan <- true
}

func longRest(longRestChan chan bool) {
    beginLongRest()
    time.Sleep(time.Minute*2)
    longRestChan <- true
}

func beginExercise() {exec.Command("say", "Exercise begins").Output()}

func endExercise() {exec.Command("say", "Exercise ends").Output()}

func beginShortRest() {}

func endShortRest() {}

func beginLongRest() {}

func endLongRest() {}

func circuitService(exerciseChan, shortRestChan, longRestChan, doneChan chan bool) {
    for {
        select {

        case endExercise := <-exerciseChan:
            _ = endExercise
            if currentExercise >= allExercises {
                go longRest(longRestChan)
                currentExercise = 1
            } else {
                currentExercise += 1
                go shortRest(shortRestChan)
            }

        case endShortRest := <-shortRestChan:
            _ = endShortRest
            go circuit(exerciseChan)

        case endLongRest := <-longRestChan:
            _ = endLongRest
            input := askUser()
            for input != "Y" && input != "N" {
                input = askUser()
            }
            if input == "Y" {
                go circuit(exerciseChan)
            } else {
                doneChan <- true
            }

        }
    }
}

func askUser() string {
    fmt.Println("Would you like to continue with another circuit? (Y/N)")
    var response string
    fmt.Scanln(&response)
    return response
}
