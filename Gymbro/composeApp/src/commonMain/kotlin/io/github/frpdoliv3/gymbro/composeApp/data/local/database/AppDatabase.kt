package io.github.frpdoliv3.gymbro.composeApp.data.local.database

import androidx.room.ConstructedBy
import androidx.room.Database
import androidx.room.RoomDatabase
import androidx.room.RoomDatabaseConstructor
import io.github.frpdoliv3.gymbro.composeApp.data.local.dao.*
import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.*

@Database(
    entities = [
        ExerciseEntity::class,
        CategoryEntity::class,
        MuscleEntity::class,
        StepEntity::class,
        ExerciseImageEntity::class,
        ExerciseMuscleEntity::class,
        ExerciseCategoryEntity::class
    ],
    version = 1
)
@ConstructedBy(AppDatabaseConstructor::class)
abstract class AppDatabase : RoomDatabase() {
    abstract fun exerciseDao(): ExerciseDao
    abstract fun categoryDao(): CategoryDao
    abstract fun muscleDao(): MuscleDao
    abstract fun stepDao(): StepDao
    abstract fun exerciseImageDao(): ExerciseImageDao
    abstract fun exerciseMuscleDao(): ExerciseMuscleDao
    abstract fun exerciseCategoryDao(): ExerciseCategoryDao
}

@Suppress("KotlinNoActualForExpect")
expect object AppDatabaseConstructor : RoomDatabaseConstructor<AppDatabase> {
    override fun initialize(): AppDatabase
}
