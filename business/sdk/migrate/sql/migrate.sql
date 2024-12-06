-- Version: 1.01
-- Description: Create table medication
CREATE TABLE medication (
	id     UUID,
	name   TEXT NOT NULL,
	form   TEXT NOT NULL,
	dosage int  NOT NULL,

	PRIMARY KEY (id)
);
