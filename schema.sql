CREATE TABLE IF NOT EXISTS player
(
    id          SERIAL PRIMARY KEY,
    username    VARCHAR(24) UNIQUE NOT NULL,
    password    VARCHAR(255)       NOT NULL,
    money       INT                NOT NULL DEFAULT 1000,
    level       INT                NOT NULL DEFAULT 1,
    exp         INT                NOT NULL DEFAULT 0,
    gold        INT                NOT NULL DEFAULT 0,
    token       INT                NOT NULL DEFAULT 0,
    hour        INT                NOT NULL DEFAULT 0,
    minute      INT                NOT NULL DEFAULT 0,
    second      INT                NOT NULL DEFAULT 0,
    vip         INT                NOT NULL DEFAULT 0,
    helper      INT                NOT NULL DEFAULT 0,
    is_online   BOOLEAN            NOT NULL DEFAULT false,
    kills       INT                NOT NULL DEFAULT 0,
    deaths      INT                NOT NULL DEFAULT 0,
    pos_x       FLOAT4             NOT NULL DEFAULT 1682.0309,
    pos_y       FLOAT4             NOT NULL DEFAULT 1448.0178,
    pos_z       FLOAT4             NOT NULL DEFAULT 10.7724,
    pos_angle   FLOAT4             NOT NULL DEFAULT 270.0725,
    language    INT                NOT NULL DEFAULT 0,
    last_login  TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    last_played TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    created_at  TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT _username_index UNIQUE (username)
);
CREATE TABLE IF NOT EXISTS skin
(
    id        SERIAL PRIMARY KEY,
    player_id INT  NOT NULL,
    skin_id   INT  NOT NULL,
    active    BOOL NOT NULL,
    FOREIGN KEY (player_id) REFERENCES player (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (player_id, skin_id),
    UNIQUE (player_id, active)
);
CREATE TABLE IF NOT EXISTS ban
(
    id        SERIAL PRIMARY KEY,
    admin_id  INT         NULL,
    username  VARCHAR(24) NOT NULL,
    ip        VARCHAR(15) NOT NULL,
    reason    VARCHAR(128),
    banned_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expire_at TIMESTAMP,
    FOREIGN KEY (admin_id) REFERENCES player (id) ON DELETE SET NULL ON UPDATE SET NULL
);
CREATE TABLE IF NOT EXISTS company
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(36) NOT NULL,
    tag         VARCHAR(3)  NOT NULL,
    description VARCHAR(80),
    balance     INT    DEFAULT 0,
    multiplier  float4 DEFAULT 1.0,
    UNIQUE (name),
    UNIQUE (tag)
);
CREATE TABLE IF NOT EXISTS company_member
(
    id         SERIAL PRIMARY KEY,
    player_id  INT  NOT NULL,
    company_id INT  NOT NULL,
    role       int2 NOT NULL DEFAULT 0,
    FOREIGN KEY (player_id) REFERENCES player (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (player_id)
);
CREATE TABLE IF NOT EXISTS company_member_info
(
    id         SERIAL PRIMARY KEY,
    player_id  INT  NOT NULL,
    company_id INT  NOT NULL,
    hour       INT  NOT NULL DEFAULT 0,
    minute     int2 NOT NULL DEFAULT 0,
    second     int2 NOT NULL DEFAULT 0,
    score      INT  not NULL DEFAULT 0,
    level      INT  not NULL DEFAULT 0,
    FOREIGN KEY (player_id) REFERENCES player (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (player_id, company_id)
);
CREATE TABLE IF NOT EXISTS company_application
(
    id          SERIAL PRIMARY KEY,
    player_id   INT                 NOT NULL,
    company_id  INT                 NOT NULL,
    description VARCHAR(80)         NULL,
    accepted    int2      DEFAULT 0 NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    expired_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP + INTERVAL '1' DAY,
    answer      VARCHAR(20)         NULL,
    answered_at TIMESTAMP           NULL,
    FOREIGN KEY (player_id) REFERENCES player (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE ON UPDATE CASCADE
);
