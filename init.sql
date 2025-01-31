-- Digkala
INSERT INTO company
VALUES (1,
        'Digikala',
        'DK',
        'Behtarin foroshgah mojod dar SanAndreas.',
        0,
        1.0);
INSERT INTO company_office
VALUES (DEFAULT,
        1,
        2183.459717,
        -2254.026855,
        2183.459717,
        -2254.026855,
        14.771435);

-- Digkala Delivery
INSERT INTO company_job
VALUES (DEFAULT,
        1,
        1,
        1);
INSERT INTO company_job_checkpoint
VALUES (DEFAULT,
        1,
        1,
        1,
        2181.576660,
        -2302.060547,
        13.546875);

-- Snapp
INSERT INTO company
VALUES (2,
        'Snapp',
        'SNP',
        'Behtarin taxi online dar SanAndreas.',
        0,
        1.0);
INSERT INTO company_office
VALUES (DEFAULT,
        2,
        1500.237793,
        -1053.283936,
        1500.237793,
        -1053.283936,
        25.062500);
INSERT INTO company_job
VALUES (DEFAULT,
        2,
        1,
        1);
