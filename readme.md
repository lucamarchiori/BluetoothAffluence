
## Hardware
The first challenge is about choosing the right device, the programming languages and the libraries that the scanner will use.
For hardware the choice is on two different platform:
- Arduino
- Raspberry Pi

Arduino supports only C++, needs additional hardware modules to acquire networking capabilities (except for Arduino R4 WiFi). Apart from this, it is a cheaper board, meaning it can be easier to deploy multiple scanners, and it is  more bug-free since it is just a micro-controller that don't deal with OS and software updates.

Raspberry Pi, on the other hand, is more expensive but it does include Bluetooth and WiFi capabilities, furthermore the Nano version can be considered to save some money. This board is a full computer so it has an operating system and supports almost all programming languages.

For this project, since it involves building a prototype, i have decide to use the Raspberry Pi since it is more flexible and allows for faster development due to the set of tools included with the OS. In production, Arduino-like board would be a more appropriate choice because of the reliability of micro-controllers. In that case, the source code of the prototype can be translated in C++ and deployed on those boards.

## Software
In my experience i have used some tools to analyze network device vulnerabilities. On of those is Bettercap.
From the Bettercap repository:
Bettercap is a powerful, easily extensible and portable framework written in Go which aims to offer to security researchers, red teamers and reverse engineers an easy to use, all-in-one solution with all the features they might possibly need for performing reconnaissance and attacking WiFi networks, Bluetooth Low Energy devices, wireless HID devices and Ethernet networks.

Since this is a very used tool and since i am looking for some language that can interact with bluetooth, my choice is to use the Go programming language and the [TinyGo Bluetooth library](https://github.com/tinygo-org/bluetooth#go-bluetooth) that is a cross-platform package for using hardware from the Go programming language.


## Roadmap
- Client: Send Json Request with data
- Client: chronjob and auotomation
- Client: tests
- Server: Webserver
- Server: Routes
- Server: Database
- Server: Data POST
- Server: FrontEnd