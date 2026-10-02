CREATE TABLE posts (
    id            INTEGER   PRIMARY KEY,
    title         text      NOT     NULL,
    subtitle      text      NOT     NULL,
    content       text      NOT     NULL,
    created_at    DATE      NOT     NULL,
    edited_at     DATE              NULL,
    language      text      NOT     NULL,
    tags          text      NOT     NULL  -- comma seperated value
);

CREATE TABLE projects (
    id            INTEGER   PRIMARY KEY,
    title         text      NOT     NULL,
    link          text      NOT     NULL,
    icon_path     text      NOT     NULL
);
