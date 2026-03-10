package io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.entity

import androidx.room.Entity
import androidx.room.PrimaryKey

@Entity(tableName = "plans")
data class PlanEntity(
    @PrimaryKey
    val id: Long,
    val name: String
)
