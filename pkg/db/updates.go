package db

// How to perform a database update:
// 1) increment the latest version number (here and in schema.sql)
// 2) add a DatabaseUpdate for migration of existing data to Updates
// 3) if the change required modificiation of the table schemas, update the
//    schemas here and in schema.sql
// 4) commit all of the above as a single commit

// Latest version of the database
const Latest DatabaseVersion = 1

// Database update history.
// Field 'From' and 'To' are the version numbers before and after the update.
// Field 'SQL' is a list of SQL statements to execute.
var Updates = []*DatabaseUpdate{
    
}

// LatestSchema is a list of table creation statements, accurate to the version
// stored in Latest. 
var LatestSchema = []Table{
    Table{"mix_user", `create table mix_user (
        uid         serial primary key
    )`},

    Table{"mix_playlist", `create table mix_playlist (
        pid         serial primary key,
        title       varchar(255),
        owner_uid   integer references mix_user (uid)
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
