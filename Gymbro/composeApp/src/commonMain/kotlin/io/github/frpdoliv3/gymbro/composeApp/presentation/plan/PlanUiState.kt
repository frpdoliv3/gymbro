package io.github.frpdoliv3.gymbro.composeApp.presentation.plan

import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseDetail
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseSummary
import io.github.frpdoliv3.gymbro.composeApp.domain.model.PlannedExercise

data class PlanUiState(
    val query: String = "",
    val selectedExercises: List<PlannedExercise> = emptyList(),
    val availableExercises: List<ExerciseSummary> = emptyList(),
    val isDetailVisible: Boolean = false,
    val detail: ExerciseDetail? = null,
    val isLoadingDetail: Boolean = false
)
