CREATE TABLE IF NOT EXISTS article (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL
);


INSERT INTO article (title, content) VALUES ('Hello World', 'This is the first article.');
INSERT INTO article (title, content) VALUES ('Hello Docker', 'This is the second article.');
INSERT INTO article (title, content) VALUES ('Hello Home Server', 'This is the third article.');
