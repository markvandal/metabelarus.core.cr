ROOT_SCRIPTPATH=$( cd "$(dirname "$0")" >/dev/null 2>&1 ; pwd -P )

CMD=$ROOT_SCRIPTPATH/../build/mbcorecrd
HOME=$ROOT_SCRIPTPATH/../build/

MONIKER=$1
PORTSHIFT=$2
HOMESUFFIX=$3

. $ROOT_SCRIPTPATH/.env

if [ -z "$MONIKER" ]
  then
    MONIKER="testnode"
fi

if [ -z "$CHAINID" ]
  then
    CHAINID="metabelarus.core.cr-test"
fi

if [ -z "$PORTSHIFT" ]
  then
    PORTSHIFT=0
fi

if [ ! -z "$HOMESUFFIX" ]
  then
    HOME=$HOME$HOMESUFFIX
  else
    HOME=$HOME"home"
fi

$CMD init $MONIKER --chain-id $CHAINID --home $HOME --overwrite --trace

sed -i '' 's/min-retain-blocks = 0/min-retain-blocks = 100/g' "$HOME/config/app.toml"
sed -i '' 's/enable = false/enable = true/g' "$HOME/config/app.toml"
sed -i '' 's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' "$HOME/config/app.toml"
sed -i '' 's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' "$HOME/config/config.toml"
sed -i '' 's/version = "v0"/version = "v2"/g' "$HOME/config/config.toml"
sed -i '' 's/addr_book_strict = true/addr_book_strict = false/g' "$HOME/config/config.toml"
sed -i '' 's/allow_duplicate_ip = false/allow_duplicate_ip = true/g' "$HOME/config/config.toml"
sed -i '' 's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "1m0s"/g' "$HOME/config/config.toml"

TCPORT=$(expr 1317 + $PORTSHIFT)
sed -i '' 's/tcp\:\/\/0.0.0.0\:1317/tcp\:\/\/0.0.0.0\:'$TCPORT'/g' "$HOME/config/app.toml"

GRPCPORT=$(expr 9090 + $PORTSHIFT)
sed -i '' 's/0.0.0.0\:9090/0.0.0.0\:'$GRPCPORT'/g' "$HOME/config/app.toml"

PROXIPORT=$(expr 26658 + $PORTSHIFT)
sed -i '' 's/tcp\:\/\/127.0.0.1\:26658/tcp\:\/\/127.0.0.1\:'$PROXIPORT'/g' "$HOME/config/config.toml"

RPCPORT=$(expr 26657 + $PORTSHIFT)
sed -i '' 's/tcp\:\/\/127.0.0.1\:26657/tcp\:\/\/127.0.0.1\:'$RPCPORT'/g' "$HOME/config/config.toml"

P2PPORT=$(expr 26656 + $PORTSHIFT)
sed -i '' 's/tcp\:\/\/0.0.0.0\:26656/tcp\:\/\/0.0.0.0\:'$P2PPORT'/g' "$HOME/config/config.toml"

PPROFPORT=$(expr 6060 + $PORTSHIFT)
sed -i '' 's/localhost\:6060/localhost\:'$PPROFPORT'/g' "$HOME/config/config.toml"


if [ -z "$HOMESUFFIX" ]
  then
    $CMD keys add mainuser --hd-path "44'/118'/0'/0/0" \
    --home $HOME --keyring-backend file --recover --trace
    $CMD add-genesis-account $($CMD keys show mainuser -a --home $HOME --keyring-backend file --trace) \
    1000token,100000000stake,20invitesuper,50invite0,50invite1,150invite2,150invite3 \
    --home $HOME --trace

    $CMD gentx mainuser --from mainuser --chain-id $CHAINID --amount 1000000stake \
    --keyring-backend file --home $HOME -y --trace

    $CMD collect-gentxs --home $HOME --trace
fi