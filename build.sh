#!/bin/bash

Version=$(git describe --tags)
BuildTime=$(date +'%Y%m%d %H:%M:%S %z')
ldflags=" \
    -X 'bgm-calendar/meta.Version=$Version' \
    -X 'bgm-calendar/meta.BuildTime=$BuildTime' \
"

go build -ldflags="$ldflags"