package io.github.frpdoliv3.gymbro.composeApp.data.local.dao

import androidx.room.*
import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.MuscleEntity
import kotlinx.coroutines.flow.Flow

@Dao
interface MuscleDao {
    @Query("SELECT * FROM muscles")
    fun getAllMuscles(): Flow<List<MuscleEntity>>

    @Query("SELECT * FROM muscles WHERE id = :id")
    suspend fun getMuscleById(id: Int): MuscleEntity?

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertMuscle(muscle: MuscleEntity)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertMuscles(muscles: List<MuscleEntity>)
}