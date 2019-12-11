CREATE TABLE comments (
    id serial NOT NULL,

    body varchar(255) NOT NULL,

    created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
    deleted_at timestamp NULL DEFAULT NULL,

    PRIMARY KEY (id)

    -- TODO
    -- Create a trigger to INSERT into comments_user AND comments_posts after insert
)