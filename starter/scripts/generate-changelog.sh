#!/usr/bin/env bash

if ! command -v bf3-uploader &> /dev/null
then
  echo "Installing git-chglog"
  go get -u github.com/git-chglog/git-chglog/cmd/git-chglog
fi

# output changelog between stable releases
git-chglog $(git describe --tags $(git rev-parse HEAD))