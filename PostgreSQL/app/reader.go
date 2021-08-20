package main
import (

    "github.com/jackc/pgx/v4"
    "github.com/jackc/pgx/v4/pgxpool"

    "strings"
    "context"
    "fmt"
    "strconv"
)

type Command struct {
    CommandType string;
    Details     string;
}

func ReadLine (line string) Command {
    commandInfo := strings.SplitN(line, " ", 2);
    if len(commandInfo) == 1 {
        return Command{commandInfo[0], ""}
    } else {
        return Command{commandInfo[0], commandInfo[1]}
    }
}

func (command *Command) PerformCommand() string {
    var result string;
    switch command.CommandType {
        case "READ_GENRES":
            result = readGenres()
        case "READ_MOVIE":
            result = readMovies(command.Details)
        case "READ_ACTOR":
            result = readActors(command.Details)
        case "SIMILAR_MOVIES":
            result = suggMovies(command.Details)
        default:
            result = "Command is not valid"
    }
    return result
}

func readGenres () string {
    query := func (conn *pgxpool.Conn) (pgx.Rows, error) {
        return conn.Query(context.Background(), "SELECT name, position FROM genres")
    }

    var name []string
    var position []int

    rowProcessor := func (rows pgx.Rows) error {
        var tempName string;
        var tempPosition int;

        err := rows.Scan(&tempName, &tempPosition)

        name = append(name, tempName)
        position = append(position, tempPosition)

        return err
    }

    QueryPipeline(query, rowProcessor)

    result := "All available genre: \n"
    for i := range name {
        result += strconv.Itoa(position[i]) + " - " + name[i] + "\n"
    }
    return result
}

func readMovies (title string) string {
    fmt.Println(title)
    query := func (conn *pgxpool.Conn) (pgx.Rows, error) {
        return conn.Query(context.Background(), "SELECT title, genre FROM movies WHERE title = $1;", title)
    }

    var movieTitle string
    var movieGenre string


    rowProcessor := func (rows pgx.Rows) error {
        err := rows.Scan(&movieTitle, &movieGenre)

        return err
    }

    QueryPipeline(query, rowProcessor)

    if movieTitle == "" {
        return "Movie is not found"
    } else {
        return movieTitle + " - " + movieGenre
    }

}

func readActors (name string) string {

    query := func (conn *pgxpool.Conn) (pgx.Rows, error) {
        return conn.Query(context.Background(), "SELECT read_actors_name($1)", name)
    }

    var actorNames []string

    rowProcessor := func (rows pgx.Rows) error {
        var tempName string
        err := rows.Scan(&tempName)

        actorNames = append(actorNames, tempName)

        return err
    }

    QueryPipeline(query, rowProcessor)

    result := fmt.Sprintf("Found %d actors: \n", len(actorNames))
    for _, name := range actorNames {
        result += name
        result += "\n"
    }

    return result
}



func suggMovies (title string) string {

    query := func (conn *pgxpool.Conn) (pgx.Rows, error) {
        return conn.Query(context.Background(), "SELECT get_similar_movies_genre($1)", title)
    }

    var movieTitles []string

    rowProcessor := func (rows pgx.Rows) error {
        var tempTitle string

        err := rows.Scan(&tempTitle)

        movieTitles = append(movieTitles, tempTitle)

        return err
    }

    QueryPipeline(query, rowProcessor)

    if len(movieTitles) == 0 {
        return "Not found given movie title"
    }

    result := fmt.Sprintf("Found %d movies : \n", len(movieTitles))
    for _, movie := range movieTitles {
        result += movie
        result += "\n"
    }

    return result
}

