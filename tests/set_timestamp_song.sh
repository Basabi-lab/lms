#!/bin/sh
find . | xargs touch --date="2018/5/12 00:00"
touch --date="2018/5/12 18:00" ./song/not_updated1.mp3
touch --date="2018/5/12 18:00" ./song/not_updated2.mp3

