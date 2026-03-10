package io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.dao

import androidx.room.*
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.entity.ExerciseImageEntity
import kotlinx.coroutines.flow.Flow

@Dao
interface ExerciseImageDao {
    @Query("SELECT * FROM exercise_images WHERE exercise_id = :exerciseId ORDER BY image_order ASC")
    fun getImagesByExerciseId(exerciseId: Int): Flow<List<ExerciseImageEntity>>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertImage(image: ExerciseImageEntity)

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertImages(images: List<ExerciseImageEntity>)

    @Delete
    suspend fun deleteImage(image: ExerciseImageEntity)
}
