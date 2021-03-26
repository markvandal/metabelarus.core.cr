# Как запустить тестовую версию mbcorecr

## Готовим виртуальную машину
1. Создайте виртуальную машину. Мы тестировали на ubuntu-server 20.04. Сгенерируете ssh ключ.
2. Скопируйте папку inventory.editme в inventory
```
cp -fr inventory.editme inventory
```
3. В файле inventory/hosts укажите ip аддресс виртуальной машины и alias хоста.
   1. Создайте файл inventory/host_vars/[alias].yml
   2. Настройте согласно своим предпочтениям
   3. Убедитесь что host_type - стоит local, если вы не хотите использовать certbot для автоматического подключения ssl сертификатов.
4. Устанавливаем все необходимые зависимости и билдим на виртуалке приложение
```
ansible-playbook machine-build.yml
```
## Производим генезис первого валидатора блокчейн
```
ansible-playbook node-init.yml
```
## Запускаем контейнеры с блокчейном
```
ansible-playbook node-start.yml
```

# Как запустить ноду с блокчейном (fullnode - не валидатор)
## Вариант с доменом
1. Проведите настройку согласно пунктам 1-4 раздела (до пункта machine-build включительно), не делайте node-init и node-start [Как запустить тестовую версию mbcorecr/Готовим виртуальную машину](#готовим-виртуальную-машину)
2. Не забудьте указать домен и host_type: "public" в `inventory/host_vars/p[host_alias].yml`
3. Chain id: metabelarus.core.cr
4. Moniker - придумайте уникальный позывной (желательно проконсультироваться с тем кто помогает вам с запуском) 
5. Seeds: "`13e91a7a0109526a4a95663fea4c6de943d12fa8@node-mark.cr.meta-belarus.org:26756`"
6. Mainuser_mnemonic: Вводим мнемонику своего аккаунта
7. Source_host: "mark" — или другое название если вы получили его у того кто вас консультирует с запуском
8. Запускаем инициализацию без генезиса
```
ansible-playbook node-init.yml --skip-tags genesis
```
3. Получите файл `genesis.json` и `info.yml` создайте папку с именем из переменной source_host 
4. Положите `genesis.json` и `info.yml` в папку `fetched/[source_host]/`
5. Запустите скрипт для установки генезиса
```
ansible-playbook node-add.yml
```
6. Запускаем ноду
```
ansible-playbook node-start.yml
```
7. Ждём пока она выровняется по высоте с сетью
## Вариант без домена
Делаем всё тоже самое, что и для случая с доменом. Но:
1. Domain: указваем "localhost"
2. Host_type: указваем "local"

# Как запустить валидатора
1. Получите 1000000stake на аккаунт связанный с вашей мнемоникой у того кто вас консультирует
2. Запустите
```
ansible-playbook validator-up.yml
```
3. Дождитесь пока валидатор выровняется с сетью по высоте
4. Запустите
```
ansible-playbook validator-init.yml
```