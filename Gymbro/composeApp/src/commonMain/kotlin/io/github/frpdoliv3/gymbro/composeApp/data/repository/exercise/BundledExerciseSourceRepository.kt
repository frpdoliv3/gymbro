package io.github.frpdoliv3.gymbro.composeApp.data.repository.exercise

import io.github.frpdoliv3.gymbro.composeApp.data.local.dao.ExerciseDao
import io.github.frpdoliv3.gymbro.composeApp.data.local.dao.ExerciseImageDao
import io.github.frpdoliv3.gymbro.composeApp.data.local.dao.StepDao
import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.ExerciseEntity
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseDetail
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseImage
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseRef
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseSummary
import kotlinx.coroutines.flow.first
import kotlinx.coroutines.flow.map

private const val BUNDLED_PROVIDER_ID = "bundled"
private const val LOCAL_ID_PREFIX = "local:"

class BundledExerciseSourceRepository(
    private val exerciseDao: ExerciseDao,
    private val stepDao: StepDao,
    private val exerciseImageDao: ExerciseImageDao
) : ExerciseSourceRepository {
    override fun observeAllExercises() = exerciseDao.getAllExercises().map { exercises ->
        exercises.map { exercise -> exercise.toSummary() }
    }

    override fun searchExercises(query: String) = exerciseDao.searchExercisesByName(query.trim()).map { exercises ->
        exercises.map { exercise -> exercise.toSummary() }
    }

    override suspend fun getExerciseDetail(ref: ExerciseRef): ExerciseDetail? {
        if (ref.providerId != BUNDLED_PROVIDER_ID) return null

        val entity = when {
            ref.exerciseId.startsWith(LOCAL_ID_PREFIX) -> {
                ref.exerciseId.removePrefix(LOCAL_ID_PREFIX).toIntOrNull()?.let { id ->
                    exerciseDao.getExerciseById(id)
                }
            }

            else -> exerciseDao.getExerciseBySourceId(ref.exerciseId)
        } ?: return null

        val steps = stepDao.getStepsByExerciseId(entity.id).first().map { it.description }
        val images = exerciseImageDao.getImagesByExerciseId(entity.id).first().map {
            ExerciseImage(
                mimeType = it.mimeType,
                bytes = it.imageBlob
            )
        }

        return ExerciseDetail(
            summary = entity.toSummary(),
            instructions = steps,
            images = images
        )
    }

    private fun ExerciseEntity.toSummary(): ExerciseSummary {
        return ExerciseSummary(
            ref = ExerciseRef(
                providerId = BUNDLED_PROVIDER_ID,
                exerciseId = sourceId ?: "$LOCAL_ID_PREFIX$id"
            ),
            name = name,
            force = force,
            level = level,
            mechanic = mechanic,
            equipment = equipment
        )
    }
}
