// +build !dynamic
// +build musl

// This file was auto-generated by librdkafka/bundle-import.sh, DO NOT EDIT.

package kafka

// #cgo CFLAGS: -I${SRCDIR}
// #cgo LDFLAGS: ${SRCDIR}/librdkafka/librdkafka_musl_linux.a  -lm -ldl -lpthread -lrt -lpthread -lrt
import "C"

// LibrdkafkaLinkInfo explains how librdkafka was linked to the Go client
const LibrdkafkaLinkInfo = "static musl_linux from librdkafka-static-bundle-v1.4.0.tgz"
