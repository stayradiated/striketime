#!/usr/bin/env bash

exec migrate -url postgres://striketime@localhost/striketime?sslmode=disable $1
