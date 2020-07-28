package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

const name = "HAILUCK CO.,LTD Usb Touch Touchpad"

type Runner interface {
	Run(cmd []string) string
}

type XinputRunner struct {
}

func (x XinputRunner) Run(cmd []string) string {
	out, err := exec.Command("xinput", cmd...).Output()
	if err != nil {
		log.Fatal(err)
	}
	return fmt.Sprintf("%s", out)

}

func GetId(r Runner) string {
	re, err := regexp.Compile(name + `\s+id=(\d+)`)
	if err != nil {
		log.Fatal(err)
	}

	out := r.Run([]string{"list"})
	m := re.FindStringSubmatch(out)
	if len(m) == 0 {
		log.Fatalf("Cannot extract device ID:\n%s", out)
	}

	return m[1]

}

func SetProp(r Runner, id, prop, value string) {
	cmd := []string{"set-prop", id, prop, value}
	r.Run(cmd)
}

func SetButtonMap(r Runner, id, value string) {
	cmd := []string{"set-buttom-map", id, value}
	r.Run(cmd)
}

func main() {
	props := [][]string{
		{"290", "1"},   // Tapping Enabled
		{"292", "0"},   // Tapping Drag Enabled
		{"298", "1"},   // Natural Scrolling Enabled
		{"310", "0.5"}, // Accel Speed
	}

	r := XinputRunner{}
	id := GetId(r)
	for _, o := range props {
		SetProp(r, id, o[0], o[1])
	}
	SetButtonMap(r, id, "1 3 3") // avoiding 2 removes middle-buttom
}
