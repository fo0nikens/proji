--
-- DROP TABLES
--
DROP TABLE IF EXISTS project;
DROP TABLE IF EXISTS project_status;
DROP TABLE IF EXISTS project_class;
DROP TABLE IF EXISTS class_label;
DROP TABLE IF EXISTS class_folder;
DROP TABLE IF EXISTS class_file;
DROP TABLE IF EXISTS class_script;
--
-- CREATE TABLES
--
CREATE TABLE IF NOT EXISTS project(
  project_id INTEGER PRIMARY KEY,
  title TEXT NOT NULL,
  project_class_id INTEGER REFERENCES project_class(project_class_id),
  install_path TEXT,
  install_data TEXT,
  project_status_id INTEGER REFERENCES project_status(project_status_id)
);
--
CREATE TABLE IF NOT EXISTS project_status(
  project_status_id INTEGER PRIMARY KEY,
  project_status TEXT NOT NULL,
  comment TEXT
);
--
CREATE TABLE IF NOT EXISTS project_class(
  project_class_id INTEGER PRIMARY KEY,
  class_name TEXT NOT NULL
);
--
CREATE TABLE IF NOT EXISTS class_label(
  class_label_id INTEGER PRIMARY KEY,
  project_class_id INTEGER REFERENCES project_class(project_class_id),
  label TEXT NOT NULL
);
--
CREATE TABLE IF NOT EXISTS class_folder(
  class_folder_id INTEGER PRIMARY KEY,
  project_class_id INTEGER REFERENCES project_class(project_class_id),
  target_path TEXT NOT NULL,
  template_name TEXT
);
--
CREATE TABLE IF NOT EXISTS class_file(
  class_file_id INTEGER PRIMARY KEY,
  project_class_id INTEGER REFERENCES project_class(project_class_id),
  target_path TEXT NOT NULL,
  template_name TEXT
);
--
CREATE TABLE IF NOT EXISTS class_script(
  class_script_id INTEGER PRIMARY KEY,
  project_class_id INTEGER REFERENCES project_class(project_class_id),
  script_name TEXT NOT NULL,
  run_as_sudo INTEGER NOT NULL
);
--
-- INSERT INTO TABLES
--
INSERT INTO
  project_status(project_status, comment)
VALUES
  (
    "active",
    "I'm actively working on this project."
  ),
  (
    "inactive",
    "I stopped working on this project for now."
  ),
  (
    "done",
    "I finished this project. There is nothing to do."
  ),
  (
    "dead",
    "This project is dead. I'm not planning to work on it anytime soon."
  );
--
INSERT INTO
  project_class(class_name)
VALUES
  ("c-plus-plus"),
  ("c"),
  ("python"),
  ("web");
--
INSERT INTO
  class_label(project_class_id, label)
VALUES
  (1, "cpp"),
  (1, "c++"),
  (1, "cc"),
  (2, "c"),
  (3, "python"),
  (3, "py"),
  (4, "html"),
  (4, "php"),
  (4, "javascript"),
  (4, "js"),
  (4, "typescript"),
  (4, "ts"),
  (4, "web");
--
INSERT INTO
  class_folder(
    project_class_id,
    target_path,
    template_name
  )
VALUES
  (NULL, ".vscode", "vscode/"),
  (1, "build/debug/", NULL),
  (1, "build/release/", NULL),
  (1, "doc/", NULL),
  (1, "include/", NULL),
  (1, "lib/", NULL),
  (1, "src/", NULL),
  (1, "test/", NULL),
  (2, "build/debug/", NULL),
  (2, "build/release/", NULL),
  (2, "doc/", NULL),
  (2, "include/", NULL),
  (2, "lib/", NULL),
  (2, "src/", NULL),
  (2, "test/", NULL),
  (3, "src/", NULL),
  (3, "doc/", NULL),
  (3, "test/", NULL),
  (4, "public_html/css/", NULL),
  (4, "public_html/img/", NULL),
  (4, "public_html/js/", NULL),
  (4, "public_html/fonts/", NULL),
  (4, "public_html/include/", NULL),
  (4, "resources/library", NULL),
  (4, "resources/templates/", NULL);
--
INSERT INTO
  class_file(project_class_id, target_path, template_name)
VALUES
  (NULL, ".gitignore", "gitignore"),
  (1, "src/main.cpp", "main.cpp"),
  (
    1,
    "CMakeLists.txt",
    "CMakeLists-cpp.txt"
  ),
  (2, "src/main.c", "main.c"),
  (
    2,
    "CMakeLists.txt",
    "CMakeLists-c.txt"
  ),
  (3, "src/__main__.py", NULL),
  (3, "src/__init__.py", NULL),
  (
    4,
    "public_html/index.php",
    "index.php"
  ),
  (4, "public_html/css/main.css", NULL),
  (4, "public_html/css/util.css", NULL),
  (4, "public_html/js/main.js", NULL);
--
INSERT INTO
  class_script(project_class_id, script_name, run_as_sudo)
VALUES
  (NULL, "init_git.sh", 0),
  (3, "init_virtualenv.sh", 0);
DROP INDEX IF EXISTS class_label_idx;
DROP INDEX IF EXISTS project_class_idx;
DROP INDEX IF EXISTS class_folder_idx;
DROP INDEX IF EXISTS class_file_idx;
DROP INDEX IF EXISTS class_script_idx;
CREATE UNIQUE INDEX class_label_idx ON class_label(label);
CREATE UNIQUE INDEX project_class_idx ON project_class(class_name);
CREATE UNIQUE INDEX class_folder_idx ON class_folder(project_class_id, target_path, template_name);
CREATE UNIQUE INDEX class_file_idx ON class_file(project_class_id, target_path, template_name);
CREATE UNIQUE INDEX class_script_idx ON class_script(project_class_id, script_name, run_as_sudo);