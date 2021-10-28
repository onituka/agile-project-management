USE test_db;

CREATE TABLE products
(
    id         CHAR(36)    NOT NULL,
    name       VARCHAR(80) NOT NULL,
    group_id   CHAR(36)    NOT NULL,
    leader_id  CHAR(36)    NOT NULL,
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
);

CREATE TABLE product_pages
(
    id         CHAR(36)    NOT NULL,
    title      VARCHAR(80) NOT NULL,
    content    TEXT        NOT NULL,
    author_id  CHAR(36)    NOT NULL,
    editor_id  CHAR(36)    NOT NULL,
    product_id CHAR(36)    NOT NULL,
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY fk_product_id (product_id)
        REFERENCES products (id)
        ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE product_page_comments
(
    id              CHAR(36) NOT NULL,
    user_id         CHAR(36) NOT NULL,
    content         TEXT     NOT NULL,
    product_page_id CHAR(36) NOT NULL,
    created_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at      DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    FOREIGN KEY fk_product_page_id (product_page_id)
        REFERENCES product_pages (id)
        ON DELETE RESTRICT ON UPDATE CASCADE
);

CREATE TABLE product_page_comment_paths
(
    comment_parent_id CHAR(36) NOT NULL,
    comment_child_id  CHAR(36) NOT NULL
);

CREATE TABLE projects
(
    id                  CHAR(36)    NOT NULL,
    group_id            CHAR(36)    NOT NULL,
    key_name            VARCHAR(10) NOT NULL,
    name                VARCHAR(80) NOT NULL,
    leader_id           CHAR(36)    NOT NULL,
    default_assignee_id CHAR(36)    NOT NULL,
    created_at          DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE uq_key_name(group_id, key_name),
    UNIQUE uq_name(group_id, name)
);
