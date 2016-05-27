cf-lager
========

A thin wrapper around [github.com/pivotal-golang/lager](https://github.com/pivotal-golang/lager) for easy use in CF components.

This library provides a flag called `logLevel`. By importing this library, various CF components can share the same name, description, and default value ("info") for this flag.

The logger returned by `cf_lager.New()` will write all logs to `os.Stdout`.

To use, simply import this package in your `main.go` and call `cf_lager.New(COMPONENT_NAME)` to get a logger.

For example:

```golang
package main

import (
    "flag"
    "fmt"

    "github.com/cloudfoundry-incubator/cf-lager"
    "github.com/pivotal-golang/lager"
)

func main() {
    cf_lager.AddFlags(flag.CommandLine)

    flag.Parse()

    logger, reconfigurableSink := cf_lager.New("my-component")
    logger.Info("starting")

    // Display the current minimum log level
    fmt.Printf("Current log level is ")
    switch reconfigurableSink.GetMinLevel() {
    case lager.DEBUG:
        fmt.Println("debug")
    case lager.INFO:
        fmt.Println("info")
    case lager.ERROR:
        fmt.Println("error")
    case lager.FATAL:
        fmt.Println("fatal")
    }

    // Change the minimum log level dynamically
    reconfigurableSink.SetMinLevel(lager.ERROR)
    logger.Debug("will-not-log")
}
```

Running the program above as `go run main.go --logLevel debug` will generate the following output:

```
{"timestamp":"1464388983.540486336","source":"my-component","message":"my-component.starting","log_level":1,"data":{}}
Current log level is debug
```
