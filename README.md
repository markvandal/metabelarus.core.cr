# mbcorecr

**mbcorecr** is a blockchain application built using Cosmos SDK and Tendermint and generated with [Starport](https://github.com/tendermint/starport).

## Development mode
### Install starpot correctly
1. Clone starport from https://github.com/tendermint/starport
2. It's higly recommended to patch: `starport/services/chain/init.go`
   
   With the following changes:
   
   Line 124: commands.AddAccount(ctx, account.Name, "")

   --> commands.AddAccount(ctx, account.Name, account.Mnemonic)

   This way one will be able to use the following mnemonics for the main genesis account: `tomato father silver rebuild volume guard donor tattoo spike staff want inherit diagram wrap teach furnace spring squeeze tuition march card foil random west`
3. Build starport with `make`
### Start development
1. Just run `starport serve` in the current dirrectory
2. Do not hesitate to run `starport serve -p` if something goes wrong
3. Run tests with `./dev/scripts/test_all.sh`

## Nodes configuration flow (pre-producation testing)
You need to make sure that one can create more than one validator node and things will still working OK.
0. Copy `./dev/.env.tpl` file to `./dev/.env` 
1. Build the app `./dev/build.sh `, the result will get to the `./build` folder.
2. Initialize the first blockchain node `./dev/init.sh`.

    One will need to provide a mnemonic string. Do not hasitate to use one from the Starport installation part.
3. Start this blockchain node `./dev/light.sh home start`
4. Initialize the scond blockchain node `./dev/init.sh testnode1 100 home1`
5. Override `./build/home1/config/gensis.json` with `./build/home/config.genesis.json` 
6. Add the first node address to the second node config:
   1. Find the node id `./dev/light.sh home tendermint show-node-id`
   2. Open `./build/home1/config/config.toml` and find `persistent_peers`
   3. Set value to `[node id you found]@127.0.0.1:26656`
7. Run the second blockchain `./build/mbcorecrd start --home [path to your project]/build/home1`
8. Create new user in the keyring of the first node
   1. `./dev/keys.sh home add slave`
   2. Pass tokens and create validator from first user to the second one `./dev/add-validator.sh mainuser slave home1`
9. Do not forget to remove persistent_peers from the config of second node after the things up and running.