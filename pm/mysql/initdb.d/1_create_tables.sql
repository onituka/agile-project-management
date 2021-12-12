USE test_db;

CREATE TABLE products
(
    id         CHAR(36)    NOT NULL,
    group_id   CHAR(36)    NOT NULL,
    name       VARCHAR(80) NOT NULL,
    leader_id  CHAR(36)    NOT NULL,
    created_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id)
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
    created_at          DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at          DATETIME    NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id),
    UNIQUE uq_key_name(group_id, key_name),
    UNIQUE uq_name(group_id, name),
    FOREIGN KEY fk_product_id(product_id)
    REFERENCES products(id)
    ON DELETE RESTRICT ON UPDATE RESTRICT
);
