package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.ForeignKey

@Entity(
    tableName = "exercise_categories",
    primaryKeys = ["exercise_id", "category_id"],
    foreignKeys = [
        ForeignKey(
            entity = ExerciseEntity::class,
            parentColumns = arrayOf("id"),
            childColumns = arrayOf("exercise_id")
        ),
        ForeignKey(
            entity = CategoryEntity::class,
            parentColumns = arrayOf("id"),
            childColumns = arrayOf("category_id")
        )
    ]
)
data class ExerciseCategoryEntity(
    @ColumnInfo(name = "exercise_id")
    val exerciseId: Int,
    @ColumnInfo(name = "category_id")
    val categoryId: Int
)