DROP TABLE IF EXISTS videos;

CREATE TABLE videos (
                        id SERIAL PRIMARY KEY,
                        url VARCHAR(255) NOT NULL,
                        title VARCHAR(255) NOT NULL,
                        description VARCHAR (255),
                        last_update DATE
);