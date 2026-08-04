CREATE TABLE posts (
    id            INTEGER   PRIMARY KEY,
    title         text      NOT     NULL,
    subtitle      text      NOT     NULL,
    content       text      NOT     NULL,
    created_at    DATE      NOT     NULL,
    edited_at     DATE              NULL,
    language      text      NOT     NULL,
    tags          text      NOT     NULL  -- comma seperated value
)
