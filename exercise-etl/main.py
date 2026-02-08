import argparse
import os


def main():
    #  This section handles the setup of command line arguments parser
    parser = argparse.ArgumentParser(
        prog="ExerciseETL",
        description="Extracts exercises from a folder, transforms into SQL rows and loads them into a SQLite database",
        epilog="The exercise folder must come from the free exercise db repository (https://github.com/yuhonas/free-exercise-db)",
    )
    parser.add_argument("exercise_folder_path")
    args = parser.parse_args()

    if not os.path.isdir(args.exercise_folder_path):
        raise NotADirectoryError(
            f"Path {args.exercise_folder_path} does not point to a valid directory"
        )


if __name__ == "__main__":
    main()
