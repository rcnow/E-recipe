CREATE TABLE recipe (
    id         INTEGER PRIMARY KEY,
    recipename TEXT,
    bottleid   TEXT,
    bottlesize INTEGER,
    pg         INTEGER,
    vg         INTEGER,
    other      INTEGER,
    nicotine   INTEGER,
    date       TEXT,
    steeptime  INTEGER,
    note       TEXT
);


CREATE TABLE flavor (
    flavor_name    TEXT,
    flavor_percent INTEGER,
    recipeid       INTEGER REFERENCES recipe (id) ON DELETE CASCADE
);

