package io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.dao

import androidx.room.Dao
import androidx.room.Insert
import androidx.room.OnConflictStrategy
import androidx.room.Query
import androidx.room.Transaction
import androidx.room.Update
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.entity.PlanExerciseEntity
import kotlinx.coroutines.flow.Flow

@Dao
interface PlanExerciseDao {
    @Query("SELECT * FROM plan_exercises WHERE plan_id = :planId ORDER BY position ASC")
    fun observePlanExercises(planId: Long): Flow<List<PlanExerciseEntity>>

    @Query("SELECT * FROM plan_exercises WHERE plan_id = :planId ORDER BY position ASC")
    suspend fun getPlanExercises(planId: Long): List<PlanExerciseEntity>

    @Insert(onConflict = OnConflictStrategy.REPLACE)
    suspend fun insertPlanExercise(planExercise: PlanExerciseEntity): Long

    @Update
    suspend fun updatePlanExercises(planExercises: List<PlanExerciseEntity>)

    @Query("DELETE FROM plan_exercises WHERE id = :id")
    suspend fun deleteById(id: Long)

    @Transaction
    suspend fun appendPlanExercise(planId: Long, sourceProviderId: String, exerciseSourceId: String): Long {
        val nextPosition = getPlanExercises(planId).size
        return insertPlanExercise(
            PlanExerciseEntity(
                planId = planId,
                sourceProviderId = sourceProviderId,
                exerciseSourceId = exerciseSourceId,
                position = nextPosition
            )
        )
    }

    @Transaction
    suspend fun removeAndReorder(planId: Long, id: Long) {
        val reordered = getPlanExercises(planId)
            .filterNot { it.id == id }
            .mapIndexed { index, entry -> entry.copy(position = index) }

        deleteById(id)
        if (reordered.isNotEmpty()) {
            updatePlanExercises(reordered)
        }
    }

    @Transaction
    suspend fun move(planId: Long, id: Long, offset: Int) {
        val current = getPlanExercises(planId)
        val fromIndex = current.indexOfFirst { it.id == id }
        if (fromIndex == -1) return

        val targetIndex = fromIndex + offset
        if (targetIndex !in current.indices) return

        val mutable = current.toMutableList()
        val moved = mutable.removeAt(fromIndex)
        mutable.add(targetIndex, moved)
        val reordered = mutable.mapIndexed { index, entry -> entry.copy(position = index) }
        updatePlanExercises(reordered)
    }
}
