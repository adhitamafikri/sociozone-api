CREATE TABLE users_posts_likes (
    post_id int NOT NULL,
    user_id int NOT NULL,

    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL,

    CONSTRAINT post_id_FK FOREIGN KEY (post_id) REFERENCES posts (id),
    CONSTRAINT user_id_FK FOREIGN KEY (user_id) REFERENCES users (id)
)