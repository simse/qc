#!/bin/bash
cd examples
echo "Resetting examples folder"
git clean -f
git clean -fd
git add -A
git reset --hard HEAD
git pull