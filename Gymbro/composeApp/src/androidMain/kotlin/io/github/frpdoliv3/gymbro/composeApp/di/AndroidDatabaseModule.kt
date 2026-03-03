package io.github.frpdoliv3.gymbro.composeApp.di

import io.github.frpdoliv3.gymbro.composeApp.data.local.AppDatabase
import io.github.frpdoliv3.gymbro.composeApp.data.local.ExerciseDatabase
import io.github.frpdoliv3.gymbro.composeApp.data.local.buildExerciseDatabase
import io.github.frpdoliv3.gymbro.composeApp.data.local.buildPlanDatabase
import org.koin.android.ext.koin.androidContext
import org.koin.dsl.module

val androidDatabaseModule = module {
    single<ExerciseDatabase> {
        buildExerciseDatabase(androidContext())
    }

    single<AppDatabase> {
        buildPlanDatabase(androidContext())
    }

    single { get<ExerciseDatabase>().exerciseDao() }
    single { get<ExerciseDatabase>().stepDao() }
    single { get<ExerciseDatabase>().exerciseImageDao() }
    single { get<AppDatabase>().planDao() }
    single { get<AppDatabase>().planExerciseDao() }
}
