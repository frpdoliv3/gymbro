package io.github.frpdoliv3.gymbro.composeApp.infrastructure.local

import androidx.room.ConstructedBy
import androidx.room.Database
import androidx.room.RoomDatabase
import androidx.room.RoomDatabaseConstructor
import androidx.sqlite.driver.bundled.BundledSQLiteDriver
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.dao.PlanDao
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.dao.PlanExerciseDao
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.entity.PlanEntity
import io.github.frpdoliv3.gymbro.composeApp.infrastructure.local.plan.entity.PlanExerciseEntity
import kotlinx.coroutines.Dispatchers

@Database(
    entities = [
        PlanEntity::class,
        PlanExerciseEntity::class
    ],
    version = 1,
    exportSchema = false
)
@ConstructedBy(AppDatabaseConstructor::class)
abstract class AppDatabase : RoomDatabase() {
    abstract fun planDao(): PlanDao
    abstract fun planExerciseDao(): PlanExerciseDao
}

@Suppress("KotlinNoActualForExpect", "EXPECT_ACTUAL_CLASSIFIERS_ARE_IN_BETA_WARNING")
expect object AppDatabaseConstructor : RoomDatabaseConstructor<AppDatabase> {
    override fun initialize(): AppDatabase
}

fun getAppRoomDatabase(builder: RoomDatabase.Builder<AppDatabase>): AppDatabase {
    return builder
        .setDriver(BundledSQLiteDriver())
        .setQueryCoroutineContext(Dispatchers.IO)
        .build()
}
