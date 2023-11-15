-- name: GetEmployees :one

SELECT * FROM employees WHERE employee_id = $1;

-- name: ListEmployees :many

SELECT * FROM employees ORDER BY name;

-- name: CreateEmployees :one

INSERT INTO
    employees(
        employee_id,
        user_id,
        name,
        position
    )
VALUES ($1, $2, $3, $4) RETURNING employee_id;

-- name: DeleteEmployees :exec

DELETE FROM employees WHERE employee_id = $1;

UPDATE employees
set
    user_id = $2,
    name = $3,
    position = $4
WHERE employee_id = $1;

-- users

-- name: Getusers :one

SELECT * FROM users WHERE user_id = $1;

-- name: Listusers :many

SELECT * FROM users ORDER BY username;

-- name: CreateUsers :one

INSERT INTO
    users (
        user_id,
        username,
        password
    )
VALUES
(
        $1,
        $2,
        $3
    ) RETURNING user_id;

-- name: DeleteUsers :exec

DELETE FROM users WHERE user_id = $1;

-- name: UpdateUsers :exec

UPDATE users
set
    username = $2,
    password = $3
WHERE user_id = $1;
