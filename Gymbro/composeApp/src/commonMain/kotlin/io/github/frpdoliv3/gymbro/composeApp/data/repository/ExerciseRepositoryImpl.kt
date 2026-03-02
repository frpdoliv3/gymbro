package io.github.frpdoliv3.gymbro.composeApp.data.repository

import io.github.frpdoliv3.gymbro.composeApp.data.local.dao.ExerciseDao
import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.ExerciseEntity
import kotlinx.coroutines.flow.Flow

class ExerciseRepositoryImpl(
    private val exerciseDao: ExerciseDao
) : ExerciseRepository {
    override fun getAllExercises(): Flow<List<ExerciseEntity>> {
        return exerciseDao.getAllExercises()
    }

    override fun searchExercisesByName(name: String): Flow<List<ExerciseEntity>> {
        return exerciseDao.searchExercisesByName(name)
    }

    override suspend fun getExerciseById(id: Int): ExerciseEntity? {
        return exerciseDao.getExerciseById(id)
    }

    override suspend fun getExerciseBySourceId(sourceId: String): ExerciseEntity? {
        return exerciseDao.getExerciseBySourceId(sourceId)
    }
}
