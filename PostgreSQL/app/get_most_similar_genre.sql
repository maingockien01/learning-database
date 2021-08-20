
CREATE OR REPLACE FUNCTION get_similar_movies_genre (
    movieTitle           text
)
RETURNS TABLE (
    title   text
) AS $$
DECLARE
    movieGenres cube;
BEGIN
    SELECT genre INTO movieGenres
    FROM movies m
    WHERE m.title = movieTitle;

    IF movieGenres IS NOT NULL THEN
        RETURN QUERY
        SELECT m.title AS title
        FROM movies m
        ORDER BY cube_distance(m.genre, movieGenres)
        LIMIT 5;
    END IF;
END;
$$ LANGUAGE plpgsql;

CREATE OR REPLACE FUNCTION get_similar_movies (
    movieGenres  cube,
    moviesNumber int
) RETURNS TABLE (
    title   text,
    genres  cube
) AS $$
BEGIN
    IF movieGenres IS NOT NULL THEN
        RETURN QUERY
        SELECT m.title AS title, m.genre AS genre
        FROM movies m
        ORDER BY cube_distance(m.genre, movieGenres)
        LIMIT moviesNumber;
    ELSE
        RAISE NOTICE 'Movie is not found!';
    END IF;
END;
$$ LANGUAGE plpgsql;


