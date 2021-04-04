# veebiprogrammeerine

## Requirements & ROUTES

### Fetch users - GET /api/user

Returns list of users, currently only 1 user created

### Start session - POST /session

As a user, I want to be able to create a session

Payload:
```json
{
    "status": "started",
    "user_id": "<used id>"
}
```

### Pause/Stop session - PATCH /session/:id

As a user, I want to be able to save my time tracking session when I am done with it

Payload:

```json
{
    "status": "<ended | paused>",
    "user_id": "<used id>"
}
```

### See sessions - GET /session

As a user, I want to see my latest sessions

## DB SQLite

### Table `users`

| Name             | Type   |
| -------          | ------ |
| id               | string |
| name             | string |

### Table `sessions`

| Name             | Type      |
| ----------       | ------    |
| id               | string    |
| user_id          | string    |
| status           | string    |
| started_at       | Timestamp |
| stopper_at       | Timestamp |
| ended_at         | Timestamp |
| seconds          | float     |

## Run with docker 

```sh
docker-compose up --build -d
```
