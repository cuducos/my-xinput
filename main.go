package main

import (
	"fmt"
	"log"
	"os/exec"
	"regexp"
)

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

func GetId(r Runner, s string) string {
	re, err := regexp.Compile(s + `\s+id=(\d+)`)
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
		log.Fatalf("Cannot extract prop ID for " + s)
	}

	return m[1]
}

func SetProp(r Runner, id, prop, value string) {
	cmd := []string{"set-prop", id, prop, value}
	r.Run(cmd)
}

func SetButtonMap(r Runner, id, value string) {
	cmd := []string{"set-button-map", id, value}
	r.Run(cmd)
}

func main() {
	ns := []string{"HAILUCK CO.,LTD Usb Touch Touchpad", "ELAN1300:00 04F3:30BE Touchpad"}
	ps := [][]string{
		{"Tapping Enabled", "1"},
		{"Tapping Drag Enabled", "0"},
		{"Natural Scrolling Enabled", "1"},
		{"Accel Speed", "0.5"},
	}
	r := XinputRunner{}

	for _, n := range ns {
		log.Output(2, "Configuring "+n)
		id := GetId(r, n)
		log.Output(2, "Device ID: "+id)

		p := ""
		for _, o := range ps {
			p = GetPropId(r, id, o[0])
			log.Output(2, "Property "+o[0]+" ID: "+p)
			SetProp(r, id, p, o[1])
		}
		SetButtonMap(r, id, "1 3 3") // removes middle-button
		log.Output(2, "Configured the button map")
	}
}
