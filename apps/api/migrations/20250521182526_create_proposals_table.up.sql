CREATE TABLE proposals (
    id SERIAL PRIMARY KEY,
    process_id INT NOT NULL,
    user_id INT,
    title VARCHAR(255) NOT NULL,
    body TEXT NOT NULL,
    votes_count INT DEFAULT 0,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE proposals
ADD CONSTRAINT fk_process
FOREIGN KEY (process_id) REFERENCES processes(id) ON DELETE CASCADE;


ALTER TABLE proposals
ADD CONSTRAINT fk_user
FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;

CREATE INDEX idx_proposals_process_id ON proposals(process_id);
CREATE INDEX idx_proposals_user_id ON proposals(user_id);
