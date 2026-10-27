CREATE TABLE tasks (
    id BIGINT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title TEXT NOT NULL,
    description TEXT NOT NULL,
    task_type TEXT NOT NULL,
    difficulty_level TEXT NOT NULL,
    deadline TIMESTAMPTZ,
    completed BOOLEAN NOT NULL DEFAULT FALSE
);