ROOT_SCRIPTPATH=$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )

CMD=$ROOT_SCRIPTPATH/../build/mbcorecrd
HOME=$ROOT_SCRIPTPATH/../build/

$CMD 

#$CMD tx staking create-validator --home $HOME \
# --moniker $1 --pubkey $2 --amount 