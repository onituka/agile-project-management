USE test_db;

-- products table test data
INSERT INTO products
(id, name, group_id, leader_id, created_at, updated_at)
values
    ("4495c574-34c2-4fb3-9ca4-3a7c79c267a6", "プロジェクト管理ツール", "124d78d6-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac184402", default, default),
    ("4487c574-34c2-4fb3-9ca4-3a7c79c267a6", "test", "224d78d6-1d03-11ec-a478-0242ac180002", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", default, default);

-- product_pages table test data
INSERT INTO product_pages
(id, title, content, author_id, editor_id, product_id, created_at, updated_at)
values
    ("50d891c2-f1fd-4291-a333-4a5815246442", "プロジェクト機能", "概要を記載します。", "d90db136-0885-4cb6-a06b-421e7e6f2198", "f147e5d9-8ad7-4008-8ee7-f17935728aef", "4495c574-34c2-4fb3-9ca4-3a7c79c267a6", default, default),
    ("60d811c2-f1fd-4291-a333-4a5815246442", "test", "プロダクトビジョンを記載します。", "d90db136-0885-4cb6-a06b-421e7e6f2198", "f147e5d9-8ad7-4008-8ee7-f17935728aef", "4487c574-34c2-4fb3-9ca4-3a7c79c267a6", default, default);

-- product_page_comments table test data
INSERT INTO product_page_comments
(id, user_id, content, product_page_id, created_at, updated_at)
values
    ("40d891c2-f1fd-4291-a333-4a5815246442", "573418ec-f6d2-49e8-81f5-b6e641dbf00a", "コメントです！", "50d891c2-f1fd-4291-a333-4a5815246442", default, default),
    ("40d811c2-f1fd-4291-a333-4a5815246442", "573418ec-f6d2-49e8-81f5-b6e641dbf00a", "test.", "60d811c2-f1fd-4291-a333-4a5815246442", default, default);

-- product_page_comment_paths table test data
INSERT INTO product_page_comment_paths
(comment_parent_id, comment_child_id)
values
    ("cf77d797-597d-40d1-91ca-171cd4aff5b9", "d0278096-04f0-47ec-aeca-b15fe3886b46"),
    ("cf77d797-597d-40d1-91ca-171cd4aff5b9", "d0278096-04f0-47ec-aeca-b15fe3886b46");

-- projects table test data
INSERT INTO projects
    (id, group_id, key_name, name, leader_id, default_assignee_id, created_at, updated_at)
values
    ("024d71d6-1d03-11ec-a478-0242ac180002", "024d78d6-1d03-11ec-a478-0242ac180002", "AJA", "プロジェクト管理ツール", "024d78d6-1d03-11ec-a478-0242ac184402", "024d78d6-1d03-11ec-a478-9242ac180002", default, default),
    ("024d7945-1d03-11ec-a478-0242ac180002", "023d76d6-1d03-11ec-a478-0242ac180002", "KANR", "kanry", "024d78d6-1d03-11ec-a478-0242ac180022", "024d78d6-1d03-11ec-a478-0242ac180502", default, default),
    ("024d61d4-1d03-11ec-a478-0242ac180002", "022d74d6-1d03-11ec-a478-0242ac180002", "DEF", "デフォルトツール", "024d78d6-1d03-11ec-a478-0245ac180002", "024d78d6-1d03-11ec-a488-0242ac180002", default, default),
    ("024d81d3-1d03-11ec-a478-0242ac180002", "021d72d6-1d03-11ec-a478-0242ac180002", "PIED", "PiedPiper", "024d78d6-1d03-11ec-a478-0242ac380002", "024d78d6-1d03-11ec-a478-0842ac180002", default, default),
    ("024d31d2-1d03-11ec-a478-0242ac180002", "025d71d6-1d03-11ec-a478-0242ac180002", "NEO", "NeoPiedPiper", "024d78d6-1d03-11ec-a478-0246ac180002", "024d78d6-1d03-11ec-a498-0242ac180002", default, default);
