CREATE OR REPLACE FUNCTION read_actors_name (
    actorName       text
)
RETURNS TABLE (
    name        text
) AS $$
DECLARE
    tempName    text;
BEGIN
    SELECT a.name INTO tempName
    FROM actors a
    WHERE a.name = actorName;

    IF tempName IS NOT NULL THEN
        RETURN QUERY
        SELECT a.name
        FROM actors a
        WHERE a.name = actorName;
    ELSE
        RETURN QUERY
        SELECT a.name
        FROM actors a
        ORDER BY levenshtein(lower(a.name), lower(actorName)) 
        LIMIT 5;
    END IF;
END;
$$ LANGUAGE plpgsql;
