package main

import "fmt"
// go does not care about your name
// go care about can you do this job or not	


type Player interface {
	play()
}

type FootballPlayer struct {
	Name string
}

type CricketPlayer struct {
	Name string
}

func (fp *FootballPlayer) play() {
	fmt.Println(fp.Name + " is playing football")
}

func (cp *CricketPlayer) play() {
	fmt.Println(cp.Name + " is playing football")
}



type Speaker interface {
	speak(msg string)
}

type Person struct {
	name string
}

func (p *Person) speak(msg string) {
	fmt.Println(p.name + " says " + msg)
}

func SpeakAlphabets(via Speaker) {
	via.speak("A")
}
func main() {
	mat := new(Person)
	mat.name = "Mat"
	SpeakAlphabets(mat)
}