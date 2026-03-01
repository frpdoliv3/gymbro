package io.github.frpdoliv3.gymbro.composeApp.di

import android.content.Context
import androidx.room.Room
import io.github.frpdoliv3.gymbro.composeApp.data.local.dao.*
import io.github.frpdoliv3.gymbro.composeApp.data.local.database.AppDatabase
import io.github.frpdoliv3.gymbro.composeApp.data.repository.ExerciseRepository
import io.github.frpdoliv3.gymbro.composeApp.data.repository.ExerciseRepositoryImpl
import org.koin.android.ext.koin.androidContext
import org.koin.core.module.dsl.singleOf
import org.koin.dsl.bind
import org.koin.dsl.module

// Android-specific implementation
actual val platformModule = module {
    single { createDatabase(androidContext()) }
    singleOf(::ExerciseRepositoryImpl) bind ExerciseRepository::class
    single { get<AppDatabase>().exerciseDao() }
    single { get<AppDatabase>().categoryDao() }
    single { get<AppDatabase>().muscleDao() }
    single { get<AppDatabase>().stepDao() }
    single { get<AppDatabase>().exerciseImageDao() }
    single { get<AppDatabase>().exerciseMuscleDao() }
    single { get<AppDatabase>().exerciseCategoryDao() }
}

fun createDatabase(context: Context): AppDatabase {
    return Room.databaseBuilder(
        context,
        AppDatabase::class.java,
        "exercises.db"
    ).createFromAsset("databases/exercises.db")  // Load from assets
        .build()
}