# Как запустить ноду с блокчейном (fullnode - не валидатор)
## Готовим виртуальную машину
1. Создайте виртуальную машину. Мы тестировали на ubuntu-server 20.04. Сгенерируете ssh ключ.
2. Перейдите под консолью в папаку ansiable (в папке куда вы склонировали репозиторий с проектом)
3. Скопируйте папку inventory.editme в inventory
```
cp -fr inventory.editme inventory
```
4. В файле inventory/hosts укажите ip аддресс виртуальной машины и alias хоста:
   1. Например `mbcr ansible_host=139.177.179.141`, где `mbcr` - alias, где ip адресс - это адресс машины где разворачиваются ноды.
5. Создаём файл `inventory/host_vars/[host_alias].yml` на основе `testhost.yml`, например: `inventory/host_vars/mbcr.yml`
6. Конфигуриреуем `inventory/host_vars/[host_alias].yml`
   1. chainid: metabelarus.core.cr
   2. moniker - придумайте уникальный позывной (желательно проконсультироваться с тем кто помогает вам с запуском) 
   3. seeds: "`d6d5f2085565073badf1c811386d96fe72d78029@node-mark.cr.meta-belarus.org:26756`"
   4. mainuser_mnemonic: Вводим мнемонику своего аккаунта
   5.  source_host: "mark" — или другое название если вы получили его у того кто вас консультирует с запуском
      1. Если отличается source_host, получите файл `genesis.json` и `info.yml` создайте папку с именем из переменной source_host 
      2. Положите `genesis.json` и `info.yml` в папку `fetched/[source_host]/`
   6. domain: доменное имя которое настроена на ваш сервис
   7. host_type: "public"
7. Устанавливаем все необходимые зависимости и билдим на виртуалке приложение
```
ansible-playbook machine-build.yml
```
## Вариант с доменом
8. Запускаем инициализацию
```
ansible-playbook node-init.yml
```
9. Запустите скрипт для установки генезиса
```
ansible-playbook node-add.yml
```
10. Запускаем ноду
```
ansible-playbook node-start.yml
```
11. Ждём пока она выровняется по высоте с сетью
## Вариант без домена
Делаем всё тоже самое, что и для случая с доменом. Но:
1. Domain: указваем "localhost"
2. Host_type: указваем "local"

# Как запустить валидатора
1. Получите 1000000stake на аккаунт связанный с вашей мнемоникой у того кто вас консультирует
2. Запустите
```
ansible-playbook validator-init.yml
```
3. Дождитесь пока валидатор выровняется с сетью по высоте
4. Запустите
```
ansible-playbook node-start.yml
```

# Как запустить тестовую версию mbcorecr

## Готовим виртуальную машину
1. Создайте виртуальную машину. Мы тестировали на ubuntu-server 20.04. Сгенерируете ssh ключ.
2. Перейдите под консолью в папаку ansiable (в папке куда вы склонировали репозиторий с проектом)
3. Скопируйте папку inventory.editme в inventory
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
ansible-playbook node-genesis.yml
```
## Запускаем контейнеры с блокчейном
```
ansible-playbook root-start.yml
```