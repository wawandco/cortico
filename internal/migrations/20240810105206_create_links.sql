CREATE TABLE links (
    id uuid PRIMARY KEY,
    original TEXT NOT NULL,
    short TEXT NOT NULL,
    created_at DATE NOT NULL
);