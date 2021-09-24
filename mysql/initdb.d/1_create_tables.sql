USE test_db;

CREATE TABLE projects
(
    id               CHAR(36)    NOT NULL,
    group_id         INT         NOT NULL,
    key_name         VARCHAR(10) NOT NULL,
    name             VARCHAR(80) NOT NULL,
    leader           INT         NOT NULL,
    default_assignee INT         NOT NULL,
    created_date     DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_date     DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE uq_key_name(group_id, key_name),
    UNIQUE uq_name(group_id, name)
);
