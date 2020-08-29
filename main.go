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

func GetPropId(r Runner, id, s string) string {
	re, err := regexp.Compile(s + `\s\((\d+)\)`)
	if err != nil {
		log.Fatal(err)
	}

	out := r.Run([]string{"list-props", id})
	m := re.FindStringSubmatch(out)
	if len(m) == 0 {
		log.Fatalf("Cannot extract prop ID:\n%s", out)
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
		{"Tapping Enabled", "1"},
		{"Tapping Drag Enabled", "0"},
		{"Natural Scrolling Enabled", "1"},
		{"Accel Speed", "0.5"},
	}

	r := XinputRunner{}
	id := GetId(r)
	log.Output(2, "Device ID: "+id)

	prop := ""
	for _, o := range props {
		prop = GetPropId(r, id, o[0])
		log.Output(2, "Property "+o[0]+" ID: "+prop)
		SetProp(r, id, prop, o[1])
	}
	SetButtonMap(r, id, "1 3 3") // avoiding 2 removes middle-buttom
}
