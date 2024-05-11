-- +goose Up
-- +goose StatementBegin

alter table rendered_video rename to rendered_videos;

alter table rendered_videos
    add column created_at date not null default now(),
    add column updated_at date not null default now(),
    add column video_data json not null default '{}';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists rendered_videos rename to rendered_video;

alter table rendered_video
    drop column created_at,
    drop column updated_at,
    drop column video_data;
-- +goose StatementEnd
