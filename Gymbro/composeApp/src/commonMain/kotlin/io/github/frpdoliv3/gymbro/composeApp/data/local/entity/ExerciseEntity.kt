package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "exercises")
data class ExerciseEntity(
    @PrimaryKey
    val id: Int,
    @ColumnInfo(name = "source_id")
    val sourceId: String,
    val name: String,
    val force: String?,
    val level: String,
    val mechanic: String?,
    val equipment: String?
)