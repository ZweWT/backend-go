-- Version: 1.1
-- Description: Create table users
CREATE TABLE users (
	user_id       UUID,
	name          TEXT,
	email         TEXT UNIQUE,
	roles         TEXT[],
	password_hash TEXT,
	date_created  TIMESTAMP,
	date_updated  TIMESTAMP,

	PRIMARY KEY (user_id)
);

-- Version: 1.2
-- Description: Create table categories
CREATE TABLE categories (
	category_id   SERIAL,
	name         TEXT,

	PRIMARY KEY (category_id)
);

-- Version: 1.3
-- Description: Create table products
CREATE TABLE products (
	product_id   UUID,
	name         TEXT,
	cost         INT,
	category_id   INT,
	date_created TIMESTAMP,
	date_updated TIMESTAMP,

	PRIMARY KEY (product_id),
	FOREIGN KEY (category_id) REFERENCES categories(category_id) ON DELETE CASCADE
);
