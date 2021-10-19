# Лекция 5. Тестирование и конкурентность в go.
# Тестирование

Документация: https://pkg.go.dev/testing

Содержание пакета test:
- client.http - запросы для удобства ручного тестирования
- support - пакет поддержки фунцкиоанльного тестирования
- app/advert/service - функциональные тесты для сервиса по работе с объявлениями

Полезные команды:
```
// Запуск тестов
CONFIG_PATH=${PWD}/config/local.toml go test ./... -v

// Запуск тестов с проверкой на гонку
CONFIG_PATH=${PWD}/config/local.toml go test ./... -v -race

// Запуск тестов по названию
CONFIG_PATH=${PWD}/config/local.toml go test ./... -v -run TestAdvertService

// Запуск тестов с указанием количества параллельных тестов
CONFIG_PATH=${PWD}/config/local.toml go test ./... -v -p 1

// Запуск тестов и сбор покрытия (видит покрытие только, если тесты лежат вместе с пакетом)
CONFIG_PATH=${PWD}/config/local.toml go test ./... -v -coverprofile cover.out
go tool cover -html=cover.out

// Запуск проверки кода на типичные ошибки
go vet -v ./...

// Запуск бенчмарка просмотров
ab -n 1000 -c 10 -m POST http://localhost:8081/advert/view/1
```

Полезные ссылки:
- Табличные тесты (https://dave.cheney.net/2019/05/07/prefer-table-driven-tests)
- Подтесты (https://ieftimov.com/post/testing-in-go-subtests/)
- Пакет testify (https://github.com/stretchr/testify)

## Виды тестов
### Unit тесты
Тесты, которые проверяют выполнение функций, методов, классов в изолированной среде.
Такие тесты не должны ходить в сторонние сервисы, БД, использовать внешние ресурсы,
поэтому широко распространенно использование моков для подмены зависимостей в целях тестирования.

Чаще всего пишутся и используются разработчиками для самопроверки, до или после написания кода.
Запускаются автоматически во время сборки. Содержатся в кодовой базе проекта.

Пример:
- messages (https://stash.kolesa-team.org/projects/GL/repos/messages/browse/app/usecase/component/template/template_test.go)
- wallet (https://stash.kolesa-team.org/projects/WAL/repos/wallet/browse/pkg/mysql/insert_test.go)
- ms-krpro (https://stash.kolesa-team.org/projects/GL/repos/ms-krpro/browse/internal/service/agent_test.go)
- krpro-autodial (https://stash.kolesa-team.org/projects/MSKR/repos/krpro-autodial/browse/internal/domain/call/processes_test.go)

### Functional тесты
Тесты, которые проверяют код с использованием всех зависимостей.
Обычно тестируют функционал по методу черного ящика, используя внешние интерфейсы модулей программы.
Функциональные тесты могут ходить в БД, во внешние сервисы.
Некоторые внешние ресурсы могут быть заменены на заглушки.
Для удобства функционального тестирования требуется изолированная среда.

Пишутся разработчиками, и могут писаться qa-инженерами.
Запускаются автоматически во время сборки. Содержатся в кодовой базе проекта.

PS: Такое название можно встретить только в "KolesaGroup", на просторах интернета
люди называют функциональными - все виды тестов, которые проверяют функциональность
приложения. А нефункциональными - стресс тесты, тесты производительности, безопасности и т.д.

Пример:
- messages (https://stash.kolesa-team.org/projects/GL/repos/messages/browse/app/usecase/repository/mysql/message_test.go)
- wallet (https://stash.kolesa-team.org/projects/WAL/repos/wallet/browse/pkg/repository/payment_log/payment_log_test.go)
- wallet (https://stash.kolesa-team.org/projects/WAL/repos/wallet/browse/pkg/repository/user_balance/user_balance_test.go)

### Integration тесты
Тесты, которые проверяют интеграцию между двумя и более модулями, сервисами,
в окружении приближенном к боевому.
Обычно тестируют внешние API сервисов и их работу друг с другом.

Чаще всего пишутся qa-инженерами, но никто не запрещает писать разработчикам.
Могут находиться в другом проекте, отличном от цели тестирования.
Также могут быть написаны с использованием другого языка.

Пример:
- ms-autotest (https://stash.kolesa-team.org/projects/MC/repos/ms-autotest/browse)
- ms-krpro (https://stash.kolesa-team.org/projects/GL/repos/ms-krpro/browse/tests/k6)

### Golden тесты (только в go)
Тесты, которые работают на основе слепков возвращаемых данных.
Слепки подготавливаются и проверяются специальной библиотекой.
А правильность слепков определяется с помощью фиксации состояния,
которое разработчик считает валдиным на данном этапе работы программы.

Документация: https://github.com/sebdah/goldie

Пример:
- avtoelon-pro (https://stash.kolesa-team.org/projects/MSKLUZ/repos/avtoelon-pro/browse/entity/rule_test.go)
- ms-krpro (https://stash.kolesa-team.org/projects/GL/repos/ms-krpro/browse/internal/input/request_test.go)

### Example тесты (только в go)
Специальный вид тестов, который встроен в go. Полезен, когда необходимо
проверить сценарий или последовательность действий. Такие тесты управляются
с помощью специального слова "Output", и проверяют stdout на наличие указанных
сообщений. Можно проверять, как четкую последовательность действий, так и просто
наличие всех строк в любом порядке с помощью ключевых слов "Unordered output"

Документация: https://pkg.go.dev/testing#hdr-Examples

Пример:
- autodial (https://stash.kolesa-team.org/projects/MI/repos/autodial/browse/_example/call_outgoing_test.go)

# Конкурентность

### Go рутины
Основной инструмент для создания конкурентной среды в языке go

Особенности:
- требуют контроля
- могут утекать
- работают даже на одном ядре

Полезные ссылки:
- https://tour.golang.org/concurrency/1
- https://gobyexample.com/goroutines
- https://golang.org/ref/spec#Go_statements

### Каналы
Средство коммуникации между частями программы.
Каналы могут иметь любой тип. От типа зависят значения, которые можно
записать в канал и прочитать из него.

Особенности:
- могут быть простыми
- могут быть буфферизированными
- можно разрешить только писать в канал или только читать из канала

Полезные ссылки:
- https://tour.golang.org/concurrency/2
- https://gobyexample.com/channels
- https://golang.org/ref/spec#Channel_types

### Пакеты sync и sync/atomic
Основной пакет для синхронизации в программах на go.
Содержит набор инструментов, которые позволяют делать
операции безопасными в конкурентной среде.

Особенности:
- Mutex, RWMutex - захватчик ресурсов, локер
- WaitGroup - ждун
- Once - контроллер одноразового выполнения
- Map - потоко-безопасная версия map
- atomic - пакет с инструментами для проведения типичных атомарных операций

Ползные ссылки:
- https://tour.golang.org/concurrency/9
- https://gobyexample.com/mutexes
- https://gobyexample.com/waitgroups
- https://gobyexample.com/atomic-counters