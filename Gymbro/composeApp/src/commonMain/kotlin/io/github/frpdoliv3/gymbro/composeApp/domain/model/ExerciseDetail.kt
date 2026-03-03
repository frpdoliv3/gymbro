package io.github.frpdoliv3.gymbro.composeApp.domain.model

data class ExerciseDetail(
    val summary: ExerciseSummary,
    val instructions: List<String>,
    val images: List<ExerciseImage>
)
