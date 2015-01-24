package db

// How to perform a database update:
// 1) increment the latest version number (here and in schema.sql)
// 2) if the change requires modificiation of the table schemas, update the
//    schemas here and in schema.sql
// 3) add a DatabaseUpdate for migration of existing data to Updates
// 4) commit all of the above as a single commit

// Latest version of the database
const Latest DatabaseVersion = 3

// Database update history.
// Field 'From' and 'To' are the version numbers before and after the update.
// Field 'SQL' is a list of SQL statements to execute.
var Updates = []*DatabaseUpdate{
    &DatabaseUpdate{
        From: 1,
        To: 2,
        SQL: []string{
            "alter table mix_user add column name varchar(255)",
        },
    },
    &DatabaseUpdate{
        From: 2,
        To: 3,
        SQL: []string{
            `create table mix_playlist_star (
                pid         integer references mix_playlist (pid),
                uid         integer references mix_user (uid),
                tstamp      timestamp
            )`,
        },
    },
}

// LatestSchema is a list of table creation statements, accurate to the version
// stored in Latest. 
var LatestSchema = []Table{
    Table{"mix_user", `create table mix_user (
        uid         serial primary key,
        name        varchar(255)
    )`},

    Table{"mix_playlist", `create table mix_playlist (
        pid         serial primary key,
        title       varchar(255),
        owner_uid   integer references mix_user (uid)
    )`},

    Table{"mix_playlist_star", `create table mix_playlist_star (
        pid         integer references mix_playlist (pid),
        uid         integer references mix_user (uid),
        tstamp      timestamp
    )`},

    Table{"mix_playlist_tag", `create table mix_playlist_tag (
        pid         integer references mix_playlist (pid),
        tag         varchar(255),
        primary key (pid, tag)
    )`},

    Table{"mix_playlist_entry", `create table mix_playlist_entry (
        eid         bigserial primary key,
        pid         integer references mix_playlist (pid),
        index       smallint,
        yt_id       char(11),
        title       varchar(255),
        artist      varchar(255),
        album       varchar(255),
        duration    smallint
    )`},
}
