# Как запустить тестовую версию mbcore

## Готовим виртуальную машину
1. Создайте виртуальную машину. Я тестировал на ubuntu-server 20.04. Сгенерируете ssh ключ.
2. Переименуте файл hosts.editme в hosts
```
mv hosts.editme hosts
```
3. В файле hosts укажите все необходимые данные для доступа к виртуальной машине.
4. Установите на машине Doker
```
ansible-playbook build-machine.yml
```
## Создаем образ
 Создаем образ. Бинарник в /usr/bin/
```
ansible-playbook make_build_mbcore.yaml
```
## Запускаем контейнер
Логинимся на удаленную машину и проверяем
```
docker run -d -p 26657:26657 -p 26656:26656 -p 1317:1317 -v mbcore:/root/.mbcorecr  core 
curl 127.0.0.1:1317
```


