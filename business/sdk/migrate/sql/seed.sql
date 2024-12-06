INSERT INTO medication (id, name, form, dosage) VALUES
	('5cf37266-3473-4006-984f-9325122678b7', 'magic pill', 'tablet', 1)
ON CONFLICT DO NOTHING;
