#!/bin/bash
echo $$ > /tmp/swap_dev.pid

# Function to clean up processes on exit
cleanup() {
    echo "Cleaning up..."
    # Prevent recursive cleanup calls
    trap - SIGINT SIGTERM SIGQUIT
    pkill -TERM -x inotifywait 2>/dev/null || true
    pkill -TERM -x swap 2>/dev/null || true
    rm -f /tmp/swap_dev.pid
    reset
    exit 0
}

# Set up signal traps
trap cleanup SIGINT SIGTERM SIGQUIT

# Function to start the swap application
start_swap(){
    export TERM=xterm-256color
    echo "starting swap"
    go run -ldflags="-X main.APIURL=localhost:8080" ./cmd/swap --debug
    echo "swap process exited"
}

# Set up file monitoring in background that will kill the app when files change
(
    while true; do
        inotifywait -q -r -e modify ./app --include '\.go$'
        echo "File change detected, killing swap to trigger restart..."
        pkill -9 -x swap 2>/dev/null || true
    done
) &

# Start the app initially
start_swap

# Main loop - after app exits, restart it
while true; do
    echo "App exited, restarting..."
    sleep 0.5
    start_swap
done
