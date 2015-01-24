# Schema version 9

create table mix_version (
    version     integer,                                # current database version
);



create table mix_user (
    uid         serial primary key,                     # user ID
    name        varchar(255)                            # username
);

create table mix_playlist (
    pid         serial primary key,                     # playlist ID
    title       varchar(255),                           # title of playlist
    owner_uid   integer references mix_user (uid),      # ID of user that owns this playlist
    created     timestamp,                              # timestamp of creation of the playlist
    search_text varchar,                                # concenation of title and tags
    parent_pid  integer references mix_playlist (pid),  # ID of playlist that this playlist is a remix of
    image_blob_id char(32),                             # blob ID of composite playlist art
);

create table mix_playlist_star (
    pid         integer references mix_playlist (pid),  # ID of playlist being starred
    uid         integer references mix_user (uid),      # ID of user that starred the playlist
    tstamp      timestamp                               # timestamp of addition of star
);

create table mix_playlist_tag (
    pid         integer references mix_playlist (pid),  # ID of playlist being tagged
    tag         varchar(255)                            # tag name
    primary key (pid, tag)
);

create table mix_playlist_entry (
    eid         bigserial primary key,                  # entry ID
    pid         integer references mix_playlist (pid),  # ID of playlist this entry is in
    index       smallint,                               # index of entry within playlist (0-based)
    yt_id       char(11),                               # 11-character YouTube video ID (https://www.youtube.com/watch?v=xxxxxxxxxxx)
    title       varchar(255),                           # track title (can be edited)
    artist      varchar(255),                           # track artist (can be edited)
    album       varchar(255),                           # track album (can be edited)
    duration    smallint,                               # duration of video in seconds
    search_text varchar,                                # concatenation of title, artist and album (for searching)
    image_blob_id char(32)                              # blob ID of album art
);


/*
Schema version changelog:
    1: initial version
    2: add mix_user.name column
    3: add mix_playlist_star table with columns pid, uid, tstamp
    4: add mix_player.created column
    5: add mix_playlist_entry.search_text column
    6: make mix_playlist_entry.search_text always lowercase
    7: add mix_playlist.search_text column
    8: add mix_playlist.parent_pid column
    9: add mix_playlist.image_blob_id and mix_playlist_entry.image_blob_id columns
*/
