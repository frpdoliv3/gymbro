import argparse
from pathlib import Path

import exercise


def main() -> None:
    """Handle the setup of command line arguments parser."""
    parser = argparse.ArgumentParser(
        prog="ExerciseETL",
        description="Extracts exercises from a folder, transforms into SQL rows and loads them into a SQLite database",
        epilog="The exercise folder must come from the free exercise db repository (https://github.com/yuhonas/free-exercise-db)",
    )
    parser.add_argument("exercise_folder_path")
    args = parser.parse_args()

    exercise_dir_path = Path(args.exercise_folder_path)

    if not exercise_dir_path.is_dir():
        err = f"Path {args.exercise_folder_path} does not point to a valid directory"
        raise NotADirectoryError(err)

    exercises = exercise.Exercise.create(exercise_dir_path)

    print(exercises)


if __name__ == "__main__":
    main()
