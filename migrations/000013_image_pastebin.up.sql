CREATE TABLE Image_Pastebin (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    url_slug TEXT NOT NULL UNIQUE,
    file_size INTEGER,
    mime_type VARCHAR(50)
);