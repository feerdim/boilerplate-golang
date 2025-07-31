CREATE TABLE IF NOT EXISTS roles (
    guid varchar NOT NULL PRIMARY KEY,
    name varchar NOT NULL,
    description varchar,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'UTC')::TIMESTAMP,
    created_by varchar,
    updated_at timestamp without time zone,
    updated_by varchar
);

CREATE TABLE IF NOT EXISTS users (
    guid varchar NOT NULL PRIMARY KEY,
    name varchar NOT NULL,
    email varchar NOT NULL UNIQUE,
    password varchar NOT NULL,
    verified_at timestamp without time zone,
    activated_at timestamp without time zone,
    activated_by varchar,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'UTC')::TIMESTAMP,
    created_by varchar,
    updated_at timestamp without time zone,
    updated_by varchar,
    deleted_at timestamp without time zone,
    deleted_by varchar
);

ALTER TABLE roles
    ADD CONSTRAINT roles_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    ADD CONSTRAINT roles_updated_by_fkey FOREIGN KEY (updated_by) REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL;

INSERT INTO roles (guid, name, created_at)
VALUES
('019288cc-3c8c-7484-ac0d-3362a88ae018', 'Admin', (now() at time zone 'UTC')::TIMESTAMP);

ALTER TABLE users
    ADD CONSTRAINT users_activated_by_fkey FOREIGN KEY (activated_by) REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    ADD CONSTRAINT users_created_by_fkey FOREIGN KEY (created_by) REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    ADD CONSTRAINT users_updated_by_fkey FOREIGN KEY (updated_by) REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    ADD CONSTRAINT users_deleted_by_fkey FOREIGN KEY (deleted_by) REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL;

INSERT INTO users (guid, name, email, password, verified_at, activated_at, created_at)
VALUES
('01928589-56f2-7065-ac68-f63261a05b4f', 'Admin', 'admin@gmail.com', '$2a$10$FE8avrLD/pyO.bDtYs82u.aYLaYtNx7zrpWVeXZNa.zHFeD3QLfzW', (now() at time zone 'UTC')::TIMESTAMP, (now() at time zone 'UTC')::TIMESTAMP, (now() at time zone 'UTC')::TIMESTAMP);

CREATE TABLE IF NOT EXISTS role_user (
    role_guid varchar NOT NULL REFERENCES roles(guid) ON UPDATE CASCADE ON DELETE CASCADE,
    user_guid varchar NOT NULL REFERENCES users(guid) ON UPDATE CASCADE ON DELETE CASCADE
);

INSERT INTO role_user (role_guid, user_guid)
VALUES
('019288cc-3c8c-7484-ac0d-3362a88ae018', '01928589-56f2-7065-ac68-f63261a05b4f');

CREATE TABLE IF NOT EXISTS permission_groups (
    guid varchar NOT NULL PRIMARY KEY,
    name varchar NOT NULL,
    description varchar,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'UTC')::TIMESTAMP,
    created_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    updated_at timestamp without time zone,
    updated_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL
);

INSERT INTO permission_groups (guid, name, created_at)
VALUES
('0195aa10-caee-7e5f-b118-b762943c8c48', 'Account Management', (now() at time zone 'UTC')::TIMESTAMP);

CREATE TABLE IF NOT EXISTS permissions (
    guid varchar NOT NULL PRIMARY KEY,
    permission_group_guid varchar NOT NULL REFERENCES permission_groups(guid) ON UPDATE CASCADE ON DELETE CASCADE,
    name varchar NOT NULL,
    description varchar,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'UTC')::TIMESTAMP,
    created_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    updated_at timestamp without time zone,
    updated_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL
);

INSERT INTO permissions (guid, permission_group_guid, name, created_at)
VALUES
('0195aa10-caee-7e5f-b118-b762943c8c48', '0195aa10-caee-7e5f-b118-b762943c8c48', 'Create', (now() at time zone 'UTC')::TIMESTAMP),
('0195aa11-451b-7ef4-be39-099ddeafe106', '0195aa10-caee-7e5f-b118-b762943c8c48', 'Update', (now() at time zone 'UTC')::TIMESTAMP),
('0195aa11-6b25-78c2-b8a2-730faa7a6a8d', '0195aa10-caee-7e5f-b118-b762943c8c48', 'Delete', (now() at time zone 'UTC')::TIMESTAMP),
('0195aa11-8cc9-728c-87a3-6f056641373c', '0195aa10-caee-7e5f-b118-b762943c8c48', 'View', (now() at time zone 'UTC')::TIMESTAMP);

CREATE TABLE IF NOT EXISTS permission_role (
    permission_guid varchar NOT NULL REFERENCES permissions(guid) ON UPDATE CASCADE ON DELETE CASCADE,
    role_guid varchar NOT NULL REFERENCES roles(guid) ON UPDATE CASCADE ON DELETE CASCADE
);

INSERT INTO permission_role (permission_guid, role_guid)
VALUES
('0195aa10-caee-7e5f-b118-b762943c8c48', '019288cc-3c8c-7484-ac0d-3362a88ae018'),
('0195aa11-451b-7ef4-be39-099ddeafe106', '019288cc-3c8c-7484-ac0d-3362a88ae018'),
('0195aa11-6b25-78c2-b8a2-730faa7a6a8d', '019288cc-3c8c-7484-ac0d-3362a88ae018'),
('0195aa11-8cc9-728c-87a3-6f056641373c', '019288cc-3c8c-7484-ac0d-3362a88ae018');

CREATE TABLE IF NOT EXISTS user_token_validations (
    guid varchar NOT NULL PRIMARY KEY,
    email varchar NOT NULL,
    type varchar NOT NULL,
    token varchar NOT NULL,
    expires_at timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'UTC')::TIMESTAMP,
    UNIQUE (email, type)
);

CREATE TABLE IF NOT EXISTS sessions (
    guid varchar NOT NULL PRIMARY KEY,
    user_guid varchar NOT NULL REFERENCES users(guid) ON UPDATE CASCADE ON DELETE CASCADE,
    access_token varchar NOT NULL,
    access_token_expires_at timestamp without time zone NOT NULL,
    refresh_token varchar NOT NULL,
    refresh_token_expires_at timestamp without time zone NOT NULL,
    ip_address varchar NOT NULL,
    user_agent varchar NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'UTC')::TIMESTAMP,
    updated_at timestamp without time zone
);
