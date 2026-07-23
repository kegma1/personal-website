CREATE TABLE posts (
    id            INTEGER   PRIMARY KEY,
    title         text      NOT     NULL,
    subtitle      text      NOT     NULL,
    content       text      NOT     NULL,
    created_at    DATETIME  DEFAULT CURRENT_TIMESTAMP,
    language      text      NOT     NULL,
    tags          text      NOT     NULL  -- comma seperated value
)
