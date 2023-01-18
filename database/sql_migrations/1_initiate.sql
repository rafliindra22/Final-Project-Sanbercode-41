-- +migrate Up

-- +migrate StatementBegin
CREATE TABLE users (
    id SERIAL,
    name VARCHAR(256),
    username VARCHAR(256),
    password VARCHAR(256),
    full_acces BOOLEAN,
    created_at TIMESTAMP(6),
    updated_at TIMESTAMP(6),
    CONSTRAINT user_pkey PRIMARY KEY (id),
    CONSTRAINT user_unkey UNIQUE (username)
);
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE TABLE brand (
    id SERIAL,
    name VARCHAR(256),
    description VARCHAR(256),
    created_at TIMESTAMP(6),
    updated_at TIMESTAMP(6),
    user_id INTEGER,
    CONSTRAINT brand_pkey PRIMARY KEY (id),
    CONSTRAINT fk_brand_user FOREIGN KEY (user_id) REFERENCES users(id)
);
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE TABLE phone (
    id SERIAL,
    type VARCHAR(256),
    year INTEGER,
    brand_id INTEGER,
    created_at TIMESTAMP(6),
    updated_at TIMESTAMP(6),
    editor_name VARCHAR(256),
    user_id INTEGER,
    CONSTRAINT phone_pkey PRIMARY KEY (id),
    CONSTRAINT fk_phone_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_phone_brand FOREIGN KEY (brand_id) REFERENCES brand(id)
);
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE TABLE spec (
    id SERIAL,
    phone_id INTEGER,
    processor VARCHAR(256),
    memory VARCHAR(256),
    storage VARCHAR(256),
    screen VARCHAR(256),
    camera VARCHAR(256),
    price VARCHAR(256),
    review VARCHAR(256),
    created_at  TIMESTAMP(6),
    updated_at TIMESTAMP(6),
    user_id INTEGER,
    CONSTRAINT spec_pkey PRIMARY KEY (id),
    CONSTRAINT fk_spec_phone FOREIGN KEY (phone_id) REFERENCES phone(id)
);
-- +migrate StatementEnd

-- +migrate StatementBegin
CREATE TABLE commentphone (
    id SERIAL,
    phone_id INTEGER,
    name VARCHAR(256),
    comment VARCHAR(256),
    created_at TIMESTAMP(6),
    user_id INTEGER,
    CONSTRAINT comment_pkey PRIMARY KEY (id),
    CONSTRAINT fk_comment_user FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT fk_commeny_phone FOREIGN KEY (phone_id) REFERENCES phone(id)
);
-- +migrate StatementEnd