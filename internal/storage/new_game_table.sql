-- +goose Up
-- +goose StatementBegin
CREATE TABLE games (
                       id SERIAL PRIMARY KEY,
                       name VARCHAR(255) NOT NULL,
                       created_at TIMESTAMP NOT NULL,
                       updated_at TIMESTAMP NOT NULL,
                       user_id INTEGER NOT NULL,
                       foreign key (user_id) references users(id),
                       status VARCHAR(255) NOT NULL,
                       winner INTEGER NOT NULL,
                       summary integer NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table if exists games;
-- +goose StatementEnd
