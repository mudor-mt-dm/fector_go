-- Insert authors
INSERT INTO authors (name) VALUES
                               ('J.K. Rowling'),
                               ('George R.R. Martin'),
                               ('Stephen King'),
                               ('Agatha Christie'),
                               ('Ernest Hemingway'),
                               ('Jane Austen'),
                               ('Mark Twain'),
                               ('F. Scott Fitzgerald'),
                               ('Harper Lee'),
                               ('Toni Morrison');

-- Insert books
INSERT INTO books (title, short_description, full_description) VALUES
                                                                   ('Harry Potter and the Philosopher''s Stone', 'The first book in the Harry Potter series.', 'The first novel in the Harry Potter series and Rowling''s debut novel.'),
                                                                   ('A Game of Thrones', 'The first book in A Song of Ice and Fire series.', 'The first novel in A Song of Ice and Fire, a series of fantasy novels by American author George R. R. Martin.'),
                                                                   ('The Shining', 'A horror novel by American author Stephen King.', 'The Shining is a horror novel by American author Stephen King.'),
                                                                   ('And Then There Were None', 'A mystery novel by Agatha Christie.', 'And Then There Were None is a mystery novel by English writer Agatha Christie.'),
                                                                   ('The Old Man and the Sea', 'A short novel written by Ernest Hemingway.', 'The Old Man and the Sea is a short novel written by the American author Ernest Hemingway.'),
                                                                   ('Pride and Prejudice', 'A romantic novel of manners by Jane Austen.', 'Pride and Prejudice is a romantic novel of manners written by Jane Austen.'),
                                                                   ('The Adventures of Huckleberry Finn', 'A novel by Mark Twain.', 'The Adventures of Huckleberry Finn is a novel by Mark Twain, first published in the United Kingdom in December 1884.'),
                                                                   ('The Great Gatsby', 'A novel by American writer F. Scott Fitzgerald.', 'The Great Gatsby is a 1925 novel by American writer F. Scott Fitzgerald.'),
                                                                   ('To Kill a Mockingbird', 'A novel by Harper Lee.', 'To Kill a Mockingbird is a novel by the American author Harper Lee.'),
                                                                   ('Beloved', 'A novel by American writer Toni Morrison.', 'Beloved is a 1987 novel by the American writer Toni Morrison.'),
                                                                   ('Harry Potter and the Chamber of Secrets', 'The second book in the Harry Potter series.', 'The second novel in the Harry Potter series and British author J. K. Rowling.'),
                                                                   ('A Clash of Kings', 'The second book in A Song of Ice and Fire series.', 'A Clash of Kings is the second novel in A Song of Ice and Fire, an epic fantasy series by American author George R. R. Martin.'),
                                                                   ('It', 'A horror novel by American author Stephen King.', 'It is a 1986 horror novel by American author Stephen King.'),
                                                                   ('Murder on the Orient Express', 'A detective novel by Agatha Christie.', 'Murder on the Orient Express is a detective novel by English writer Agatha Christie.'),
                                                                   ('For Whom the Bell Tolls', 'A novel by Ernest Hemingway.', 'For Whom the Bell Tolls is a novel by Ernest Hemingway published in 1940.'),
                                                                   ('Sense and Sensibility', 'A novel by Jane Austen.', 'Sense and Sensibility is a novel by Jane Austen, published in 1811.'),
                                                                   ('The Adventures of Tom Sawyer', 'A novel by Mark Twain.', 'The Adventures of Tom Sawyer is a novel by Mark Twain, first published in 1876.'),
                                                                   ('Tender Is the Night', 'A novel by F. Scott Fitzgerald.', 'Tender Is the Night is a novel by American writer F. Scott Fitzgerald.'),
                                                                   ('Go Set a Watchman', 'A novel by Harper Lee.', 'Go Set a Watchman is a novel by Harper Lee published in 2015.'),
                                                                   ('Song of Solomon', 'A novel by Toni Morrison.', 'Song of Solomon is a 1977 novel by American author Toni Morrison.');

-- Insert book_authors relationships
INSERT INTO book_authors (book_id, author_id) VALUES
                                                  (1, 1),
                                                  (2, 2),
                                                  (3, 3),
                                                  (4, 4),
                                                  (5, 5),
                                                  (6, 6),
                                                  (7, 7),
                                                  (8, 8),
                                                  (9, 9),
                                                  (10, 10),
                                                  (11, 1),
                                                  (12, 2),
                                                  (13, 3),
                                                  (14, 4),
                                                  (15, 5),
                                                  (16, 6),
                                                  (17, 7),
                                                  (18, 8),
                                                  (19, 9),
                                                  (20, 10),
                                                  (1, 2), -- Harry Potter and the Philosopher's Stone by J.K. Rowling and George R.R. Martin
                                                  (2, 3), -- A Game of Thrones by George R.R. Martin and Stephen King
                                                  (3, 4), -- The Shining by Stephen King and Agatha Christie
                                                  (4, 5), -- And Then There Were None by Agatha Christie and Ernest Hemingway
                                                  (5, 6), -- The Old Man and the Sea by Ernest Hemingway and Jane Austen
                                                  (6, 7), -- Pride and Prejudice by Jane Austen and Mark Twain
                                                  (7, 8), -- The Adventures of Huckleberry Finn by Mark Twain and F. Scott Fitzgerald
                                                  (8, 9), -- The Great Gatsby by F. Scott Fitzgerald and Harper Lee
                                                  (9, 10), -- To Kill a Mockingbird by Harper Lee and Toni Morrison
                                                  (10, 1); -- Beloved by Toni Morrison and J.K. Rowling
