ROOT_SCRIPTPATH=$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )

CMD=$ROOT_SCRIPTPATH/../build/mbcorecrd
HOME=$ROOT_SCRIPTPATH/../build/home
SECONDHOME=$ROOT_SCRIPTPATH/../build/$3

. $ROOT_SCRIPTPATH/.env

$CMD tx bank send \
 $($CMD keys show $1 -a --home $HOME --keyring-backend file) \
 $($CMD keys show $2 -a --home $HOME --keyring-backend file) \
 1000000stake --chain-id $CHAINID \
 --from mainuser -y --keyring-backend file --home $HOME

$CMD tx staking create-validator --amount 1000000stake --from $2 \
 --moniker $2 --pubkey $($CMD tendermint show-validator --home $SECONDHOME) \
 -y --keyring-backend file --chain-id $CHAINID --home $HOME \
 --commission-max-change-rate 0.010000000000000000 \
 --commission-max-rate 0.200000000000000000 \
 --commission-rate 0.100000000000000000 \
 --min-self-delegation 1