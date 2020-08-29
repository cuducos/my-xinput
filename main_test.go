package main

import "testing"

const list = `
⎡ Virtual core pointer                    	id=2	[master pointer  (3)]
⎜   ↳ Virtual core XTEST pointer              	id=4	[slave  pointer  (2)]
⎜   ↳ HAILUCK CO.,LTD Usb Touch Touchpad      	id=19	[slave  pointer  (2)]
⎣ Virtual core keyboard                   	id=3	[master keyboard (2)]
    ↳ HAILUCK CO.,LTD Usb Touch               	id=11	[slave  keyboard (3)]
    ↳ HAILUCK CO.,LTD Usb Touch Consumer Control	id=13	[slave  keyboard (3)]
    ↳ HAILUCK CO.,LTD Usb Touch Wireless Radio Control	id=16	[slave  keyboard (3)]
`

const listProps = `
Device 'HAILUCK CO.,LTD Usb Touch Touchpad':
	Device Enabled (153):	1
	Coordinate Transformation Matrix (155):	1.000000, 0.000000, 0.000000, 0.000000, 1.000000, 0.000000, 0.000000, 0.000000, 1.000000
	libinput Tapping Enabled (290):	0
	libinput Tapping Enabled Default (291):	0
	libinput Tapping Drag Enabled (292):	1
	libinput Tapping Drag Enabled Default (293):	1
	libinput Tapping Drag Lock Enabled (294):	0
	libinput Tapping Drag Lock Enabled Default (295):	0
	libinput Tapping Button Mapping Enabled (296):	1, 0
	libinput Tapping Button Mapping Default (297):	1, 0
	libinput Natural Scrolling Enabled (298):	0
	libinput Natural Scrolling Enabled Default (299):	0
	libinput Disable While Typing Enabled (300):	1
	libinput Disable While Typing Enabled Default (301):	1
	libinput Scroll Methods Available (302):	1, 1, 0
	libinput Scroll Method Enabled (303):	1, 0, 0
	libinput Scroll Method Enabled Default (304):	1, 0, 0
	libinput Click Methods Available (305):	1, 1
	libinput Click Method Enabled (306):	1, 0
	libinput Click Method Enabled Default (307):	1, 0
	libinput Middle Emulation Enabled (308):	0
	libinput Middle Emulation Enabled Default (309):	0
	libinput Accel Speed (310):	0.500000
	libinput Accel Speed Default (311):	0.000000
	libinput Left Handed Enabled (312):	0
	libinput Left Handed Enabled Default (313):	0
	libinput Send Events Modes Available (275):	1, 1
	libinput Send Events Mode Enabled (276):	0, 0
	libinput Send Events Mode Enabled Default (277):	0, 0
	Device Node (278):	"/dev/input/event23"
	Device Product ID (279):	9610, 8214
	libinput Drag Lock Buttons (314):	<no items>
	libinput Horizontal Scroll Enabled (315):	1
`

type MockRunner struct {
}

func (m MockRunner) Run(cmd []string) string {
	switch cmd[0] {
	case "list":
		return list
	case "list-props":
		return listProps
	default:
		return ""
	}
}

func TestGetId(t *testing.T) {
	mock := MockRunner{}
	if id := GetId(mock); id != "19" {
		t.Errorf("Expected 19, but got %s", id)
	}

}

func TestGetPropId(t *testing.T) {
	mock := MockRunner{}
	if id := GetPropId(mock, "1", "Tapping Enabled"); id != "290" {
		t.Errorf("Expected 290, but got %s", id)
	}
}
