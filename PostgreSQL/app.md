# Overview
This is my small assignment of how to use PostgreSQL and Go Lang
The app will be simple and console-interaction.

The app will have the function of getting movies based on genre and actor. It will show all result for strict matching and top 5 (genres or actors) for fuzzy searching.

# Functions
- Receive command from the console. COMMAND will have a format of "[COMMAND] [DETAIL]". Each command will have their own detail.
- Print out the result (in other thread) - multithread processing

# Commands
READ_GENRES //print generes and position
READ_MOVIE title //strictly matching title
READ_ACTOR name //strictly matching if found, otherwise fuzzy searching

SIMILAR_MOVIES_GENRE 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 //genre type
SIMILAR_MOVIES title


INSERT_ACTOR name
INSERT_MOVIE title 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 // values of 18 genres

# Credit
The assignment is inspired by the book "7 Database in 7 weeks" - Chapter 2. 
The data and the data model is from the book.

