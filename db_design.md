# Database Design

This document is to show a brief sketch of the database design.
Once the schema is implemented in the database then this document
will become obsolete and I'll just remove it from the source.

## Tables

Database tables in the database

### Users

This table holds information about the users in a particular channel.

| Field | Type | Description |
|-------|------|-------------|
| id | char(36), PrimaryKey | the user id (uuid v4) |
| user_name | varchar(50) | The username of the user (this should also be the mention name) |
| steam_name | varchar(50) | The steam username of the user |
| psn_name | varchar(50) | The psn username of the user |
| xbox_name | varchar(50) | The xbox username of the user |
| twitch_name | varchar(50) | The twitch username of the user |
| rank | tinyint | The current rank of the member |
| roles | varchar | semicolon-split roles that a user has |
| permissions | bigint | bit flags for permissions |
| deleted | bool | soft-delete flag|
| created | datetime | the date the record was created |
| updated | datetime | the date the record was modified |

### SoundCommands

This table holds admin-uploaded sound filepaths as well as commands that trigger playing of the sounds.

| Field | Type | Description|
|-------|------|------------|
| id | char(36), PrimaryKey | The id of the sound command hook|
| added_by | char(36), ForeignKey(Users.id) | The id of the user that added the sound command |
| filepath | varchar(256) | The filepath to the sound byte |
| command | varchar(20) | The command to type in to trigger the audio|
| deleted | bool | soft-delete flag|
| created | datetime | the date the record was created |
| updated | datetime | the date the record was modified |

### TwitterSubscriptions

This table lists the twitter feeds the bot listens for updates on.

| Field | Type | Description|
|-------|------|------------|
| id | char(36), PrimaryKey | The id of the twitter subscriptions |
| added_by | char(36), ForeignKey(Users.id) | The id of the user that added the twitter subscription |
| twitter_user | varchar(100) | The twitter user to subscribe to |
| active | bool | flag to indicate whether the subscription is activated |
| deleted | bool | soft-delete flag|
| created | datetime | the date the record was created |
| updated | datetime | the date the record was modified |

### DatabaseLogs

This table lists delete, modify, and create events for each table in the database. Mostly for record keeping. For fixing data issues that's what we have backups for :P.

| Field | Type | Description|
|-------|------|------------|
| log_id | bigint, PrimaryKey | The id of the data log (just an integer that autoincrements) |
| user | char(36), ForeignKey(Users.id) | The id of the user that modified data |
| table_name | varchar(100) | The name of the table where the data was modified |
| row_id | char(36) | the uuid of the row modified |
| event_type | tinyint | the event type (0 deleted, 1 created, 2, updated) |

### Tasks

This table tracks tasks for the bot to run every now and then.

| Field | Type | Description|
|-------|------|------------|
| id | char(36), PrimaryKey | The id of the task |
| added_by | char(36), ForeignKey(Users.id) | The id of the user that added the task |
| task_name | varchar(100) | The name given to the task |
| task_description | varchar(500) | Description of the task |
| last_run | datetime | the datetime when the task was last performed. Can be Null |
| remaining_executions | smallint | How many executions remain. -1 indicates it runs forever |
| interval | bigint | The interval in seconds that this task runs |
| deleted | bool | soft-delete flag|
| created | datetime | the date the record was created |
| updated | datetime | the date the record was modified |
