## Charmil Logger
Charmil logger provides a unified way to log messages across the entire Cobra CLI. 

## Functions provided by logger
Logger provides multiple functions, which can be used at many use-cases like Info, Error etc. For now logger supports the following functions:
```go
Info(args ...interface{})
Error(args ...interface{})
Infof(format string, args ...interface{})
Errorf(format string, args ...interface{})
Infoln(args ...interface{})
Errorln(args ...interface{})
```

## How to use
When the factory is initialised, the logger is provided by default. To know about factory visit [here](charmil_factory.md)

```go
newFactory.Logger.Info("provide a message")
newFactory.Logger.Errorf("Error: %s", err)
newFactory.Logger.Infoln("provide a message")
```