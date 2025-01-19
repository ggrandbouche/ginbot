#!/bin/bash

set -e

if tmux list-windows | grep -q "Gin Game"; then
  tmux kill-window -t "Gin Game"
fi

tmux new-window -n "Gin Game"
sleep 0.1
tmux split-window -h "go run cmd/server/main.go"
sleep 0.1
tmux kill-pane -t 0
sleep 0.1
tmux split-window -v "go run cmd/client/main.go"
sleep 0.1
tmux split-window -h "go run cmd/client/main.go"

