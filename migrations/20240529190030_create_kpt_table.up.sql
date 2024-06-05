CREATE TABLE kpt (
    id SERIAL PRIMARY KEY,
    cad_quarter VARCHAR(20) UNIQUE NOT NULL,
    date_formation TIMESTAMP NOT NULL,
    file_name VARCHAR(255) NOT NULL,
    size INT NOT NULL,
    content_type VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
)
