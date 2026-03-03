package io.github.frpdoliv3.gymbro.composeApp.di

import io.github.frpdoliv3.gymbro.composeApp.data.local.AppDatabase
import io.github.frpdoliv3.gymbro.composeApp.data.local.PlanDatabase
import io.github.frpdoliv3.gymbro.composeApp.data.local.buildExerciseDatabase
import io.github.frpdoliv3.gymbro.composeApp.data.local.buildPlanDatabase
import org.koin.android.ext.koin.androidContext
import org.koin.dsl.module

val androidDatabaseModule = module {
    single<AppDatabase> {
        buildExerciseDatabase(androidContext())
    }

    single<PlanDatabase> {
        buildPlanDatabase(androidContext())
    }

    single { get<AppDatabase>().exerciseDao() }
    single { get<AppDatabase>().stepDao() }
    single { get<AppDatabase>().exerciseImageDao() }
    single { get<PlanDatabase>().planDao() }
    single { get<PlanDatabase>().planExerciseDao() }
}
