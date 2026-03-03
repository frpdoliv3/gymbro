package io.github.frpdoliv3.gymbro.composeApp.domain.model

data class PlannedExercise(
    val id: Long,
    val position: Int,
    val exercise: ExerciseSummary
)
