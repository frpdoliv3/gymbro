package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.PrimaryKey
import androidx.room.ForeignKey

@Entity(
    tableName = "steps",
    foreignKeys = [
        ForeignKey(
            entity = ExerciseEntity::class,
            parentColumns = arrayOf("id"),
            childColumns = arrayOf("exercise_id"),
            onDelete = ForeignKey.CASCADE
        )
    ]
)
data class StepEntity(
    @PrimaryKey
    val id: Int,
    @ColumnInfo(name = "exercise_id")
    val exerciseId: Int,
    val description: String
)