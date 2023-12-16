CREATE TABLE IF NOT EXISTS templates (
    code varchar(255) PRIMARY KEY NOT NULL,
    title text NOT NULL,
    body text NOT NULL
);