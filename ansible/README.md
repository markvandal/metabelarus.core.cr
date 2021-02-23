# Как запустить тестовую версию mbcore

## Готовим виртуальную машину
1. Создайте виртуальную машину. Я тестировал на ubuntu-server 20.04. Сгенерируете ssh ключ.
2. Переименуте файл hosts.editme
```
mv hosts.editme hosts
```
3. В файле hosts укажите все необходимые данные для доступа к виртуальной машине.
4. Установите на машине Doker
```
ansible-playbook build-machine.yml
```
## Создаем образ
5. Создайте образ с бинарным файлом mbcorercd
```
ansible-playbook make_build_mbcore.yaml
```
6. На удаленной машине можно протестировать
```
docker run -t core 
```
