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
    created_at  TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at  TIMESTAMP                   DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT _username_index UNIQUE (username)
);
CREATE TABLE IF NOT EXISTS player_job
(
    id        SERIAL PRIMARY KEY,
    player_id INT NOT NULL,
    job_id    INT NOT NULL,
    score     INT NOT NULL DEFAULT 0,
    FOREIGN KEY (player_id) REFERENCES player (id),
    UNIQUE (player_id, job_id)
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
    name        VARCHAR(36)        NOT NULL,
    tag         VARCHAR(3)         NOT NULL,
    description VARCHAR(80),
    balance     INT    DEFAULT 0   NOT NULL,
    multiplier  float4 DEFAULT 1.0 NOT NULL,
    UNIQUE (name),
    UNIQUE (tag)
);
CREATE TABLE IF NOT EXISTS company_office
(
    id         SERIAL PRIMARY KEY,
    company_id INT    NOT NULL UNIQUE,
    icon_x     float4 NOT NULL,
    icon_y     float4 NOT NULL,
    pickup_x   float4 NOT NULL,
    pickup_y   float4 NOT NULL,
    pickup_z   float4 NOT NULL,
    FOREIGN KEY (company_id) REFERENCES company (id)
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
    score      INT  not NULL DEFAULT 0,
    level      INT  not NULL DEFAULT 0,
    FOREIGN KEY (player_id) REFERENCES player (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (player_id, company_id)
);
CREATE TABLE IF NOT EXISTS company_application
(
    id          SERIAL PRIMARY KEY,
    player_id   INT                                                    NOT NULL,
    company_id  INT                                                    NOT NULL,
    description VARCHAR(80)                                            NULL,
    accepted    int2      DEFAULT 0                                    NOT NULL,
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP                    NOT NULL,
    expired_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP + INTERVAL '1' DAY NOT NULL,
    answer      VARCHAR(20)                                            NULL,
    answered_at TIMESTAMP                                              NULL,
    FOREIGN KEY (player_id) REFERENCES player (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE ON UPDATE CASCADE
);
CREATE TABLE IF NOT EXISTS company_job
(
    id         SERIAL PRIMARY KEY,
    company_id INT  NOT NULL,
    job_id     int2 NOT NULL,
    job_group  int2 NOT NULL,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (company_id, job_id)
);
CREATE TABLE IF NOT EXISTS company_job_checkpoint
(
    id         SERIAL PRIMARY KEY,
    company_id INT    NOT NULL,
    job_id     int2   NOT NULL,
    type       int2   NOT NULL,
    x          float4 NOT NULL,
    y          float4 NOT NULL,
    z          float4 NOT NULL,
    FOREIGN KEY (company_id) REFERENCES company (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS house
(
    id               SERIAL PRIMARY KEY,
    owner_id         INT                 NOT NULL,
    name             VARCHAR(100) UNIQUE NULL,
    price_multiplier float4              NOT NULL DEFAULT 1.0,
    pickup_x         float4              NOT NULL,
    pickup_y         float4              NOT NULL,
    pickup_z         float4              NOT NULL,
    entrance_x       float4              NOT NULL,
    entrance_y       float4              NOT NULL,
    entrance_z       float4              NOT NULL,
    garage           BOOLEAN             NOT NULL DEFAULT false,
    garage_slot      INT                 NULL,
    FOREIGN KEY (owner_id) REFERENCES player (id) ON DELETE SET NULL ON UPDATE SET NULL
);

CREATE TABLE IF NOT EXISTS vehicle
(
    id       SERIAL PRIMARY KEY,
    model    INT NOT NULL,
    owner_id INT NOT NULl,
    spoiler INT NULL,
    hood INT NULL,
    roof INT NULL,
    sideskirt INT NULL,
    exhaust INT NULl,
    wheel INT NULL,
    lamp INT NULL,
    nitro INT NULL,
    stereo INT NULL,
    hydraulics INT NULL,
    bullbar INT NULL,
    rear_bullbar INT NULL,
    front_bullbar INT NULL,
    front_bumper INT NULL,
    rear_bumper INT NULL,
    vent INT NULL,
    FOREIGN KEY (owner_id) REFERENCES player (id) ON DELETE CASCADE ON UPDATE CASCADE
);

CREATE TABLE IF NOT EXISTS house_interior
(
    id          SERIAL PRIMARY KEY,
    house_id    INT NOT NULL,
    interior_id INT NOT NULL,
    FOREIGN KEY (house_id) REFERENCES house (id) ON DELETE CASCADE ON UPDATE CASCADE,
    UNIQUE (house_id, interior_id)
);

CREATE TABLE IF NOT EXISTS house_slot
(
    id         SERIAL PRIMARY KEY,
    house_id   INT NOT NULL,
    vehicle_id INT NOT NULL,
    FOREIGN KEY (house_id) REFERENCES house (id) ON DELETE CASCADE ON UPDATE CASCADE,
    FOREIGN KEY (vehicle_id) REFERENCES vehicle (id) ON DELETE CASCADE ON UPDATE CASCADE
);

-- Functions
CREATE OR REPLACE FUNCTION update_or_create_player_job(pid INT, jid INT, sc INT)
    RETURNS player_job AS
$$
DECLARE
    player_job_return player_job%ROWTYPE;
BEGIN
    IF EXISTS (SELECT FROM player_job WHERE player_job.player_id = pid AND player_job.job_id = jid) THEN
        UPDATE player_job
        SET score = sc
        WHERE player_job.player_id = pid
          AND player_job.job_id = jid
        RETURNING * INTO player_job_return;
    ELSE
        INSERT INTO player_job(player_id, job_id, score)
        VALUES (pid, jid, sc)
        RETURNING * INTO player_job_return;
    END IF;
    return player_job_return;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION update_or_create_player_membership(pid INT, cid INT)
    RETURNS TABLE
            (
                company_member      company_member,
                company_member_info company_member_info
            )
AS
$$
DECLARE
    company_member_return      company_member%ROWTYPE;
    company_member_info_return company_member_info%ROWTYPE;
    company_member_exists      bool;
    company_member_info_exists bool;
BEGIN
    SELECT TRUE
    INTO company_member_exists
    FROM company_member
    WHERE player_id = pid
      AND company_id = cid;
    SELECT TRUE
    INTO company_member_info_exists
    FROM company_member_info
    WHERE player_id = pid
      AND company_id = cid;
    IF company_member_exists THEN
        SELECT *
        INTO company_member_return
        FROM company_member
        WHERE player_id = pid
          AND company_id = cid;
    ELSE
        INSERT INTO company_member (player_id, company_id)
        VALUES (pid, cid)
        RETURNING * INTO company_member_return;
    END IF;
    IF company_member_info_exists THEN
        SELECT *
        INTO company_member_info_return
        FROM company_member_info
        WHERE player_id = pid
          AND company_id = cid;
    ELSE
        INSERT INTO company_member_info (player_id, company_id)
        VALUES (pid, cid)
        RETURNING * INTO company_member_info_return;
    END IF;

    RETURN QUERY SELECT company_member_return, company_member_info_return;
END;
$$ LANGUAGE plpgsql;
