package main

// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

import "github.com/jacobsa/go-serial/serial"
import "fmt"
import "log"
import "bufio"
import "strings"
import "os/exec"

// Sample output from the Andino X1's serial port:
//
// :0294{0000,0000}{0,0}
// :0295{0000,0000}{0,0}
// :0296{0000,0000}{0,0}
// :0297{0001,0000}{1,0}
// :0298{0001,0000}{1,0}
// :0299{0001,0000}{1,0}
//
// I think the first field after the : is the timestamp.  We care about
// the transition in the 4th line, where it goes from {0,0} to {1,0}.
// That indicates that the power-loss signal appeared on the first input.
// Once that happens, we want to trigger a system shutdown.
//
// The Andino UPS's firmware supposedly handles the power loss/shutdown/recovery
// race case correctly, so we don't have to do anything about it.

func main() {
	options := serial.OpenOptions{
		PortName:        "/dev/ttyAMA0",
		BaudRate:        38400,
		DataBits:        8,
		StopBits:        1,
		MinimumReadSize: 4,
	}

	port, err := serial.Open(options)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	defer port.Close()

	fmt.Println("Opened.")

	scanner := bufio.NewScanner(port)
	for scanner.Scan() {
		line := scanner.Text()
		if line[0] == ':' {
			s := strings.Split(line, "{")
			if (len(s) == 3) && (s[2][0:2] == "1,") {
				fmt.Printf("Power loss, attempting system shutdown!")
				cmd := exec.Command("/sbin/poweroff")
				err := cmd.Run()
				if err != nil {
					log.Fatal(err)
				}
			}
		} else {
			fmt.Printf("Read unknown line: %s\n", line)
		}
	}
}
