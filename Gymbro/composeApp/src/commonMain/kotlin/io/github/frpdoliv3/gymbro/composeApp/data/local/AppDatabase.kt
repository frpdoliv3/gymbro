package io.github.frpdoliv3.gymbro.composeApp.data.local

import androidx.room.ConstructedBy
import androidx.room.Database
import androidx.room.RoomDatabase
import androidx.room.RoomDatabaseConstructor
import androidx.sqlite.driver.bundled.BundledSQLiteDriver
import io.github.frpdoliv3.gymbro.composeApp.data.local.dao.*
import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.*
import kotlinx.coroutines.Dispatchers

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

@Suppress("KotlinNoActualForExpect", "EXPECT_ACTUAL_CLASSIFIERS_ARE_IN_BETA_WARNING")
expect object AppDatabaseConstructor : RoomDatabaseConstructor<AppDatabase> {
    override fun initialize(): AppDatabase
}

fun getRoomDatabase(
    builder: RoomDatabase.Builder<AppDatabase>
): AppDatabase {
    return builder
        .setDriver(BundledSQLiteDriver())
        .setQueryCoroutineContext(Dispatchers.IO)
        .build()
}
