#!/bin/sh

ARGS=""

if [ ! -z "$BRAND" ]; then
    ARGS="-brand $BRAND"
fi

wiki $ARGS
