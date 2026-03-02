package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

import androidx.room.Entity
import androidx.room.Index
import androidx.room.PrimaryKey

@Entity(
    tableName = "muscles",
    indices = [
        Index(value = ["name"], unique = true)
    ]
)
data class MuscleEntity(
    @PrimaryKey
    val id: Int,
    val name: String
)
