USE supermarket;

-- products table
INSERT INTO products (id, name, price) VALUES
    ('2f0495b9-099e-4c3f-9803-a4b8e32448a5', 'Onion', 3.50),
    ('671641b5-8d8d-43d3-9b44-4d8addeb5108', 'Patato', 3.99),
    ('c7eee60e-4066-4029-abfd-2d4ff6bc0ecc', 'Eggs', 1.50),
    ('e087d9f4-f377-432c-a375-e37b200a8fcf', 'Cheese', 9.99),
    ('a6e0e62b-b3ff-4ef5-887d-29e41e0a5d53', 'Tomato', 3.50),
    ('25346fba-5f97-4793-b479-b2047d823016', 'Sprite 1L', 5.50),
    ('9a9a2f60-555b-461d-9ec6-0422a08348d4', 'Coca Cola 1L', 5.25),
    ('d7e9c188-a9da-4461-a061-3bf21d89b012', 'Pepsi 1L', 5.75)
;

-- orders table
INSERT INTO orders (id) VALUES
    ('0268feac-8135-4d5f-9e8d-01bc61263eba'),
    ('2e77d060-8a5e-440b-a5c8-0d82aa8983d2')
;

INSERT INTO order_product (order_id, product_id, amount) VALUES
    -- only drinks
    ('0268feac-8135-4d5f-9e8d-01bc61263eba', '25346fba-5f97-4793-b479-b2047d823016', 3),
    ('0268feac-8135-4d5f-9e8d-01bc61263eba', '9a9a2f60-555b-461d-9ec6-0422a08348d4', 3),
    ('0268feac-8135-4d5f-9e8d-01bc61263eba', 'd7e9c188-a9da-4461-a061-3bf21d89b012', 5),

    -- only breakfast
    ('2e77d060-8a5e-440b-a5c8-0d82aa8983d2', 'a6e0e62b-b3ff-4ef5-887d-29e41e0a5d53', 1),
    ('2e77d060-8a5e-440b-a5c8-0d82aa8983d2', 'c7eee60e-4066-4029-abfd-2d4ff6bc0ecc', 1),
    ('2e77d060-8a5e-440b-a5c8-0d82aa8983d2', 'e087d9f4-f377-432c-a375-e37b200a8fcf', 1)
;