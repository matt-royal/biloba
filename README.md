# biloba

## Improved test output in GoLand
biloba can help you get improved ginkgo test output in GoLand.
To try it out, replace your existing calls to 

`RunSpecs(t, "My Suite")`

with the following... 

```go
RunSpecsWithDefaultAndCustomReporters(t, "My Suite", biloba.GoLandReporter())
```

Outside of GoLand the output will be unchanged. When running tests from GoLand, it will add an additional reporter to 
help GoLand (via `go tool test2json`) parse individual ginkgo tests
