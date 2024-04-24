-- +goose Up
-- +goose StatementBegin
create table category
(
    id serial primary key,
    name text not null,
    slug text not null,
    created_at timestamp not null default now(),
    updated_at timestamp
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table category;
-- +goose StatementEnd
