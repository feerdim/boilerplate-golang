CREATE TABLE IF NOT EXISTS directories (
    guid varchar NOT NULL PRIMARY KEY,
    directory_guid varchar,
    name varchar NOT NULL,
    description text,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'UTC')::TIMESTAMP,
    created_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    updated_at timestamp without time zone,
    updated_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    deleted_at timestamp without time zone,
    deleted_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL
);

ALTER TABLE directories
    ADD CONSTRAINT directories_directory_guid_fkey FOREIGN KEY (directory_guid) REFERENCES directories(guid) ON UPDATE CASCADE ON DELETE RESTRICT;

CREATE TABLE IF NOT EXISTS files (
    guid varchar NOT NULL PRIMARY KEY,
    directory_guid varchar NOT NULL REFERENCES directories(guid) ON UPDATE CASCADE ON DELETE RESTRICT,
    name varchar NOT NULL,
    description text,
    path text NOT NULL,
    size bigint NOT NULL,
    extension varchar NOT NULL,
    mime_type varchar NOT NULL,
    created_at timestamp without time zone NOT NULL DEFAULT (now() at time zone 'UTC')::TIMESTAMP,
    created_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    updated_at timestamp without time zone,
    updated_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL,
    deleted_at timestamp without time zone,
    deleted_by varchar REFERENCES users(guid) ON UPDATE CASCADE ON DELETE SET NULL
);
