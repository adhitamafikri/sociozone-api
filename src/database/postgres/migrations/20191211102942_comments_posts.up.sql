CREATE TABLE comments_posts (
    post_id int NOT NULL,
    comment_id int NOT NULL,

    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL,

    CONSTRAINT post_id_FK FOREIGN KEY (post_id) REFERENCES posts (id),
    CONSTRAINT comment_id_FK FOREIGN KEY (comment_id) REFERENCES comments (id)
)