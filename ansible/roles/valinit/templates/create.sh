mbcorecrd tx staking create-validator \
  --amount {{ amount_to_delegate }} --from mainuser \
  --moniker {{ moniker }} --pubkey $(mbcorecrd tendermint show-validator) \
  -y --keyring-backend file --chain-id {{ chainid }} \
  --commission-max-change-rate 0.010000000000000000 \
  --commission-max-rate 0.200000000000000000 \
  --commission-rate 0.100000000000000000 \
  --min-self-delegation 1