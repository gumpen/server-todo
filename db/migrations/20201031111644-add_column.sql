
-- +migrate Up
ALTER TABLE messages
ADD updated_at DATETIME default current_timestamp;

-- +migrate Down
ALTER TABLE messages
DROP updated_at;