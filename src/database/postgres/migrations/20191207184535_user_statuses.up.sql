CREATE TABLE user_statuses (
  id smallserial NOT NULL,

  name varchar(50) NOT NULL,

  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,

  PRIMARY KEY (id)
);
