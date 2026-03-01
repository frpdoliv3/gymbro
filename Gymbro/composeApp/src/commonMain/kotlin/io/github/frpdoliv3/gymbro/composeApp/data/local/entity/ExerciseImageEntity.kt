package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

import androidx.room.Entity
import androidx.room.PrimaryKey
import androidx.room.ForeignKey

@Entity(
    tableName = "exercise_images",
    foreignKeys = [
        ForeignKey(
            entity = ExerciseEntity::class,
            parentColumns = arrayOf("id"),
            childColumns = arrayOf("exercise_id"),
            onDelete = ForeignKey.CASCADE
        )
    ]
)
data class ExerciseImageEntity(
    @PrimaryKey(autoGenerate = true)
    val id: Int,
    val exercise_id: Int,
    val image_order: Int,
    val image_blob: ByteArray,
    val mime_type: String
)