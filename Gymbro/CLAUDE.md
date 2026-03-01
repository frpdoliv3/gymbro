# Project: Gymbro
An application to track progress in the gym

# Technologies
- Kotlin Multiplatform WITHOUT SHARED UI

# Commands
Builds the project
```./gradlew build```

# Architecture
- composeApp: Holds all the application logic from the infrastructure layer to the ViewModels
- androidApp: Contains the connection with the ViewModels and UI written in Jetpack Compose

# Important Notes
- ALWAYS RUN THE CODE WITH THE COMMAND THAT BUILDS THE PROJECT THAT IS REFERENCED IN THIS FILE BEFORE FINISHING AND SAYING THAT YOU WERE SUCCESSFUL
- NEVER INCLUDE UI CODE IN composeApp