create table user (
    uid         serial primary key                  # user ID
);

create table playlist (
    pid         serial primary key,                 # playlist ID
    title       varchar(255),                       # title of playlist
    owner_uid   integer references user (uid)       # ID of user that owns this playlist
);

create table playlist_tag (
    pid         integer references playlist (pid),  # ID of playlist being tagged
    tag         varchar(255)                        # tag name
    primary key (pid, tag)
);

create table playlist_entry (
    eid         bigserial primary key,              # entry ID
    pid         integer references playlist (pid),  # ID of playlist this entry is in
    index       smallint,                           # index of entry within playlist (0-based)
    yt_id       char(11),                           # 11-character YouTube video ID (https://www.youtube.com/watch?v=xxxxxxxxxxx)
    title       varchar(255),                       # track title (can be edited)
    artist      varchar(255),                       # track artist (can be edited)
    album       varchar(255),                       # track album (can be edited)
    duration    smallint                            # duration of video in seconds
);
