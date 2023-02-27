-- +goose Up
-- +goose StatementBegin
CREATE TABLE playlist (
  id bigserial primary key,
  name varchar(255) unique not null,
  duration float(8) not null
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP playlist;
-- +goose StatementEnd
