BEGIN;

DO
$do$
DECLARE
    migration integer := 1;
BEGIN
    IF NOT EXISTS (SELECT 1 FROM migrations WHERE migration_id = migration) THEN

        CREATE TABLE users
        (
            id SERIAL,
    		FirstName  varchar(20) NOT NULL,
    		LastName   varchar(20) NOT NULL,
            Email      varchar(20) NOT NULL,
    		Password   varchar(40) NOT NULL,
    		ctime      TIMESTAMP,
    		utime      TIMESTAMP,
    		PRIMARY KEY (id)
        );

        INSERT INTO migrations (migration_id, ctime) VALUES(1, NOW());
    ELSE
        RAISE NOTICE 'Migration already run: %.', migration;
    END IF;
END
$do$;

COMMIT;