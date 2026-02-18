INSERT INTO roles (
    name
) VALUES (
    'Админ'
),
(
    'Менеджер'
),
(
    'Сотрудник'
);

INSERT INTO users (
    login,
    password_hash,
    name,
    surname,
    patronymic,
    role_id
) VALUES (
    'admin',
    '$2a$10$w3O3KyAxRBtYdpcSTieNUOUkiPxZ3GXOcoVAElBBZ.S4Ce.WdOmyC',
    'Максим',
    'Валявкин',
    'Александрович',
    1
), (
    'manager',
    '$2a$10$iLtvP.pOGdfDPlx6IrOc3eHaSMC0vmos.iOyikYJfahu2ffZeVErW',
    'Иван',
    'Сараканайс',
    'Юрьевич',
    2
), (
    'worker',
    '$2a$10$Pzlsgbd.edQOAyJ/5LKwAuk.8jxE2pCs7JQrn5SWPDlAmhO3Sm6Uq',
    'Араик',
    'Григорянц',
    'Арменакович',
    3
) ON CONFLICT (login) DO NOTHING;

INSERT INTO units (
    code,
    sign,
    name
) VALUES (
    796,
    'шт',
    'Штука'
), (
    778,
    'упак',
    'Упаковка'
), (
    717,
    'кор',
    'Коробка'
), (
    166,
    'кг',
    'Килограмм'
), (
    163,
    'г',
    'Грамм'
), (
    112,
    'л',
    'Литр'
), (
    111,
    'мл',
    'Милилитр'
);