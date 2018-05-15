#!/bin/sh
find ../tests/song | xargs touch --date="2018/5/12 00:00"
find ../tests/song | grep not_updated | xargs touch --date="2018/5/12 12:00"

