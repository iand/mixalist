package db

// How to perform a database update:
// 1) increment the latest version number (here and in schema.sql)
// 2) if the change requires modificiation of the table schemas, update the
//    schemas here and in schema.sql
// 3) add a DatabaseUpdate for migration of existing data to Updates
// 4) log the update in the schema version changelog at the bottom of schema.sql
// 5) commit all of the above as a single commit

// Latest version of the database
const Latest DatabaseVersion = 9

// Database update history.
// Field 'From' and 'To' are the version numbers before and after the update.
// Field 'SQL' is a list of SQL statements to execute.
var Updates = []*DatabaseUpdate{
	&DatabaseUpdate{
		From: 1,
		To:   2,
		SQL: []string{
			"alter table mix_user add column name varchar(255)",
		},
	},
	&DatabaseUpdate{
		From: 2,
		To:   3,
		SQL: []string{
			`create table mix_playlist_star (
				pid         integer references mix_playlist (pid),
				uid         integer references mix_user (uid),
				tstamp      timestamp
			)`,
		},
	},
	&DatabaseUpdate{
		From: 3,
		To:   4,
		SQL: []string{
			"alter table mix_playlist add column created timestamp",
		},
	},
	&DatabaseUpdate{
		From: 4,
		To:   5,
		SQL: []string{
			"alter table mix_playlist_entry add column search_text varchar",
			"update mix_playlist_entry set search_text = title || ' ' || artist || ' ' || album",
		},
	},
	&DatabaseUpdate{
		From: 5,
		To:   6,
		SQL: []string{
			"update mix_playlist_entry set search_text = lower(search_text)",
		},
	},
	&DatabaseUpdate{
		From: 6,
		To:   7,
		SQL: []string{
			"alter table mix_playlist add column search_text varchar",
			`update mix_playlist
			 set search_text = subquery.search_text
			 from (
				select mix_playlist.pid as pid,
					   lower(mix_playlist.title || ' ' || string_agg(mix_playlist_tag.tag, ' ')) as search_text
				from mix_playlist
				inner join mix_playlist_tag on mix_playlist.pid = mix_playlist_tag.pid
				group by mix_playlist.pid
			 ) as subquery
			 where mix_playlist.pid = subquery.pid`,
		},
	},
	&DatabaseUpdate{
		From: 7,
		To:   8,
		SQL: []string{
			"alter table mix_playlist add column parent_pid integer references mix_playlist (pid)",
		},
	},
	&DatabaseUpdate{
		From: 8,
		To:   9,
		SQL: []string{
			"alter table mix_playlist add column image_blob_id char(32)",
			"update mix_playlist set image_blob_id = ''",
			"alter table mix_playlist_entry add column image_blob_id char(32)",
			"update mix_playlist_entry set image_blob_id = ''",
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
		owner_uid   integer references mix_user (uid),
		created     timestamp,
		search_text varchar,
		parent_pid  integer references mix_playlist (pid),
		image_blob_id char(32)
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
		duration    smallint,
		search_text varchar,
		image_blob_id char(32)
	)`},
}
