package io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.dao

import androidx.room.*
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.entity.ExerciseCategoryEntity
import kotlinx.coroutines.flow.Flow

@Dao
interface ExerciseCategoryDao {
    @Query("SELECT * FROM exercise_categories WHERE exercise_id = :exerciseId")
    fun getCategoriesByExerciseId(exerciseId: Int): Flow<List<ExerciseCategoryEntity>>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertExerciseCategory(exerciseCategory: ExerciseCategoryEntity)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertExerciseCategories(exerciseCategories: List<ExerciseCategoryEntity>)

    @Delete
    suspend fun deleteExerciseCategory(exerciseCategory: ExerciseCategoryEntity)
}