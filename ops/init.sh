#!/bin/sh
rm initlog
rm -rf /root/.mbcorecr
. .env
echo "Наш пароль $DEBUG_PASSWD" > initlog
echo "Мнемоника: $DEBUG_AUTH" >> initlog
echo " ------------------------------- Начинаю инициализацию ----------------------" >> initlog
mbcorecrd init testnode --chain-id metabelarus.core.cr-test  --overwrite --trace 2>>initlog
echo " --- Создаю пользователя" >> initlog 
(echo $DEBUG_AUTH; echo $DEBUG_PASSWD; echo $DEBUG_PASSWD;) | mbcorecrd keys add mainuser --hd-path "44'/118'/0'/0/0" --keyring-backend file --recover --trace 2>> initlog
echo "Добавляю genesis-account" >> initlog
mbcorecrd add-genesis-account $( (echo  $DEBUG_PASSWD) | mbcorecrd keys show mainuser -a --keyring-backend file ) \
1000token,100000000stake,20invitesuper,50invite0,50invite1,150invite2,150invite3 2>>initlog
(echo  $DEBUG_PASSWD) | mbcorecrd gentx mainuser --from mainuser --chain-id metabelarus.core.cr-test --keyring-backend file 2>> initlog
mbcorecrd collect-gentxs 2>> initlog
echo "Правлю config"
sed -i  's/min-retain-blocks = 0/min-retain-blocks = 100/g' $HOME_NODE/config/app.toml
sed -i  's/enable = false/enable = true/g' $HOME_NODE/config/app.toml
sed -i  's/enabled-unsafe-cors = false/enabled-unsafe-cors = true/g' $HOME_NODE/config/app.toml
sed -i  's/cors_allowed_origins = \[\]/cors_allowed_origins = \["\*"\]/g' $HOME_NODE/config/config.toml
sed -i  's/laddr = "tcp:\/\/127.0.0.1:26657"/laddr = "tcp:\/\/0.0.0.0:26657"/g' $HOME_NODE/config/config.toml
sed -i  's/version = "v0"/version = "v2"/g' $HOME_NODE/config/config.toml
sed -i  's/addr_book_strict = true/addr_book_strict = false/g' $HOME_NODE/config/config.toml
sed -i  's/allow_duplicate_ip = false/allow_duplicate_ip = true/g' $HOME_NODE/config/config.toml
sed -i  's/create_empty_blocks_interval = "0s"/create_empty_blocks_interval = "60s"/g' $HOME_NODE/config/config.toml
echo "Запускаю демона"
mbcorcrd start



