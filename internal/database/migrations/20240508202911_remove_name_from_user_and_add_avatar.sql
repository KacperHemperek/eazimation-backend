-- +goose Up
-- +goose StatementBegin
alter table if exists users drop column if exists name;

alter table if exists users add column avatar text not null default '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
alter table if exists users add column name text not null default '';

alter table if exists users drop column avatar;
-- +goose StatementEnd
