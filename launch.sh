#!/bin/sh
if [[ $1 != "test" ]]; then
    make airbloc &> ./build/airbloc.log
    if [[ $? -eq 0 ]]; then
        rm ./build/airbloc.log
        echo Build successful.
        sleep 1
    else
        cat ./build/airbloc.log
        echo Build failed. See log for details.
        exit 1
    fi
fi

# Launch Airbloc tmux
tmux new -d -s airbloc
tmux rename-window 'Airbloc'

if [[ $1 = "test" ]]; then
    tmux send-keys "go run ./cmd/airbloc/main.go" "C-m"
    tmux split-window -v
    tmux send-keys "go run ./cmd/airbloc/main.go userdelegate --config config-userdelegate.yml" "C-m"
else
    tmux send-keys "./build/bin/airbloc" "C-m"
    tmux split-window -v
    tmux send-keys "./build/bin/airbloc userdelegate --config config-userdelegate.yml" "C-m"
fi
tmux -2 attach-session -t airbloc

if [[ $1 = "test" ]]; then
    echo Removing airbloc session.
    tmux kill-session -t airbloc
fi
