package io.github.frpdoliv3.gymbro.composeApp.di

import io.github.frpdoliv3.gymbro.composeApp.data.local.AppDatabase
import io.github.frpdoliv3.gymbro.composeApp.data.local.AndroidDatabaseFactory
import io.github.frpdoliv3.gymbro.composeApp.data.local.DatabaseFactory
import io.github.frpdoliv3.gymbro.composeApp.data.local.ExerciseDatabase
import org.koin.android.ext.koin.androidContext
import org.koin.dsl.module

val androidDatabaseModule = module {
    single<DatabaseFactory> { AndroidDatabaseFactory(androidContext()) }

    single<ExerciseDatabase> {
        get<DatabaseFactory>().createExerciseDatabase()
    }

    single<AppDatabase> {
        get<DatabaseFactory>().createAppDatabase()
    }

    single { get<ExerciseDatabase>().exerciseDao() }
    single { get<ExerciseDatabase>().stepDao() }
    single { get<ExerciseDatabase>().exerciseImageDao() }
    single { get<AppDatabase>().planDao() }
    single { get<AppDatabase>().planExerciseDao() }
}
