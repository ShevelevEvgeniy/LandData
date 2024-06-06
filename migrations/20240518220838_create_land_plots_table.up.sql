CREATE TABLE land_plots (
                            id SERIAL PRIMARY KEY,
                            cad_number VARCHAR(20) UNIQUE NOT NULL,
                            coordinates GEOMETRY(POLYGON, 4326) NOT NULL,
                            category TEXT NOT NULL,
                            permitted_use TEXT NOT NULL,
                            area DECIMAL(10, 2) NOT NULL,
                            okato VARCHAR(20) NOT NULL,
                            kladr VARCHAR(20) NOT NULL,
                            readable_address TEXT NOT NULL,
                            created_at TIMESTAMP NOT NULL DEFAULT NOW(),
                            updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE INDEX land_plots_cad_number_idx ON land_plots USING hash (cad_number);