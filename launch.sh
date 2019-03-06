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

# ./build/bin/airbloc \
tmux send-keys "go run cmd/airbloc/*.go server --config \"config.yml\" --verbose --ethereum \"ws://localhost:8545\" --deployment \"http://localhost:8500\" --metadb \"mongodb://localhost:27017/airbloc\" --private \"0x3d690ff25a05e195d4ef8f65ac34ee2d5e71e41c414a25c48164499299fece40\"" "C-m"

if [[ $1 != "test" ]]; then
tmux split-window -v
tmux send-keys "./build/bin/airbloc userdelegate --config \"config-userdelegate.yml\" --verbose --ethereum \"ws://localhost:8545\" --deployment \"http://localhost:8500\" --metadb \"mongodb://localhost:27017/airbloc\" --private \"0xa0746c23a619f9c6a25057de7da0fb4f46dbd012a5a2fe5607b4f38e7ca5a3cc\"" "C-m"
fi

tmux -2 attach-session -t airbloc

if [[ $1 = "test" ]]; then
    echo Removing airbloc session.
    tmux kill-session -t airbloc
fi
