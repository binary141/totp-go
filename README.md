# Time-Based One-Time Password (TOTP) Algorigthm

## What is this?
A Golang implementation of [RFC 6238](https://datatracker.ietf.org/doc/html/rfc6238). Includes helpers that aren't defined in the RFC to make it easier to use / manage in code.

## Install
Requires >=1.24.2

```sh
go get github.com/binary141/totp-go@v1.0.0
go get github.com/binary141/hotp-go@v1.0.0
```
## Usage
```golang
package main

import (
	"fmt"

	"github.com/binary141/hotp-go"
	"github.com/binary141/totp-go"
)

func main() {

	secret := "12345678901234567890"
	encodedSecret := hotp.EncodeSecret([]byte(secret))
	decodedSecret, err := hotp.DecodeSecret(encodedSecret)
	if err != nil {
		panic(err)
	}

	digits := 6
	var t0 int64 = 0
	var timeStep int64 = 30

	otp := totp.CreateTotp(decodedSecret, digits, t0, timeStep, "LABEL")

	code, err := otp.Calculate()
	if err != nil {
		panic(err)
	}

	fmt.Println("code: ", code)
}
```
