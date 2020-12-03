#!/bin/bash
set -e

case "$1" in
    bash)
        set -- "$@"
    ;;
    *)
        set -- ./r1wallet "$@"
    ;;
esac

exec "$@"
