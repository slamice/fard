BEGIN;

DO
$do$
DECLARE
    migration integer := 0;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM migrations WHERE migration_id = migration) THEN

        CREATE TABLE migrations
        (
            id SERIAL,
            migration_id INTEGER NOT NULL UNIQUE,
            ctime TIMESTAMP,
    		PRIMARY KEY (id)
        );

        IF NOT EXISTS (SELECT 1 FROM migrations WHERE migration_id=0)
        INSERT INTO migrations (migration_id, ctime)
        VALUES (0, current_timestamp);

    ELSE
        RAISE NOTICE 'Migration already run: %.', migration;
    END IF;
END
$do$;

COMMIT;