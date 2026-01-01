# CRUD-API-GO

A simple CRUD API for managing tasks, built with Go, Fiber, and PostgreSQL. Supports creating, reading, updating, and deleting tasks.

Author : [Kenean Dita](https://www.github.com/keneandita/)

## Features

- Create a new task
- List all tasks
- Get a task by ID
- Update a task
- Delete a task
- Auto-increment task IDs

## Database Setup

```sql
CREATE TABLE tasks (
    id BIGSERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT,
    status BOOLEAN NOT NULL DEFAULT false
);
````

Populate sample data:

```sql
INSERT INTO tasks (title, description, status) VALUES
('Learn Go', 'Learn Go basics and syntax', true),
('Fiber setup', 'Initialize Fiber project and routing', true),
('Postgres setup', 'Install PostgreSQL and create database', true);
```

---

## API Endpoints

| Method | Endpoint   | Description       |
| ------ | ---------- | ----------------- |
| POST   | /tasks     | Create a new task |
| GET    | /tasks     | List all tasks    |
| GET    | /tasks/:id | Get a task by ID  |
| PUT    | /tasks/:id | Update a task     |
| DELETE | /tasks/:id | Delete a task     |

---

## Example Request

- **Create Task**

```bash
POST /tasks
Content-Type: application/json

{
  "title": "Learn Fiber",
  "description": "Build a CRUD API with Fiber",
  "status": false
}
```

- **Response**

```json
{
  "message": "Task created successfully",
  "id": 1
}
```

## Running Locally

1. Install dependencies:

```bash
go mod tidy
```

2. Start PostgreSQL and create the `tasks` table.

3. Run the server:

```bash
go run .
```

Server runs at `http://localhost:3000`.

## Notes

* IDs are auto-incremented using `BIGSERIAL`.
* Status is a boolean (`true` = completed, `false` = pending).
* Modify Go code if you want UUID IDs instead of auto-increment.
