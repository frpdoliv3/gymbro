import json
from dataclasses import dataclass
from pathlib import Path
from typing import Self


@dataclass(frozen=True)
class Exercise:
    """Represents a physical exercise with its metadata and images.

    All string fields are automatically converted to title case on
    initialization. Optional fields (force, mechanic, equipment) may
    be None if not applicable to the exercise.

    Attributes:
        source_id: Unique identifier from the source dataset.
        name: Display name of the exercise.
        force: Type of force involved (e.g. Push, Pull).
        level: Difficulty level (e.g. Beginner, Intermediate).
        mechanic: Movement mechanic (e.g. Compound, Isolation).
        equipment: Required equipment (e.g. Barbell, Dumbbell).
        primary_muscles: Muscles primarily targeted by the exercise.
        secondary_muscles: Muscles secondarily engaged.
        instructions: Step-by-step instructions to perform the exercise.
        category: Categories the exercise belongs to.
        images: Raw image data associated with the exercise.
        id: Database identifier, defaults to 0 when not yet persisted.

    """

    source_id: str
    name: str
    force: str
    level: str
    mechanic: str
    equipment: str
    primary_muscles: list[str]
    secondary_muscles: list[str]
    instructions: list[str]
    category: list[str]
    images: list[bytes]
    id: int = 0

    def __post_init__(self) -> None:
        """Process strings to transform all to title."""
        object.__setattr__(
            self,
            "force",
            self.force.title() if self.force else None,
        )
        object.__setattr__(self, "level", self.level.title())
        object.__setattr__(
            self,
            "mechanic",
            self.mechanic.title() if self.mechanic else None,
        )
        object.__setattr__(
            self,
            "equipment",
            self.equipment.title() if self.equipment else None,
        )
        object.__setattr__(
            self,
            "primary_muscles",
            [x.title() for x in self.primary_muscles],
        )
        object.__setattr__(
            self,
            "secondary_muscles",
            [x.title() for x in self.secondary_muscles],
        )
        object.__setattr__(
            self,
            "instructions",
            [x.title() for x in self.instructions],
        )
        object.__setattr__(
            self,
            "category",
            [x.title() for x in self.category],
        )

    @staticmethod
    def create_from_folder(exercise_folder: str | Path) -> list[Exercise]:
        """Find and creates all exercises from a directory.

        It looks for all .json files in the given folder and assumes the
        image folder shares the same name as the JSON file (without extension).

        Args:
            exercise_folder: The directory to scan for exercise files.

        Returns:
            A list of initialized Exercise objects.

        """
        exercise_folder = Path(exercise_folder)
        if not exercise_folder.is_dir():
            raise NotADirectoryError(exercise_folder)

        return [
            Exercise.create_from_metadata(exercise)
            for exercise in exercise_folder.glob("*.json")
        ]

    @classmethod
    def create_from_metadata(cls, metadata_file_path: Path) -> Self:
        """Create an instance from a metadata JSON file.

        Reads the metadata file and loads the associated images relative
        to the metadata file's directory.

        Args:
            metadata_file_path: Path to the JSON metadata file.

        Returns:
            A new instance populated with data from the metadata file.

        Raises:
            FileNotFoundError: If the metadata file or any referenced
                image does not exist.
            KeyError: If a required field is missing from the metadata.
            json.JSONDecodeError: If the metadata file is not valid JSON.

        """
        with Path.open(metadata_file_path, "r") as metadata_file:
            metadata = json.loads(metadata_file.read())

        image_file_path: list[Path] = [
            metadata_file_path.parent / image_path for image_path in metadata["images"]
        ]
        return cls(
            source_id=metadata["id"],
            name=metadata["name"],
            force=metadata["force"],
            level=metadata["level"],
            mechanic=metadata["mechanic"],
            equipment=metadata["equipment"],
            primary_muscles=metadata["primaryMuscles"],
            secondary_muscles=metadata["secondaryMuscles"],
            instructions=metadata["instructions"],
            category=metadata["category"],
            images=[image_path.read_bytes() for image_path in image_file_path],
        )
