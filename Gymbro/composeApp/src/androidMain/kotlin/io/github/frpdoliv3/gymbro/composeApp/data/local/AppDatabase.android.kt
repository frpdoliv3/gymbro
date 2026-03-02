package io.github.frpdoliv3.gymbro.composeApp.data.local

import android.content.Context
import androidx.room.Room
import androidx.room.RoomDatabase

private const val DATABASE_NAME = "exercises.db"
private const val DATABASE_RESOURCE_PATH = "databases/exercises.db"

fun getDatabaseBuilder(context: Context): RoomDatabase.Builder<AppDatabase> {
    val appContext = context.applicationContext
    val dbFile = appContext.getDatabasePath(DATABASE_NAME)
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
