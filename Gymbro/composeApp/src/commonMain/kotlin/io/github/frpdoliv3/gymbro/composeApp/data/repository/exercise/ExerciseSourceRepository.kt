package io.github.frpdoliv3.gymbro.composeApp.data.repository.exercise

import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseDetail
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseRef
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseSummary
import kotlinx.coroutines.flow.Flow

interface ExerciseSourceRepository {
    fun observeAllExercises(): Flow<List<ExerciseSummary>>
    fun searchExercises(query: String): Flow<List<ExerciseSummary>>
    suspend fun getExerciseDetail(ref: ExerciseRef): ExerciseDetail?
}
