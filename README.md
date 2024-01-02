# mLock

[![Go Report Card](https://goreportcard.com/badge/github.com/Kimbbakar/mlock)](https://goreportcard.com/report/github.com/Kimbbakar/mlock)
[![GoDoc](https://pkg.go.dev/badge/github.com/Kimbbakar/mlock?status.svg)](https://pkg.go.dev/github.com/Kimbbakar/mlock?tab=doc)

`mlock` empowers users to enforce locks on the same critical section based on different types, providing a flexible and effective way to control access and maintain concurrency in diverse scenarios.

## Features

- Obtain locks based on different keys.
- Automatically cleans up unused locks at regular intervals.

## Installation

To use `mlock` in your Golang project, you can simply run:

```bash
go get -u github.com/Kimbbakar/mlock
```

## Usage

#### Importing the Package
```go
import "github.com/yourusername/mlock"
```

#### Initializing and Cleaning Up Locks
Before using mlock, it's recommended to set up a cleanup routine to remove unused locks periodically. This can be done using the KeepClean function:
```go
interval := 30 * time.Minute
mlock.KeepClean(&interval)
```

This sets up a cleanup routine to run every 30 minutes. You can customize the cleanup interval as needed.

#### Obtaining and Releasing Locks
To obtain a lock based on a key, use the Lock function:
```go
mlock.Lock("your_key_here")
defer mlock.Unlock("your_key_here")
```

#### Example
```go
package main

import (
	"fmt"
	"github.com/yourusername/mlock"
	"time"
)

func main() {
	// Set up cleanup routine
	interval := 30 * time.Minute
	mlock.KeepClean(&interval)

	// Obtain a lock
	mlock.Lock("example_key")
	defer mlock.Unlock("example_key")

	// Your critical section here
	fmt.Println("Lock acquired, performing operations...")
}
```

#### Contribution
Contributions are welcome! If you find any issues or have suggestions for improvement, feel free to open an issue or create a pull request.

# License
This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
