PRAGMA foreign_keys = ON;
PRAGMA encoding = "UTF-8";

-- Delete Tables
DROP TABLE IF EXISTS exercise_categories;
DROP TABLE IF EXISTS exercise_muscles;
DROP TABLE IF EXISTS exercise_images;
DROP TABLE IF EXISTS steps;
DROP TABLE IF EXISTS muscles;
DROP TABLE IF EXISTS categories;
DROP TABLE IF EXISTS exercises;

-- Create Tables
CREATE TABLE exercises(
    id INTEGER NOT NULL PRIMARY KEY,
    source_id TEXT UNIQUE,
    name TEXT NOT NULL,
    force TEXT,
    level TEXT NOT NULL,
    mechanic TEXT,
    equipment TEXT
);

CREATE TABLE categories(
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE muscles(
    id INTEGER NOT NULL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE steps(
    id INTEGER NOT NULL PRIMARY KEY,
    exercise_id INTEGER NOT NULL REFERENCES exercises(id),
    description TEXT NOT NULL,
    step_order INTEGER NOT NULL
);

CREATE TABLE exercise_images(
    id INTEGER NOT NULL PRIMARY KEY,
    exercise_id INTEGER NOT NULL REFERENCES exercises(id),
    image_order INTEGER NOT NULL,
    image_blob BLOB NOT NULL,
    mime_type TEXT NOT NULL,
    UNIQUE(exercise_id, image_order)
);

CREATE TABLE exercise_muscles(
    exercise_id INTEGER NOT NULL REFERENCES exercises(id),
    muscle_id INTEGER NOT NULL REFERENCES muscles(id),
    muscle_type TEXT NOT NULL CHECK(muscle_type IN ('primary', 'secondary')),
    PRIMARY KEY(exercise_id, muscle_id, muscle_type)
);

CREATE TABLE exercise_categories(
    exercise_id INTEGER NOT NULL REFERENCES exercises(id),
    category_id INTEGER NOT NULL REFERENCES categories(id),
    PRIMARY KEY(exercise_id, category_id)
);

CREATE INDEX index_steps_exercise_id ON steps(exercise_id);
CREATE INDEX index_exercise_images_exercise_id ON exercise_images(exercise_id);
CREATE INDEX index_exercise_muscles_muscle_id ON exercise_muscles(muscle_id);
CREATE INDEX index_exercise_categories_category_id ON exercise_categories(category_id);
