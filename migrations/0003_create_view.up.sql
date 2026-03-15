CREATE OR REPLACE VIEW users_view AS
SELECT 
	U.id,
	U.login,
	U.name,
	U.surname,
	U.patronymic,
	R.name AS role
FROM users U
LEFT JOIN roles R
ON U.role_id = R.id;

CREATE OR REPLACE VIEW numenclatures_view AS
SELECT 
	N.id,
	N.name,
	N.image_url,
	CASE 
        WHEN N.use_serial THEN 'Да'
        ELSE 'Нет'
    END AS use_serial,
    
    CASE 
        WHEN N.use_marks THEN 'Да'
        ELSE 'Нет'
    END AS use_marks,
	N.article,
	U.name AS unit,
	M.name AS manufacturer
FROM numenclatures N
LEFT JOIN units U
ON N.unit_id = U.id
LEFT JOIN manufacturers M
ON N.manufacturer_id = M.id;

CREATE OR REPLACE VIEW barcodes_view AS
SELECT 
	B.id,
	B.code,
	N.name AS numenclature,
	N.id AS numenclature_id
FROM barcodes B
LEFT JOIN numenclatures N
ON B.numenclature_id = N.id;

CREATE OR REPLACE VIEW marks_view AS
SELECT 
	M.id,
	M.code,
	N.name AS numenclature,
	N.id AS numenclature_id
FROM marks M
LEFT JOIN numenclatures N
ON M.numenclature_id = N.id;