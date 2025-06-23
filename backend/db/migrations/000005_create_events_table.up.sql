CREATE TABLE events (
    id SERIAL PRIMARY KEY,
    experiment_id INTEGER NOT NULL REFERENCES experiments(id) ON DELETE CASCADE,
    variant_id INTEGER NOT NULL REFERENCES variants(id) ON DELETE CASCADE,
    user_identifier TEXT NOT NULL,
    event_type TEXT NOT NULL CHECK (event_type IN ('view', 'conversion')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
)