CREATE TABLE variants (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    weight FLOAT NOT NULL DEFAULT 1.0,
    payload JSONB NOT NULL DEFAULT '{}'::jsonb,
    experiment_id INTEGER NOT NULL REFERENCES experiments(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
