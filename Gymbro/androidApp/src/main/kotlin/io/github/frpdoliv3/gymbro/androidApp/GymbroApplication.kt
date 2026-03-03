package io.github.frpdoliv3.gymbro.androidApp

import android.app.Application
import io.github.frpdoliv3.gymbro.composeApp.di.appModule
import io.github.frpdoliv3.gymbro.composeApp.di.androidDatabaseModule
import org.koin.android.ext.koin.androidContext
import org.koin.core.context.startKoin

class GymbroApplication : Application() {
    override fun onCreate() {
        super.onCreate()

        startKoin {
            androidContext(this@GymbroApplication)
            modules(listOf(appModule, androidDatabaseModule))
        }
    }
}
