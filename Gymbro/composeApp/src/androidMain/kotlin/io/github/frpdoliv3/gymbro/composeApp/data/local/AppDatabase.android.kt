package io.github.frpdoliv3.gymbro.composeApp.data.local

import android.content.Context
import androidx.room.Room
import androidx.room.RoomDatabase
import kotlinx.coroutines.Dispatchers

private const val EXERCISE_DATABASE_NAME = "exercises.db"
private const val DATABASE_RESOURCE_PATH = "databases/exercises.db"
private const val PLAN_DATABASE_NAME = "gymbro.db"

fun getExerciseDatabaseBuilder(context: Context): RoomDatabase.Builder<AppDatabase> {
    val appContext = context.applicationContext
    val dbFile = appContext.getDatabasePath(EXERCISE_DATABASE_NAME)
    return Room.databaseBuilder<AppDatabase>(
        context = appContext,
        name = dbFile.absolutePath
    )
        .fallbackToDestructiveMigration(dropAllTables = true)
        .createFromInputStream {
            checkNotNull(AppDatabase::class.java.classLoader?.getResourceAsStream(DATABASE_RESOURCE_PATH)) {
                "Missing bundled database resource: $DATABASE_RESOURCE_PATH"
            }
        }
}

fun getPlanDatabaseBuilder(context: Context): RoomDatabase.Builder<PlanDatabase> {
    val appContext = context.applicationContext
    val dbFile = appContext.getDatabasePath(PLAN_DATABASE_NAME)
    return Room.databaseBuilder<PlanDatabase>(
        context = appContext,
        name = dbFile.absolutePath
    )
        .fallbackToDestructiveMigration(dropAllTables = true)
}

fun buildExerciseDatabase(context: Context): AppDatabase {
    return getExerciseDatabaseBuilder(context)
        .setQueryCoroutineContext(Dispatchers.IO)
        .build()
}

fun buildPlanDatabase(context: Context): PlanDatabase {
    return getPlanDatabaseBuilder(context)
        .setQueryCoroutineContext(Dispatchers.IO)
        .build()
}
