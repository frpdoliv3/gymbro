# General Instructions
- Always write shell commands relative to the current working directory, and never use parent-directory traversal such as ../ or otherwise navigate upward in the filesystem. The only exception is a command explicitly written in this AGENTS.md file.
- Keep exactly one committed Room schema JSON file for the current database version, and ensure it stays aligned with exercise-etl/db/create.sql.

# Project: exercise-etl
A simple go application that parses exercises inside sources/free-exercise-db into a SQLite database

## Technologies
- Go Programming language
- SQLite Database

## Important notes
- The updated schema is stored inside exercise-etl/db/create.sql
- The database should be stored in exercise-etl/db/exercises.db and then copied to Gymbro/composeApp/src/commonMain/resources/databases/exercises.db
- To create a new version of the database from the repository root run ```go run ./exercise-etl --override --schema exercise-etl/db/create.sql ../sources/free-exercise-db/exercises exercise-etl/db/exercises.db```
- The exercise_images table contains the images stored as binary inside image_blob and the format of the blob is determined by mime_type

# Project: Gymbro
An application to track progress in the gym this application is stored in the folder Gymbro/

Features are described inside Gymbro/spec/ in markdown files

## Technologies
- Kotlin Multiplatform WITHOUT SHARED UI
- Room for database mapping in all platforms
- Koin for dependency injection
- Android Gradle Plugin 9 (attention this is new the data you have in the model might be outdated)

## Commands
Builds the project
```./gradlew build```

## Architecture
- composeApp: Holds all the application logic from the infrastructure layer to the ViewModels
- androidApp: Contains the connection with the ViewModels and UI written in Jetpack Compose
- database: The main exercise database of the application is stored inside Gymbro/composeApp/src/commonMain/resources/databases/exercises.db. This database contains the exercises that will be shown in the app 

## Important Notes
- ALWAYS RUN THE CODE WITH THE COMMAND THAT BUILDS THE PROJECT THAT IS REFERENCED IN THIS FILE BEFORE FINISHING AND SAYING THAT YOU WERE SUCCESSFUL
- NEVER INCLUDE UI CODE IN composeApp
- The bundled `Gymbro/composeApp/src/commonMain/resources/databases/exercises.db` is a static, app-agnostic exercise source.
- Store Gymbro application data in a separate app database, not in `exercises.db`.
- If a feature needs new persistence, create or extend the app database instead of changing the bundled exercise source database.
- Keep exercise access generic enough to support future plugin-based exercise providers.
