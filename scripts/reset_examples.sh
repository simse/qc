#!/bin/bash
cd examples
echo "Resetting examples folder"
git clean -f
git clean -fd
git reset --hard HEAD
git pull