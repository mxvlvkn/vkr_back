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
ON U.role_id = R.id
