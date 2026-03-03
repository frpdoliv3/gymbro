package io.github.frpdoliv3.gymbro.composeApp.presentation.plan

import androidx.lifecycle.ViewModel
import androidx.lifecycle.viewModelScope
import io.github.frpdoliv3.gymbro.composeApp.data.repository.exercise.ExerciseSourceRepository
import io.github.frpdoliv3.gymbro.composeApp.data.repository.plan.PlanRepository
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseRef
import io.github.frpdoliv3.gymbro.composeApp.domain.model.PlannedExercise
import kotlinx.coroutines.flow.MutableStateFlow
import kotlinx.coroutines.flow.SharingStarted
import kotlinx.coroutines.flow.StateFlow
import kotlinx.coroutines.flow.combine
import kotlinx.coroutines.flow.flatMapLatest
import kotlinx.coroutines.flow.stateIn
import kotlinx.coroutines.flow.update
import kotlinx.coroutines.launch

@OptIn(kotlinx.coroutines.ExperimentalCoroutinesApi::class)
class PlanViewModel(
    private val exerciseSourceRepository: ExerciseSourceRepository,
    private val planRepository: PlanRepository
) : ViewModel() {
    private val query = MutableStateFlow("")
    private val detailState = MutableStateFlow(DetailState())

    val uiState: StateFlow<PlanUiState> = combine(
        query,
        query.flatMapLatest { value ->
            if (value.isBlank()) {
                exerciseSourceRepository.observeAllExercises()
            } else {
                exerciseSourceRepository.searchExercises(value)
            }
        },
        exerciseSourceRepository.observeAllExercises(),
        planRepository.observePlanEntries(),
        detailState
    ) { currentQuery, availableExercises, allExercises, planEntries, detail ->
        val exercisesByRef = allExercises.associateBy { it.ref }
        val selected = planEntries.mapNotNull { entry ->
            exercisesByRef[entry.exerciseRef]?.let { summary ->
                PlannedExercise(
                    id = entry.id,
                    position = entry.position,
                    exercise = summary
                )
            }
        }

        PlanUiState(
            query = currentQuery,
            selectedExercises = selected,
            availableExercises = availableExercises,
            isDetailVisible = detail.isVisible,
            detail = detail.exerciseDetail,
            isLoadingDetail = detail.isLoading
        )
    }.stateIn(
        scope = viewModelScope,
        started = SharingStarted.WhileSubscribed(5_000),
        initialValue = PlanUiState()
    )

    fun onQueryChange(value: String) {
        query.value = value
    }

    fun addExercise(exerciseRef: ExerciseRef) {
        viewModelScope.launch {
            planRepository.addExercise(exerciseRef)
        }
    }

    fun removeExercise(planEntryId: Long) {
        viewModelScope.launch {
            planRepository.removeExercise(planEntryId)
        }
    }

    fun moveExerciseUp(planEntryId: Long) {
        viewModelScope.launch {
            planRepository.moveExerciseUp(planEntryId)
        }
    }

    fun moveExerciseDown(planEntryId: Long) {
        viewModelScope.launch {
            planRepository.moveExerciseDown(planEntryId)
        }
    }

    fun openDetail(exerciseRef: ExerciseRef) {
        viewModelScope.launch {
            detailState.value = DetailState(isVisible = true, isLoading = true)
            val detail = exerciseSourceRepository.getExerciseDetail(exerciseRef)
            detailState.value = DetailState(
                isVisible = true,
                exerciseDetail = detail,
                isLoading = false
            )
        }
    }

    fun closeDetail() {
        detailState.update { DetailState() }
    }
}

private data class DetailState(
    val isVisible: Boolean = false,
    val exerciseDetail: io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseDetail? = null,
    val isLoading: Boolean = false
)
