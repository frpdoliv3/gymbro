package io.github.frpdoliv3.gymbro.composeApp.domain.model

data class ExerciseSummary(
    val ref: ExerciseRef,
    val name: String,
    val force: String?,
    val level: String,
    val mechanic: String?,
    val equipment: String?
)
