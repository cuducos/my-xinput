# My `xinput`

A simple program to apply my custom settings to my track-pad using `xinput`.

How I use it:

1. Build it with `go build main.go`
2. Change its permissions with `chmod a+x my-xinput`
3. Move it to `/usr/local/bin`
4. Adds `exec_always my-xinput` to `.config/i3/config`