#!/bin/bash

export API_PORT=10999
export API_BIND=0.0.0.0
export LOGLEVEL=debug

go run init.go app.go

