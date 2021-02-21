#!/bin/sh

ROOT_SCRIPTPATH=$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )

CMD=$ROOT_SCRIPTPATH/../build/mbcorecrd
HOME=$ROOT_SCRIPTPATH/../build/$1

. $ROOT_SCRIPTPATH/.env

$CMD q ${@:2} --home $HOME --chain-id $CHAINID