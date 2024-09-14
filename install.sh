#!/bin/bash

TAG=""

# Parse command-line arguments
while getopts ":t:" opt; do
  case ${opt} in
    d ) TAG="$OPTARG";;
    \? ) echo "Usage: $0 -t <tag-version>"; exit 0;;
  esac
done

VERSION="${TAG:-@latest}"
REPO_URL="github.com/aless10/CLI-Todo-App-In-Go"

echo "Installing the todo cli app from ${REPO_URL}"

go install ${REPO_URL}${VERSION}

echo "CLI installed" 

CLI-Todo-App-In-Go -h

echo "If you want to improve your ux, set an alias for the command in your shell rc file."
