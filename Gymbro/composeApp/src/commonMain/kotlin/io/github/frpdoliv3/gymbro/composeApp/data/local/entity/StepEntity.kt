package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.ForeignKey
import androidx.room.PrimaryKey

@Entity(
    tableName = "steps",
    foreignKeys = [
        ForeignKey(
            entity = ExerciseEntity::class,
            parentColumns = arrayOf("id"),
            childColumns = arrayOf("exercise_id")
        )
    ]
)
data class StepEntity(
    @PrimaryKey
    val id: Int,
    @ColumnInfo(name = "exercise_id")
    val exerciseId: Int,
    val description: String,
    @ColumnInfo(name = "step_order")
    val stepOrder: Int
)
