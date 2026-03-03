package io.github.frpdoliv3.gymbro.composeApp.data.repository.plan

import io.github.frpdoliv3.gymbro.composeApp.data.local.plan.dao.PlanDao
import io.github.frpdoliv3.gymbro.composeApp.data.local.plan.dao.PlanExerciseDao
import io.github.frpdoliv3.gymbro.composeApp.data.local.plan.entity.PlanEntity
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseRef
import io.github.frpdoliv3.gymbro.composeApp.domain.model.PlanEntry
import kotlinx.coroutines.flow.Flow
import kotlinx.coroutines.flow.map

private const val DEFAULT_PLAN_ID = 1L
private const val DEFAULT_PLAN_NAME = "Workout Plan"

class PlanRepositoryImpl(
    private val planDao: PlanDao,
    private val planExerciseDao: PlanExerciseDao
) : PlanRepository {
    override fun observePlanEntries(): Flow<List<PlanEntry>> {
        return planExerciseDao.observePlanExercises(DEFAULT_PLAN_ID).map { entries ->
            entries.map {
                PlanEntry(
                    id = it.id,
                    position = it.position,
                    exerciseRef = ExerciseRef(
                        providerId = it.sourceProviderId,
                        exerciseId = it.exerciseSourceId
                    )
                )
            }
        }
    }

    override suspend fun addExercise(exerciseRef: ExerciseRef) {
        ensureDefaultPlan()
        planExerciseDao.appendPlanExercise(
            planId = DEFAULT_PLAN_ID,
            sourceProviderId = exerciseRef.providerId,
            exerciseSourceId = exerciseRef.exerciseId
        )
    }

    override suspend fun removeExercise(planEntryId: Long) {
        ensureDefaultPlan()
        planExerciseDao.removeAndReorder(
            planId = DEFAULT_PLAN_ID,
            id = planEntryId
        )
    }

    override suspend fun moveExerciseUp(planEntryId: Long) {
        ensureDefaultPlan()
        planExerciseDao.move(
            planId = DEFAULT_PLAN_ID,
            id = planEntryId,
            offset = -1
        )
    }

    override suspend fun moveExerciseDown(planEntryId: Long) {
        ensureDefaultPlan()
        planExerciseDao.move(
            planId = DEFAULT_PLAN_ID,
            id = planEntryId,
            offset = 1
        )
    }

    private suspend fun ensureDefaultPlan() {
        if (planDao.getPlanById(DEFAULT_PLAN_ID) == null) {
            planDao.insertPlan(
                PlanEntity(
                    id = DEFAULT_PLAN_ID,
                    name = DEFAULT_PLAN_NAME
                )
            )
        }
    }
}
