#!/bin/sh
if [[ $1 = "nobuild" ]]; then
    echo Skipping Build
else
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
tmux send-keys "./build/bin/airbloc server -d . -c config.yml --verbose" "C-m"
tmux split-window -v
tmux send-keys "./build/bin/airbloc userdelegate -c config-userdelegate.yml --verbose" "C-m"
tmux -2 attach-session -t airbloc

if [[ $1 = "test" ]]; then
    echo Removing airbloc session.
    tmux kill-session -t airbloc
fi
