USE test_db;

CREATE TABLE products
(
    id         CHAR(36)    NOT NULL,
    group_id   CHAR(36)    NOT NULL,
    name       VARCHAR(80) NOT NULL,
    leader_id  CHAR(36)    NOT NULL,
    created_at DATETIME    NOT NULL,
    updated_at DATETIME    NOT NULL,
    PRIMARY KEY (id),
    UNIQUE uq_name(group_id, name)
);

CREATE TABLE product_notes
(
    id         CHAR(36)     NOT NULL,
    product_id CHAR(36)     NOT NULL,
    group_id   CHAR(36)     NOT NULL,
    title      VARCHAR(255) NOT NULL,
    content    TEXT COLLATE utf8_unicode_ci,
    created_by CHAR(36)     NOT NULL,
    updated_by CHAR(36)     NOT NULL,
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,
    PRIMARY KEY (id),
    UNIQUE uq_title(product_id, title),
    FOREIGN KEY fk_product_notes_product_id(product_id)
        REFERENCES products(id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE projects
(
    id                  CHAR(36)    NOT NULL,
    product_id          CHAR(36)    NOT NULL,
    group_id            CHAR(36)    NOT NULL,
    key_name            VARCHAR(10) NOT NULL,
    name                VARCHAR(80) NOT NULL,
    leader_id           CHAR(36)    NOT NULL,
    default_assignee_id CHAR(36)    NOT NULL,
    trashed_at          DATETIME DEFAULT NULL,
    created_at          DATETIME    NOT NULL,
    updated_at          DATETIME    NOT NULL,
    PRIMARY KEY (id),
    UNIQUE uq_key_name(group_id, key_name),
    UNIQUE uq_name(group_id, name),
    FOREIGN KEY fk_projects_product_id(product_id)
        REFERENCES products(id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
);

CREATE TABLE project_notes
(
    id         CHAR(36)     NOT NULL,
    project_id CHAR(36)     NOT NULL,
    group_id   CHAR(36)     NOT NULL,
    title      VARCHAR(255) NOT NULL,
    content    TEXT COLLATE utf8_unicode_ci,
    created_by CHAR(36)     NOT NULL,
    updated_by CHAR(36)     NOT NULL,
    created_at DATETIME     NOT NULL,
    updated_at DATETIME     NOT NULL,
    PRIMARY KEY (id),
    UNIQUE uq_title(project_id, title),
    FOREIGN KEY fk_project_notes_project_id(project_id)
        REFERENCES projects(id)
        ON DELETE RESTRICT ON UPDATE RESTRICT
);