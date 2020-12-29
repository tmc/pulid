# pulid

[![Project status](https://img.shields.io/github/release/tmc/pulid.svg?style=flat-square)](https://github.com/tmc/pulid/releases/latest)
[![Build Status](https://github.com/tmc/pulid/workflows/Test/badge.svg)](https://github.com/tmc/pulid/actions?query=workflow%3ATest)
[![Go Report Card](https://goreportcard.com/badge/tmc/pulid?cache=0)](https://goreportcard.com/report/tmc/pulid)
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/tmc/pulid)

Prefixed Universally Unique Lexicographically Sortable Identifier

Wraps [oklog/ulid](https://github.com/oklog/ulid) and adds a two-character prefix on ids.

This is similar to the pattern that Twilio uses in their entity IDs.

This allows encoding of entity types into identifiers.
