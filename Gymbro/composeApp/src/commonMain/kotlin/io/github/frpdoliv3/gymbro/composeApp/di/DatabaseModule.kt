package io.github.frpdoliv3.gymbro.composeApp.di

import io.github.frpdoliv3.gymbro.composeApp.data.repository.ExerciseRepository
import io.github.frpdoliv3.gymbro.composeApp.data.repository.ExerciseRepositoryImpl
import org.koin.core.module.dsl.factoryOf
import org.koin.dsl.bind
import org.koin.dsl.module

actual val platformModule = module {
    factoryOf(::ExerciseRepositoryImpl) bind ExerciseRepository::class
}