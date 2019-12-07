CREATE TABLE posts (
    id serial NOT NULL,
    user_id int NOT NULL,

    title varchar(50) NOT NULL,
    body varchar(255) NOT NULL,
    pictures json NOT NULL,

    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL,

    PRIMARY KEY (id),

    CONSTRAINT user_id_FK FOREIGN KEY (user_id) REFERENCES users (id)
)