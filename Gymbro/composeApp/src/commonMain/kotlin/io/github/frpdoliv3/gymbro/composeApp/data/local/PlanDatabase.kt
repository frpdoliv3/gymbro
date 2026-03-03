package io.github.frpdoliv3.gymbro.composeApp.data.local

import androidx.room.ConstructedBy
import androidx.room.Database
import androidx.room.RoomDatabase
import androidx.room.RoomDatabaseConstructor
import androidx.sqlite.driver.bundled.BundledSQLiteDriver
import io.github.frpdoliv3.gymbro.composeApp.data.local.plan.dao.PlanDao
import io.github.frpdoliv3.gymbro.composeApp.data.local.plan.dao.PlanExerciseDao
import io.github.frpdoliv3.gymbro.composeApp.data.local.plan.entity.PlanEntity
import io.github.frpdoliv3.gymbro.composeApp.data.local.plan.entity.PlanExerciseEntity
import kotlinx.coroutines.Dispatchers

@Database(
    entities = [
        PlanEntity::class,
        PlanExerciseEntity::class
    ],
    version = 1,
    exportSchema = false
)
@ConstructedBy(PlanDatabaseConstructor::class)
abstract class PlanDatabase : RoomDatabase() {
    abstract fun planDao(): PlanDao
    abstract fun planExerciseDao(): PlanExerciseDao
}

@Suppress("KotlinNoActualForExpect", "EXPECT_ACTUAL_CLASSIFIERS_ARE_IN_BETA_WARNING")
expect object PlanDatabaseConstructor : RoomDatabaseConstructor<PlanDatabase> {
    override fun initialize(): PlanDatabase
}

fun getPlanRoomDatabase(builder: RoomDatabase.Builder<PlanDatabase>): PlanDatabase {
    return builder
        .setDriver(BundledSQLiteDriver())
        .setQueryCoroutineContext(Dispatchers.IO)
        .build()
}
