package io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.dao

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.entity.PlanEntity

@Dao
interface PlanDao {
    @Insert(onConflict = OnConflictStrategy.IGNORE)
    suspend fun insertPlan(plan: PlanEntity)

    @Query("SELECT * FROM plans WHERE id = :planId")
    suspend fun getPlanById(planId: Long): PlanEntity?
}
