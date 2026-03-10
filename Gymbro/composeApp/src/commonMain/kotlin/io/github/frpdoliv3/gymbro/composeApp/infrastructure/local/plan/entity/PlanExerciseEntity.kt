package io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.entity

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.ForeignKey
import androidx.room.Index
import androidx.room.PrimaryKey

@Entity(
    tableName = "plan_exercises",
    foreignKeys = [
        ForeignKey(
            entity = PlanEntity::class,
            parentColumns = ["id"],
            childColumns = ["plan_id"],
            onDelete = ForeignKey.CASCADE
        )
    ],
    indices = [
        Index(value = ["plan_id", "position"], unique = true),
        Index(value = ["plan_id"])
    ]
)
data class PlanExerciseEntity(
    @PrimaryKey(autoGenerate = true)
    val id: Long = 0,
    @ColumnInfo(name = "plan_id")
    val planId: Long,
    @ColumnInfo(name = "source_provider_id")
    val sourceProviderId: String,
    @ColumnInfo(name = "exercise_source_id")
    val exerciseSourceId: String,
    val position: Int
)
