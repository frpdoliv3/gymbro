package io.github.frpdoliv3.gymbro.composeApp.data.local.dao

import androidx.room.*
import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.ExerciseEntity
import kotlinx.coroutines.flow.Flow

@Dao
interface ExerciseDao {
    @Query("SELECT * FROM exercises ORDER BY name ASC")
    fun getAllExercises(): Flow<List<ExerciseEntity>>

    @Query("SELECT * FROM exercises WHERE id = :id")
    suspend fun getExerciseById(id: Int): ExerciseEntity?

    @Query("SELECT * FROM exercises WHERE source_id = :sourceId")
    suspend fun getExerciseBySourceId(sourceId: String): ExerciseEntity?

    @Query("SELECT * FROM exercises WHERE name LIKE '%' || :name || '%' ORDER BY name ASC")
    fun searchExercisesByName(name: String): Flow<List<ExerciseEntity>>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertExercise(exercise: ExerciseEntity)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertExercises(exercises: List<ExerciseEntity>)

    @Delete
    suspend fun deleteExercise(exercise: ExerciseEntity)
}
