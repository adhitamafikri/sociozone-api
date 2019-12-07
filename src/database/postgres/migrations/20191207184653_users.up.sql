CREATE TABLE users (
  id serial NOT NULL,

  name varchar(150) NOT NULL,
  email varchar(255) UNIQUE NOT NULL,
  uername varchar(50) UNIQUE NOT NULL,
  password varchar(50) NOT NULL,
  profile_picture varchar(255),
  bio varchar(255),
  website varchar(150),

  -- created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  -- updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  -- PostgreSQL doesnt't have the feature like above
  -- we should create a trigger instead
  created_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  deleted_at timestamp NULL DEFAULT NULL,

  PRIMARY KEY (id)
);
