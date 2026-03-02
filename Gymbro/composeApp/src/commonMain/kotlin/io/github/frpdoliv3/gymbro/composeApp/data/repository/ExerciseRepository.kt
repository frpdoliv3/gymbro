package io.github.frpdoliv3.gymbro.composeApp.data.repository

import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.ExerciseEntity
import kotlinx.coroutines.flow.Flow

interface ExerciseRepository {
    fun getAllExercises(): Flow<List<ExerciseEntity>>
    fun searchExercisesByName(name: String): Flow<List<ExerciseEntity>>
    suspend fun getExerciseById(id: Int): ExerciseEntity?
    suspend fun getExerciseBySourceId(sourceId: String): ExerciseEntity?
}
