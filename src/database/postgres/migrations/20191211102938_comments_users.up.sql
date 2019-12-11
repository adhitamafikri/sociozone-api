CREATE TABLE comments_users (
    user_id int NOT NULL,
    comment_id int NOT NULL,

    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL,

    CONSTRAINT user_id_FK FOREIGN KEY (user_id) REFERENCES users (id),
    CONSTRAINT comment_id_FK FOREIGN KEY (comment_id) REFERENCES comments (id)
)