# veebiprogrammeerine

## Requirements


POST /session Payload optional name
As a user, I want to be able to start a time tracking session

PUT /session?session_group_id=1
As a user, I want to be able to save my time tracking session when I am done with it

GET /session?timeframe=day
As a user, I want an overview of my sessions for the day, week and month

PUT /session/{id}?stop=true
As a user, I want to be able to stop a time tracking session

PUT /session/{id} Payload name
As a user, I want to be able to name my time tracking session

## DB

### Table User

| Name    | Type   |
| ------- | ------ |
| user_id | string |

### Table Session

| Name             | Type   |
| ----------       | ------ |
| session_id       | string |
| session_group_id | string |
| status           | string |
| start            | int    |
| end              | int    |
| user_id          | string |

## REST paths needed

## OPEN API specs

https://mermade.github.io/openapi-gui/
