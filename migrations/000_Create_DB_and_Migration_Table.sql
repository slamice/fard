BEGIN;

DO
$do$
DECLARE
    migration integer := 0;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM migrations WHERE migration_id = migration) THEN

        CREATE DATABASE frd WITH ENCODING='UTF8' CONNECTION LIMIT=100;

        CREATE TABLE migrations
        (
            id SERIAL,
            migration_id INTEGER NOT NULL UNIQUE
        );

        INSERT INTO migrations (migration_id, ctime) VALUES(0, NOW()) WHERE NOT EXISTS (SELECT 1 FROM migrations WHERE migration_id=0);

    ELSE
        RAISE NOTICE 'Migration already run: %.', migration;
    END IF;
END
$do$;

COMMIT;