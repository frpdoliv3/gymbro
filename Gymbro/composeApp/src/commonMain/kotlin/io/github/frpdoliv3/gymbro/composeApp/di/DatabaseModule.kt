package io.github.frpdoliv3.gymbro.composeApp.di

import io.github.frpdoliv3.gymbro.composeApp.data.repository.exercise.BundledExerciseSourceRepository
import io.github.frpdoliv3.gymbro.composeApp.data.repository.exercise.ExerciseSourceRepository
import io.github.frpdoliv3.gymbro.composeApp.data.repository.plan.PlanRepository
import io.github.frpdoliv3.gymbro.composeApp.data.repository.plan.PlanRepositoryImpl
import org.koin.core.module.dsl.singleOf
import org.koin.dsl.bind
import org.koin.dsl.module

val databaseModule = module {
    singleOf(::BundledExerciseSourceRepository) bind ExerciseSourceRepository::class
    singleOf(::PlanRepositoryImpl) bind PlanRepository::class
}
