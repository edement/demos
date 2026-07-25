1. Прежде чем работать с проектом нужно установить все необходимые пакеты
    cmd\bash `go mod tidy`

2. Если ругается на модуль golang-migrate или просто migrate, то скачать его отдельно
    cmd\bash `go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest`

3. Так же необходимо поднять базу данных Postgres через pgAdmin имя базы даных, пароль, юзер, порт указываются в конфиг файле, заменить в postgres на нужные

4. Сам запуск приложения
    cmd\bash `air`