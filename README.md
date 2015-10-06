# golang-worker-example
Example buffered channel + worker example

## Building & Running

Build the project:

```
go build -o queued .
```

To start the worker, run:
```
./queued
```

To issue commands, run:
```
curl -v -X POST "localhost:8000/work?delay=1s&name=foo"
```
