package io.github.frpdoliv3.gymbro.composeApp.data.repository.plan

import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseRef
import io.github.frpdoliv3.gymbro.composeApp.domain.model.PlanEntry
import kotlinx.coroutines.flow.Flow

interface PlanRepository {
    fun observePlanEntries(): Flow<List<PlanEntry>>
    suspend fun addExercise(exerciseRef: ExerciseRef)
    suspend fun removeExercise(planEntryId: Long)
    suspend fun moveExerciseUp(planEntryId: Long)
    suspend fun moveExerciseDown(planEntryId: Long)
}
