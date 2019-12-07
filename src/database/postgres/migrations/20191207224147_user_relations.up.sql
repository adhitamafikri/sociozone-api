CREATE TABLE user_relations (
    followee_id int NOT NULL,
    follower_id int NOT NULL,

    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL,

    CONSTRAINT user_id_FK_1 FOREIGN KEY (followee_id) REFERENCES users (id),
    CONSTRAINT user_id_FK_2 FOREIGN KEY (follower_id) REFERENCES users (id)
)