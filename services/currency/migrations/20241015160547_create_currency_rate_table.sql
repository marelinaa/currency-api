-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS currency_rate (
    id BIGSERIAL PRIMARY KEY,
    date DATE NOT NULL UNIQUE,
    rate NUMERIC(20, 15) NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS currency_rate;
-- +goose StatementEnd