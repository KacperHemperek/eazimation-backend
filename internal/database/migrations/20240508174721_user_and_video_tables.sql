-- +goose Up
-- +goose StatementBegin

create table if not exists users
(
    id serial primary key,
    email text not null unique,
    name text not null
);

create table if not exists rendered_video
(
    id serial primary key,
    user_id int not null references users,
    video_id text not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table if exists users;

drop table if exists rendered_video;

-- +goose StatementEnd

