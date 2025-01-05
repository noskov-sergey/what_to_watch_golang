-- +goose Up
-- +goose StatementBegin
ALTER TABLE opinions
DROP COLUMN added_by;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE opinions
ADD added_by VARCHAR;
-- +goose StatementEnd
