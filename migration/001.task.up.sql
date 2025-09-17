CREATE TABLE tasks
(
    id          SERIAL PRIMARY KEY,
    title       TEXT NOT NULL,
    description TEXT,
    completed   BOOLEAN     DEFAULT FALSE,
    priority    VARCHAR(10) DEFAULT 'medium' CHECK (priority IN ('low', 'medium', 'high')),
    due_date    TIMESTAMP,
    created_at  TIMESTAMP   DEFAULT NOW(),
    updated_at  TIMESTAMP   DEFAULT NOW()
);
