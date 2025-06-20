CREATE TABLE variant_assignments (
    id SERIAL PRIMARY KEY,
    user_identifier TEXT NOT NULL,
    experiment_id INT NOT NULL,
    variant_id INT NOT NULL,
    assigned_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    UNIQUE (user_identifier, experiment_id),

    FOREIGN KEY (experiment_id) REFERENCES experiments(id) ON DELETE CASCADE,
    FOREIGN KEY (variant_id) REFERENCES variants(id) ON DELETE CASCADE
);