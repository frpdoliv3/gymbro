package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.ForeignKey

@Entity(
    tableName = "exercise_muscles",
    primaryKeys = ["exercise_id", "muscle_id", "muscle_type"],
    foreignKeys = [
        ForeignKey(
            entity = ExerciseEntity::class,
            parentColumns = arrayOf("id"),
            childColumns = arrayOf("exercise_id")
        ),
        ForeignKey(
            entity = MuscleEntity::class,
            parentColumns = arrayOf("id"),
            childColumns = arrayOf("muscle_id")
        )
    ]
)
data class ExerciseMuscleEntity(
    @ColumnInfo(name = "exercise_id")
    val exerciseId: Int,
    @ColumnInfo(name = "muscle_id")
    val muscleId: Int,
    @ColumnInfo(name = "muscle_type")
    val muscleType: String // 'primary' or 'secondary'
)