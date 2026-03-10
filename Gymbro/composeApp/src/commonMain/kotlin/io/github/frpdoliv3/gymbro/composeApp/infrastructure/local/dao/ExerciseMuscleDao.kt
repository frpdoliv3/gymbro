package io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.dao

import androidx.room.*
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.entity.ExerciseMuscleEntity
import kotlinx.coroutines.flow.Flow

@Dao
interface ExerciseMuscleDao {
    @Query("SELECT * FROM exercise_muscles WHERE exercise_id = :exerciseId")
    fun getMusclesByExerciseId(exerciseId: Int): Flow<List<ExerciseMuscleEntity>>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertExerciseMuscle(exerciseMuscle: ExerciseMuscleEntity)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertExerciseMuscles(exerciseMuscles: List<ExerciseMuscleEntity>)

    @Delete
    suspend fun deleteExerciseMuscle(exerciseMuscle: ExerciseMuscleEntity)
}
