-- +goose Up
CREATE TYPE mission_status AS ENUM ('IN_PROGRESS', 'COMPLETED');
CREATE TABLE missions (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
    status mission_status NOT NULL,

    cat_id uuid REFERENCES cats (id) ON DELETE CASCADE
);

CREATE UNIQUE INDEX unique_cat_id_if_not_null ON missions (cat_id) WHERE cat_id IS NOT NULL;

CREATE TYPE target_status AS ENUM ('IN_PROGRESS', 'COMPLETED');
CREATE TABLE targets (
    id uuid PRIMARY KEY DEFAULT gen_random_uuid(),

    name varchar(128) NOT NULL,
    country varchar(128) NOT NULL,
    notes text NOT NULL,
    status target_status NOT NULL,

    mission_id uuid REFERENCES missions (id) ON DELETE CASCADE NOT NULL
);

-- +goose StatementBegin
CREATE OR REPLACE FUNCTION check_max_targets_per_mission()
RETURNS TRIGGER AS $$
BEGIN
    IF (SELECT COUNT(*) FROM targets WHERE mission_id = NEW.mission_id) > 3 THEN
        RAISE EXCEPTION 'Exceeded maximum number of targets per mission: 3';
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;
-- +goose StatementEnd
CREATE TRIGGER enforce_max_targets_per_mission
BEFORE INSERT OR UPDATE ON targets
FOR EACH ROW
EXECUTE FUNCTION check_max_targets_per_mission();

-- +goose Down
DROP TABLE IF EXISTS targets;
DROP TABLE IF EXISTS missions;
DROP FUNCTION IF EXISTS check_max_targets_per_mission ();
DROP TRIGGER IF EXISTS enforce_max_targets_per_mission ();
