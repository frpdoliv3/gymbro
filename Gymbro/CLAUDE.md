# Project: Gymbro
An application to track progress in the gym

# Technologies
- Kotlin Multiplatform WITHOUT SHARED UI

# Architecture
- composeApp: Holds all the application logic from the infrastructure layer to the ViewModels
- androidApp: Contains the connection with the ViewModels and UI written in Jetpack Compose

# Important Notes
- NEVER ask me to run Gradle i should run the application by hand UNLESS YOU WANT TO CHECK COMPILATION ERRORS
- NEVER INCLUDE UI CODE IN composeApp