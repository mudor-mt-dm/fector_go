CREATE TABLE authors (
                         id SERIAL PRIMARY KEY,
                         name VARCHAR(255) NOT NULL
);

CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       short_description VARCHAR(200),
                       full_description TEXT
);

CREATE TABLE book_authors (
                              book_id INT NOT NULL,
                              author_id INT NOT NULL,
                              PRIMARY KEY (book_id, author_id),
                              FOREIGN KEY (book_id) REFERENCES books(id),
                              FOREIGN KEY (author_id) REFERENCES authors(id)
);
