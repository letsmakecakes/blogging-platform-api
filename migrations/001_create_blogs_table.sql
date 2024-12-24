CREATE TABLE IF NOT EXISTS blogs (
                                     id SERIAL PRIMARY KEY,
                                     title VARCHAR(255) NOT NULL,
                                     content TEXT NOT NULL,
                                     category VARCHAR(100) NOT NULL,
                                     tags TEXT[] DEFAULT '{}', -- Use an array for tags for better structure
                                     created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                     updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                     CONSTRAINT chk_title_length CHECK (char_length(title) > 0), -- Ensures non-empty titles
                                     CONSTRAINT chk_category_length CHECK (char_length(category) > 0) -- Ensures non-empty categories
);