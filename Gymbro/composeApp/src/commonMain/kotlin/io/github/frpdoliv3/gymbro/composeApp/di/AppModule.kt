package io.github.frpdoliv3.gymbro.composeApp.di

import org.koin.core.module.dsl.factoryOf
import org.koin.dsl.module

val appModule = module {
    includes(databaseModule)
}
