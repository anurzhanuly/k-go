# в образе для сборки alpine-dev есть компилятор go, make, go-утилиты gotestsum, migrate, air
ARG build_image=docker-hub-utils.kolesa-team.org:5000/build/golang:1.16-alpine-dev
ARG base_image=docker-hub-utils.kolesa-team.org:5000/base/alpine:latest

# Стадия сборки - для разработки, тестирования и компиляции
FROM ${build_image} as build_stage

# локальном скачиваем зависимости напрямую,
# а при сборке в бамбу подменим goproxy на кэширующий прокси
ARG goproxy=https://proxy.golang.org|direct
ENV GOPROXY=${goproxy} \
    GOPRIVATE=*.kolesa-team.org \
    GO111MODULE=on

WORKDIR /code

# если для вашего проекта нужно доустановить пакеты или библиотеки - делайте это здесь

# установка библиотек
COPY go.mod go.sum ./
RUN go mod download -x

# компиляция
COPY . .
RUN go build -v -o out/binary

# опционально: если вы ведёте разработку в docker-compose и хотите иметь горячую перезагрузку при изменении кода,
# то эта команда будет запускать сервер в режиме live reload (см. пример файла .air.conf на соседней странице)
# ENTRYPOINT air -c .air.conf

# Финальная стадия для запуска готовой программы
FROM ${base_image}

# принимаем информацию о сборке в аргументах --build-arg revision=... и т.д.
ARG revision
ARG branch
ARG build_number
ARG build_url
ARG build_date
# устанавливаем переменные окружения на основе переданных аргументов.
# если ваш проект использует библиотеку https://stash.kolesa-team.org/projects/GL/repos/build-env/browse,
# то эти переменные окружения будут автоматически прочитаны и доступны через env.GetBuildInfo().
# в ином случае нужно написать самостоятельное считывание и вывод этих переменных по адресу /release
ENV RELEASE_REVISION=${revision} \
    RELEASE_BRANCH=${branch} \
    RELEASE_BUILD_NUMBER=${build_number} \
    RELEASE_BUILD_URL=${build_url} \
    RELEASE_BUILD_DATE=${build_date}

ENV CONFIG_FILE=/config/local.toml

# копируем программу из предыдущей стадии
COPY --from=build_stage /code/out/binary /usr/local/bin/

# Загружаем конфиги
ADD config /config

# Запуск сервера
ENTRYPOINT /usr/local/bin/binary serve --config=${CONFIG_FILE}