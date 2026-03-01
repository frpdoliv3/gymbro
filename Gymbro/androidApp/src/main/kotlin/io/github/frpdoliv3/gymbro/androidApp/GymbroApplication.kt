package io.github.frpdoliv3.gymbro.androidApp

import android.app.Application
import org.koin.android.ext.koin.androidContext
import org.koin.core.context.startKoin
import io.github.frpdoliv3.gymbro.composeApp.di.appModule

class GymbroApplication : Application() {
    override fun onCreate() {
        super.onCreate()

        startKoin {
            androidContext(this@GymbroApplication)
            modules(listOf(appModule))
        }
    }
}
