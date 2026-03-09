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