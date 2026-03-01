package io.github.frpdoliv3.gymbro.composeApp.presentation

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import io.github.frpdoliv3.gymbro.composeApp.data.repository.ExerciseRepository
import io.github.frpdoliv3.gymbro.composeApp.data.local.entity.ExerciseEntity
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.map
import kotlinx.coroutines.flow.stateIn

class ExerciseViewModel(
    private val exerciseRepository: ExerciseRepository
) : ViewModel() {

    val exercises: StateFlow<List<ExerciseEntity>> = exerciseRepository.getAllExercises()
        .stateIn(
            scope = viewModelScope,
            started = SharingStarted.WhileSubscribed(5000),
            initialValue = emptyList()
        )

    fun searchExercises(query: String): StateFlow<List<ExerciseEntity>> = exerciseRepository.searchExercisesByName(query)
        .stateIn(
            scope = viewModelScope,
            started = SharingStarted.WhileSubscribed(5000),
            initialValue = emptyList()
        )

    fun getExerciseById(id: Int) = exerciseRepository.getExerciseById(id)
}