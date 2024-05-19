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

-- Добавление данных в таблицу users
INSERT INTO users (name) VALUES
('Alice'),
('Bob'),
('Charlie');

-- Добавление данных в таблицу labels
INSERT INTO labels (name) VALUES
('Urgent'),
('Bug'),
('Feature');

-- Добавление данных в таблицу tasks
INSERT INTO tasks (opened, closed, author_id, assigned_id, title, content) VALUES
(CURRENT_TIMESTAMP, NULL, 1, 2, 'Fix login issue', 'User unable to login with correct credentials'),
(CURRENT_TIMESTAMP, NULL, 2, 3, 'Add user profile page', 'Create a new page for user profiles'),
(CURRENT_TIMESTAMP, NULL, 3, 1, 'Update password policy', 'Password must be at least 12 characters long');

-- Добавление связей между tasks и labels
INSERT INTO tasks_labels (task_id, label_id) VALUES
(1, 2), -- 'Fix login issue' has 'Bug' label
(2, 3), -- 'Add user profile page' has 'Feature' label
(3, 1); -- 'Update password policy' has 'Urgent' label

-- +goose Down
DROP TABLE tasks_labels;
DROP TABLE tasks;
DROP TABLE labels;
DROP TABLE users;
