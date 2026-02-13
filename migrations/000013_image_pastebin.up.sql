CREATE TABLE Image_Pastebin (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    url_slug TEXT NOT NULL UNIQUE,
    image_data BYTEA DEFAULT NULL,
    file_size INTEGER,
    mime_type VARCHAR(25)
);