-- +goose Up
DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE labels (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    opened TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    closed TIMESTAMP,
    author_id INTEGER REFERENCES users(id) DEFAULT 0,
    assigned_id INTEGER REFERENCES users(id) DEFAULT 0,
    title VARCHAR(255),
    content TEXT
);

CREATE TABLE tasks_labels(
    task_id INTEGER REFERENCES tasks(id),
    label_id INTEGER REFERENCES labels(id)
);


-- +goose Down
DROP TABLE tasks_labels;
DROP TABLE tasks;
DROP TABLE labels;
DROP TABLE users;
