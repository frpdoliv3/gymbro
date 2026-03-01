package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

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
    val exercise_id: Int,
    val muscle_id: Int,
    val muscle_type: String // 'primary' or 'secondary'
)