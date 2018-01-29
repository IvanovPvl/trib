## Trib

RESTful-сервис по вычислению чисел трибоначчи. В проекте используется простой итеративный алгоритм на базе пакета math/big.
Для ускорения используется кэширование с помощью Map.

Запуск:
```bash
git clone https://github.com/IvanovPvl/trib.git
cd trib
docker build -t trib .
docker run -p 3333:3333 --name trib-prod trib:latest
```

После запуска сервис доступен по маршруту вида /trib?number=

Запуск тестов:
```bash
docker build -t trib-test -f Dockerfile.test .
```

Также тесты можно запускать внутри контейнера:
```bash
docker run --rm -it --name trib-test trib-test:latest
go test
exit
```

Развернутый сервис доступен по адресу:
[http://185.143.172.248:3333/trib?number=1000](http://185.143.172.248:3333/trib?number=1000)