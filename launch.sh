#!/bin/sh

make airbloc &> ./build/airbloc.log
if [ $? -eq 0 ]; then
    rm ./build/airbloc.log
    echo Build successful.
    sleep 1
else
    cat ./build/airbloc.log
    echo Build failed. See log for details.
    exit 1
fi

# Launch Airbloc tmux
tmux new-session -d -s airbloc '/usr/local/bin/zsh'
tmux send-keys "./build/bin/airbloc userdelegate --config config-userdelegate.yml" "C-m"
tmux rename-window 'Airbloc'
tmux split-window -h '/usr/local/bin/zsh'
tmux send-keys "./build/bin/airbloc" "C-m"
tmux -2 attach-session -t airbloc

