#!/bin/bash
echo $$ > /tmp/swap_dev.pid
cleanup() {
    echo "Ctrl+C detected. Cleaning up..."
    pkill -TERM -x inotifywait
    pkill -TERM -x swap
    pkill -P $$
    rm -f /tmp/swap_dev.pid
    exit 0
}
trap cleanup SIGINT SIGTERM
start_swap(){
    export TERM=xterm-256color
    echo "starting swap"
    go run -ldflags="-X main.APIURL=localhost:8080" ./cmd/swap --debug
}

start_tracker(){
    inotifywait -r -m -e modify ./app --include '\.go$' | while read -r event; do
        echo "Detected modification: $event"
        pkill -9 -x swap || true
        sleep 0.2
    done
}

(start_tracker)&
while :;do start_swap; sleep 1; done
