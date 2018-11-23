#!/bin/sh
# Go build trick— fake GOPATH environment, borrowed from go-ethereum :P

name="airbloc-go"
set -e

if [ ! -f "env.sh" ]; then
    echo "$0 must be run from the root of the repository."
    exit 2
fi

# Create fake Go workspace if it doesn't exist yet.
fakegopath="$PWD/build"
root="$PWD"
workdir="$fakegopath/src/github.com/airbloc"
if [ ! -L "$workdir/$name" ]; then
    mkdir -p "$workdir"
    cd "$workdir"
    ln -s ../../../../. $name
    cd "$root"
fi

# Set up the environment to use the fakegopath.
GOPATH="$fakegopath"
export GOPATH

# Run the command inside the fakegopath.
cd "$workdir/$name"
PWD="$workdir/$name"

# Launch the arguments with the configured environment.
exec "$@"
