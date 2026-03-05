package io.github.frpdoliv3.gymbro.composeApp.data.local

import android.content.Context
import androidx.room.Room
import androidx.room.RoomDatabase
import java.io.File

private const val EXERCISE_DATABASE_NAME = "exercises.db"
private const val DATABASE_RESOURCE_PATH = "databases/exercises.db"
private const val PLAN_DATABASE_NAME = "gymbro.db"

private fun getExerciseDatabaseBuilder(context: Context): RoomDatabase.Builder<ExerciseDatabase> {
    val appContext = context.applicationContext
    val dbFile = appContext.getDatabasePath(EXERCISE_DATABASE_NAME)
    copyBundledExerciseDatabaseIfNeeded(dbFile)
    return Room.databaseBuilder<ExerciseDatabase>(
        context = appContext,
        name = dbFile.absolutePath
    )
        .fallbackToDestructiveMigration(dropAllTables = true)
}

private fun getPlanDatabaseBuilder(context: Context): RoomDatabase.Builder<AppDatabase> {
    val appContext = context.applicationContext
    val dbFile = appContext.getDatabasePath(PLAN_DATABASE_NAME)
    return Room.databaseBuilder<AppDatabase>(
        context = appContext,
        name = dbFile.absolutePath
    )
        .fallbackToDestructiveMigration(dropAllTables = true)
}

private fun copyBundledExerciseDatabaseIfNeeded(dbFile: File) {
    if (dbFile.exists()) return

    dbFile.parentFile?.mkdirs()
    checkNotNull(ExerciseDatabase::class.java.classLoader?.getResourceAsStream(DATABASE_RESOURCE_PATH)) {
        "Missing bundled database resource: $DATABASE_RESOURCE_PATH"
    }.use { input ->
        dbFile.outputStream().use { output ->
            input.copyTo(output)
        }
    }
}

class AndroidDatabaseFactory(private val context: Context) : DatabaseFactory {
    override fun createExerciseDatabase(): ExerciseDatabase {
        return getExerciseRoomDatabase(getExerciseDatabaseBuilder(context))
    }

    override fun createAppDatabase(): AppDatabase {
        return getAppRoomDatabase(getPlanDatabaseBuilder(context))
    }
}
