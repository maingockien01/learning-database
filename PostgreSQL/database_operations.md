# Installation


# Running

# Administrating DB

This is not the main point of this projects so I won't go deep into this topic. 
I will list related problems and solutions which I face during the project here. I can be used as the reference for myself and for anyone else review my project.

## Create DB
Link: https://stackoverflow.com/questions/43734650/createdb-database-creation-failed-error-permission-denied-to-create-database

for Unix users:

You need to change to `postgres` user to be system user
```
$ sudo su - postgres
```

Now, you are system user, you can create databases by using command `createdb [mydb]` in the command line.
```
$ createdb 7db
```

## Create user


## Access to DB

You need to be superuser like postgre or change to users granted permission to be able to access the database. More information could be found in the documentation of postgresSQL about users config.

```
psql 7db
```

# Operations

CRUD = CREATE, READ, UPDATE, DELETE.
As long as you can perform CRUD operations, you can do any operations no matter how complicated it is with basic CRUD. 

## Conditions
### LIKE

### AND

### OR

### EQUAL

### NOT EQUAL

## Create table / CREATE
```
CREATE TABLE countries (
country_code char(2) PRIMARY_KEY,
country_name text UNIQUE
);
```

## Insert data / CREATE
Either you can list the order of the column that you will insert
```
INSERT INTO countries (country_code, country_name)
VALUES 
('us', 'United States'),
('mx', 'Mexico'),
('au', 'Australia'),
('gb', 'United Kingdom'),
('de', 'Germany'),
('ll', 'Loompaland');
```

Or you can skip specify your order and assume the order as the order in the schema of the table.
```
INSERT INTO countries 
VALUES ('uk', 'United Kingdom');
);
```

## SELECT

`*` is used as the wildcard to include all columns in the table.

```
SELECT * FROM countries;
```

## DELETE

```
DELETE FROM countries
WHERE country_code = 'll';
```

## Data condition on table schema
`CHECK` is used as the schema condition for specified columns

## Primary key
Primary key can be one column or a set of columns

## Table with foreign key

```
CREATE TABLE cities (
name  text NOT NULL, postal_code varchar (9) CHECK (postal_code <> ''), 
country_code char(2) REFERENCES countries,
PRIMARY KEY (country_code, postal_code)
);
```

## UPDATE

```
UPDATE cities
SET postal_code = '97206'
WHERE name='Portland';
```

## INNER JOIN read

```
SELECT cities.*, country_name
FROM cities INNER JOIN countries
ON cities.country_code = countries.country_code;
```

### JOIN read with compound key

```
CREATE TABLE
7db=# CREATE TABLE venues (
venue_id SERIAL PRIMARY KEY,
name varchar(255),
street_address text,
type char(7) CHECK (type in ('public', 'private')) DEFAULT 'public',
postal_code varchar(9),
country_code char(2),
FOREIGN KEY (country_code, postal_code)
REFERENCES cities (country_code, postal_code) MATCH FULL
);
```

```
INSERT INTO venues (name, postal_code, country_code) 
VALUES ('Crystal Ballroom', '97206', 'us');
```

Join with compound key 
```
SELECT v.venue_id, v.name, c.name FROM venues as v INNER JOIN cities AS c ON v.postal_code = c.postal_code AND v.country_code = c.country_code;
```

## INSERT with returning data
```
INSERT INTO venues (name, postal_code, country_code) VALUES ('Voodoo Doughnut', '97206', 'us') RETURNING venue_id;
```

## OUTER JOIN
Outer join is the way of merging 2 tables when the results of one table must always be returned event there is no matching data on the other table. 
LEFT JOIN is outer join with left table is returned, while right join returns right table.

## Index
### Hash index
```
CREATE INDEX events_title ON events USING hash(title);
```

### BTree index
```
SELECT * FROM events WHERE starts >= '2018-04-01';
```

```
CREATE INDEX events_starts ON events USING btree(starts);
```

## JOIN multiple table
```
SELECT ctr.country_name FROM events AS e
JOIN venues AS v ON e.venue_id = v.venue_id
JOIN countries as ctr ON v.country_code = ctr.country_code
WHERE e.title = 'Fight Club';
```

# Advanced queries

## Aggregate functions
List: https://www.postgresql.org/docs/9.5/functions-aggregate.html

```
SELECT count(title) FROM events WHERE title LIKE '%Day%';
```

```
SELECT min(starts), max(ends) FROM events JOIN venues ON events.venue_id = venues.venue_id WHERE venues.name = 'Crystal Ballroom';
```

## sub-SELECT

Sub-SELECT is like a sub function return a value in a bigger function.
Instead of querrying 2 functions separatedly, we combine them into 1 and query one time only.

```
INSERT INTO events (title, starts, ends, venue_id) VALUES ('Moby', '2018-02-06 21:00', '2018-02-06 23:00', (SELECT venue_id FROM venues WHERE name = 'Crystal Ballroom'));
```

## Grouping
with GROUP BY, you tell PostgreSQL to place rows in groups and perform some kind of aggregate functions on those group.

```
SELECT venue_id, count(*)
FROM events
GROUP BY venue_id;
```
This command, we tell Postgres to group events based on venue_id then perfrom count function on each group.

```
SELECT venue_id FROM events GROUP BY venue_id HAVING count(*) >= 2 AND venue_id IS NOT NULL;
```

Select data without repeating data
```
SELECT venue_id FROM events GROUP BY venue_id;
```
Or
```
SELECT DISTINCT venue_id FROM events;
```

## Window function
```
SELECT title, count(*) OVER (PARTITION BY venue_id) FROM events;
```

## Transaction
`All or nothing`

Transaction ensures that every command of a set is executed. If anything fails along the way, all of the commands are rolled back as if they never happened.

PostgreSQL follow ACID compliance:
- Atomic (either all operations succeed or none do)
- Consistent (the data will always be in a good state and never in an inconsistent state)
- Isolated (transaction don't interfere with one another)
- Durable (a commited transaction is safe, event after a server crash)

```
BEGIN TRANSACTION;
    DELETE FROM events;
ROLLBACK;
SELECT * FROM events;
```

```
BEGIN TRANSACTION;
    ...
END;
```

Transaction is useful when you want to keep 2 tables in sync after operations regardless all unexpected situations.

## Stored Procedure
Customize SQL, CRUD operations.

## Pull the trigger
Trigger auto fire stored procedures when some event happens, such as an insert or update.

## Viewing the world
Use the result of a complex query just like a table.

```
CREATE VIEW holidays AS SELECT event_id AS holiday_id, title AS name, starts AS date FROM events WHERE title LIKE '%Day%' AND venue_id IS NULL;
```

```
SELECT name, to_char(date, 'Month DD, YYYY') AS date FROM holidays WHERE date <= '2018-04-01';
```

## RULEs
A RULE is a description of how to alter the parsed query tree.

## Pivot table
To use crsstab(), the query must return 3 columns: ID, category, and value.


