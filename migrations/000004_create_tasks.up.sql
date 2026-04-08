CREATE TABLE IF NOT EXISTS tasks (
    id          SERIAL PRIMARY KEY,
    project_id  INT NOT NULL REFERENCES projects(id) ON DELETE CASCADE,
    created_by  INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    title       VARCHAR(200) NOT NULL,
    description TEXT DEFAULT '',
    status      task_status DEFAULT 'todo',
    due_date    DATE,
    created_at  TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);