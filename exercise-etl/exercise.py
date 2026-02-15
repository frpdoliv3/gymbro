from pathlib import Path


class Exercise:
    """Represents an exercise with associated metadata and image resources.

    This class manages the relationship between a JSON metadata file and an
    optional directory containing exercise execution images.
    """

    def __init__(self, metadata_file: Path, image_folder: Path | None) -> None:
        """Initialize an Exercise instance.

        Args:
            metadata_file: Path to the JSON file containing exercise data.
            image_folder: Path to the directory containing exercise images.

        Raises:
            NotADirectoryError: If image_folder is provided but isn't a directory.
            FileNotFoundError: If the metadata_file does not exist.

        """
        if not (image_folder is None or image_folder.is_dir()):
            err = f"Path {image_folder} does not point to a valid directory"
            raise NotADirectoryError(err)

        if not metadata_file.is_file():
            err = f"Metadata file not found: {metadata_file}"
            raise FileNotFoundError(err)

        self.__metadata_file = metadata_file
        self.__image_folder = image_folder

    def __str__(self) -> str:
        """Return a string summary of the exercise paths."""
        return f"Exercise: {self.__metadata_file} images at {self.__image_folder}"

    def __repr__(self) -> str:
        """Return a string that can recreate the instance."""
        return (
            f"Exercise(metadata_file={self.__metadata_file!r}, "
            f"image_folder={self.__image_folder!r})"
        )

    @classmethod
    def create(cls, exercise_folder: str | Path) -> list[Exercise]:
        """Find and creates all exercises from a directory.

        It looks for all .json files in the given folder and assumes the
        image folder shares the same name as the JSON file (without extension).

        Args:
            exercise_folder: The directory to scan for exercise files.

        Returns:
            A list of initialized Exercise objects.

        """
        exercise_folder = Path(exercise_folder)

        return [
            cls(metadata_file, metadata_file.parent / metadata_file.stem)
            for metadata_file in exercise_folder.glob("*.json")
        ]
