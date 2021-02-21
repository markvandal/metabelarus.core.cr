#!/bin/sh

ROOT_SCRIPTPATH=$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )

cd $ROOT_SCRIPTPATH

cd ../
make
mkdir ./build -p
mv mbcorecrd ./build