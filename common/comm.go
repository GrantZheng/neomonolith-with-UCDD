package common

import (
	log "github.com/go-kit/kit/log"
	opentracinggo "github.com/opentracing/opentracing-go"
)

var Tracer opentracinggo.Tracer
var Logger log.Logger
