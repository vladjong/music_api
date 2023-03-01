# WAREHOUSE_API

<details>
<summary>1.Вопросы для разогрева</summary>

### 1. Опишите самую интересную задачу в программировании, которую вам приходилось решать?

Мне необходимо было реализовать новый алгоритм расчета оптимального пути доставки между двумя точками на больших дааных.
Задача была поставлена бизнессом, чтобы новый алгоритм увеличил метрики доставки. Выбор алгоритма полностью ложился на мои плечи. Главная сложность была в анализе реальных данных, так-как это был не учебный проект. В команде я работал с аналитиками, которые помогали с тестированием алгоритма. В результате нам удалось увеличить показатели оптимальной доставки и обнаружены пути, которые ухудшают показатели. На данный момент, этот сервис внедрен в витрину доствки в Willdberis

### 2. Расскажите о своем самом большом факапе? Что вы предприняли для решения проблемы?

Это произошло на хакатоне от Школы 21 по разработке telegram бота.

Разработка бота проводилась в три этапа:
- Аналитика
- Разработка
- Презентация

Я собрал команду из знакомых ребят. В команде я был тимлидом.
Поначалу все шло хорошо, но на этапе разработке все пошло не по плану. Ребята расслабились и весь код я писал в одного, и в один момент я запутался в логике и проект застыл.
Это была моя ошибка писать код в одного, а не поговорить с ребятами и обсудить наши проблемы.
В результате мы собрались и обсудили обстановку в команде, после этого мы начали работать вместе. Мы успешно защитили проект и заняли второе место

### 3. Каковы ваши ожидания от участия в буткемпе?

1. Прокачать свои hard скиллы и выйти на новый уровень!
2. Познакомиться с компетентным менторами и перенять их лучшие качества
3. Окунуться в корпоративную культуру cloud

</details>

<details>
<summary>2.Разработка музыкального плейлиста</summary>

### `timer` - реализация таймера для плейлиста

### `async_list` - реализация асинхронного двухсвязного списка

### `playlist` - модуль для работы с плейлистом

</details>

<details>

<summary>3.Построение API для музыкального плейлиста</summary>

## Инструкция

1. Склонировать репозиторий
```
git clone https://github.com/vladjong/music_api.git
```

2. Добавить `.env` файл в проект

3. Запустить проекта через docker compose
```
make docker
```
4. Пример работы

[Видео примера работы сервиса на youtube](https://youtu.be/_dzC4VLlupk)

- [Добавить несколько песен в плейлист](#добавить-песню-в-плейлист)
- [Начать воспроизвидение плейлиста](#начать-воспроизведение-плейлиста)
- [Узнать какая песня играет сейчас](#текущая-песня-в-плейлисте)
- [Остановить плейлист](#остановить-воспроизведение-плейлиста)
- [Включить следующую песню](#следующая-песня-в-плейлисте)
- [Узнать какая песня играет сейчас](#текущая-песня-в-плейлисте)
- [Включить предыдущую песню](#предыдущая-песня-в-плейлисте)
- [Узнать какая песня играет сейчас](#текущая-песня-в-плейлисте)

5. Завершить проект
```
make clean
```
6. Запустить тесты
```
make test
```

## Методы API

### Добавить песню в плейлист

<details>
<summary>пример запроса:</summary>

```
curl --location 'http://0.0.0.0:8080/api/v1/playlist' \
--header 'X-Forwarded-For: 123.0.0.12' \
--header 'Content-Type: application/json' \
--data '{
    "name": "test_1",
    "duration": 5
}'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "status": "Song add in playlist"
}
```
</details>

### Посмотреть все песни в плейлисте

<details>
<summary>пример запроса:</summary>

```
curl --location 'http://0.0.0.0:8080/api/v1/playlist/song'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
[
    {
        "id": 1,
        "name": "test_1",
        "duration": 5
    },
    {
        "id": 2,
        "name": "test_2",
        "duration": 5
    },
    {
        "id": 3,
        "name": "test_3",
        "duration": 5
    },
    {
        "id": 4,
        "name": "test_4",
        "duration": 5
    }
]
```
</details>

### Посмотреть песню в плейлисте по id

<details>
<summary>пример запроса:</summary>

```
curl --location 'http://0.0.0.0:8080/api/v1/playlist/song/1'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "id": 1,
    "name": "test_1",
    "duration": 5
}
```
</details>

### Обновить песню в плейлисте по id

<details>
<summary>пример запроса:</summary>

```
curl --location --request PUT 'http://0.0.0.0:8080/api/v1/playlist/song' \
--header 'Content-Type: application/json' \
--data '{
    "id": 1,
    "name": "test_new",
    "duration": 25
}'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "status": "Song id=1 update in playlist"
}
```
</details>

### Удалить песню в плейлисте по id

<details>
<summary>пример запроса:</summary>

```
curl --location --request DELETE 'http://0.0.0.0:8080/api/v1/playlist/song/1'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "status": "Song id=1 delete in playlist"
}
```
</details>


### Начать воспроизведение плейлиста

<details>
<summary>пример запроса:</summary>

```
curl --location 'http://0.0.0.0:8080/api/v1/playlist/play'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "status": "Play apply"
}
```
</details>

### Остановить воспроизведение плейлиста

<details>
<summary>пример запроса:</summary>

```
curl --location 'http://0.0.0.0:8080/api/v1/playlist/stop'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "status": "Stop apply"
}
```
</details>


### Следующая песня в плейлисте

<details>
<summary>пример запроса:</summary>

```
curl --location 'http://0.0.0.0:8080/api/v1/playlist/next'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "status": "Next apply"
}
```
</details>

### Предыдущая песня в плейлисте

<details>
<summary>пример запроса:</summary>

```
curl --location 'http://0.0.0.0:8080/api/v1/playlist/prev'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "status": "Prev apply"
}
```
</details>

### Текущая песня в плейлисте

<details>
<summary>пример запроса:</summary>

```
curl --location 'http://0.0.0.0:8080/api/v1/playlist/play_song'
```
</details>

<details>
<summary>пример ответа:</summary>

```json
{
    "Id": 2,
    "Name": "test_2",
    "Duration": 5
}
```
</details>

</details>
