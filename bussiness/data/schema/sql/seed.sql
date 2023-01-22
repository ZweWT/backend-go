INSERT INTO users (user_id, name, email, roles, password_hash, date_created, date_updated) VALUES
	('5cf37266-3473-4006-984f-9325122678b7', 'Mg Admin', 'admin@example.com', '{ADMIN}', '$2a$10$cfcK/1sekPiJVMlg3mf06eUQCmPbJKRIEsl5Wo4fs8v.UNjO.4opG', '2019-03-24 00:00:00', '2019-03-24 00:00:00'),
	('45b5fbd3-755f-4379-8f07-a58d4a30fa2f', 'Mg User', 'user@example.com', '{USER}', '$2a$10$cfcK/1sekPiJVMlg3mf06eUQCmPbJKRIEsl5Wo4fs8v.UNjO.4opG', '2019-03-24 00:00:00', '2019-03-24 00:00:00')
	ON CONFLICT DO NOTHING;

INSERT INTO categories(name) VALUES
	('Shoe'),
	('Shirt')
	ON CONFLICT DO NOTHING;

INSERT INTO products (product_id, category_id, name, cost, date_created, date_updated) VALUES
	('a2b0639f-2cc6-44b8-b97b-15d69dbb511e', 1, 'Couple Shoe', 3000, '2019-01-01 00:00:01.000001+00', '2019-01-01 00:00:01.000001+00'),
	('72f8b983-3eb4-48db-9ed0-e45cc6bd716b', 2, 'Couple Shirt', 5000, '2019-01-01 00:00:02.000001+00', '2019-01-01 00:00:02.000001+00')
	ON CONFLICT DO NOTHING;
