CREATE TABLE songs (
    id SERIAL PRIMARY KEY,
    release_date DATE,
    group varchar(255),
    song varchar(255),
     text_parts TEXT,
    link TEXT
);
