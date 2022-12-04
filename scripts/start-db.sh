#!/usr/bin/env bash

echo "Starting Surreal db..."
surreal start --log trace --user root --pass root memory
