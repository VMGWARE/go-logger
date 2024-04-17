# go.vmgware.dev/logger

## Introduction

`go.vmgware.dev/logger` is a powerful logging utility crafted for Go applications. It provides color-coded console outputs and supports multiple log levels, facilitating effective debugging and monitoring in both development and production environments.

## Features

- **Multiple Log Levels:** Supports detailed categorization with DEBUG, INFO, WARN, and ERROR levels.
- **Colorful Console Outputs:** Enhances log readability with color-coded messages, improving visibility for better troubleshooting.
- **Formatted Messages:** Offers methods to log messages with formatting, allowing dynamic content integration.
- **Concurrency-Safe:** Designed to be safe for concurrent use within Go applications, ensuring log integrity across multiple goroutines.
- **Easy to Integrate:** Simple setup and customizable, making it an excellent choice for any Go project.

## Installation

```bash
go get go.vmgware.dev/logger
```

## Usage

After installing, import the logger into your project and set up the logging as needed:

```go
package main

import (
 "go.vmgware.dev/logger"
)

func main() {
 logger.Setup(logger.DEBUG, "app.log")
 defer logger.Close()

 // Log messages at various levels
 logger.Info("MAIN", "Application is starting")
 logger.Warn("MAIN", "Running on deprecated version")
 logger.Error("DATABASE", "Failed to connect to the database")
 logger.Debugf("DATABASE", "Connection attempt to %s failed", "server.local")
}
```

## Configuration

The logger is ready to use with minimal configuration but can be fully customized to fit the needs of any project:

- Set global log levels to control output verbosity.
- Configure log file paths for persistent logging.

## Acknowledgements

Inspired by Uptime Kuma's robust logging mechanism, adapted and enhanced for the Go environment.

## License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.

## Source

The original inspiration for this logger can be found in the Uptime Kuma repository: [Uptime Kuma Log Source](https://github.com/louislam/uptime-kuma/blob/master/src/util.ts)
