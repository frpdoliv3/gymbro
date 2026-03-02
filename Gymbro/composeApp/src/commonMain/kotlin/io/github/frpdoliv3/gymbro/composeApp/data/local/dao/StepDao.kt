package io.github.frpdoliv3.gymbro.composeApp.data.local.dao

import androidx.room.*
import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.StepEntity
import kotlinx.coroutines.flow.Flow

@Dao
interface StepDao {
    @Query("SELECT * FROM steps WHERE exercise_id = :exerciseId ORDER BY step_order ASC")
    fun getStepsByExerciseId(exerciseId: Int): Flow<List<StepEntity>>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertStep(step: StepEntity)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertSteps(steps: List<StepEntity>)

    @Delete
    suspend fun deleteStep(step: StepEntity)
}
