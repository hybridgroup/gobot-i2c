# gobot-i2c

Gobot (http://gobot.io/) is a library for robotics and physical computing using Go

This library provides drivers for i2c devices (https://en.wikipedia.org/wiki/I%C2%B2C). You would not normally use this library directly, instead it is used by Gobot adaptors that have i2c support.

## Getting Started
Install the library with: `go get -u github.com/hybridgroup/gobot-i2c`

## Examples
```go
package main

import (
	"fmt"
	"github.com/hybridgroup/gobot"
	"github.com/hybridgroup/gobot-firmata"
	"github.com/hybridgroup/gobot-i2c"
)

func main() {
	firmata := new(gobotFirmata.FirmataAdaptor)
	firmata.Name = "firmata"
	firmata.Port = "/dev/ttyACM0"

	wiichuck := gobotI2C.NewWiichuck(firmata)
	wiichuck.Name = "wiichuck"

	work := func() {
		go func() {
			for {
				fmt.Println("joystick", gobot.On(wiichuck.Events["joystick"]))
			}
		}()
		go func() {
			for {
				fmt.Println("c", gobot.On(wiichuck.Events["c_button"]))
			}
		}()
		go func() {
			for {
				fmt.Println("z", gobot.On(wiichuck.Events["z_button"]))
			}
		}()
	}

	robot := gobot.Robot{
		Connections: []interface{}{firmata},
		Devices:     []interface{}{wiichuck},
		Work:        work,
	}

	robot.Start()
}
```
## Hardware Support
Gobot has a extensible system for connecting to hardware devices. The following i2c devices are currently supported:

- Wii Nunchuck Controller

More drivers are coming soon...

## Documentation
We're busy adding documentation to our web site at http://gobot.io/  please check there as we continue to work on Gobot

Thank you!

## Contributing
In lieu of a formal styleguide, take care to maintain the existing coding style. Add unit tests for any new or changed functionality.


## License
Copyright (c) 2013 The Hybrid Group. Licensed under the Apache 2.0 license.
