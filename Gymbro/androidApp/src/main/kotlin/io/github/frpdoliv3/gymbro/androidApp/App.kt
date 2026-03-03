package io.github.frpdoliv3.gymbro.androidApp

import android.graphics.BitmapFactory
import androidx.compose.foundation.Image
import androidx.compose.foundation.background
import androidx.compose.foundation.clickable
import androidx.compose.foundation.layout.Arrangement
import androidx.compose.foundation.layout.Box
import androidx.compose.foundation.layout.Column
import androidx.compose.foundation.layout.IntrinsicSize
import androidx.compose.foundation.layout.Row
import androidx.compose.foundation.layout.Spacer
import androidx.compose.foundation.layout.fillMaxHeight
import androidx.compose.foundation.layout.fillMaxSize
import androidx.compose.foundation.layout.fillMaxWidth
import androidx.compose.foundation.layout.height
import androidx.compose.foundation.layout.heightIn
import androidx.compose.foundation.layout.padding
import androidx.compose.foundation.layout.safeDrawingPadding
import androidx.compose.foundation.layout.size
import androidx.compose.foundation.layout.width
import androidx.compose.foundation.lazy.LazyColumn
import androidx.compose.foundation.lazy.items
import androidx.compose.foundation.rememberScrollState
import androidx.compose.foundation.shape.RoundedCornerShape
import androidx.compose.foundation.verticalScroll
import androidx.compose.material3.AlertDialog
import androidx.compose.material3.Button
import androidx.compose.material3.Card
import androidx.compose.material3.CardDefaults
import androidx.compose.material3.CircularProgressIndicator
import androidx.compose.material3.FilledTonalButton
import androidx.compose.material3.HorizontalDivider
import androidx.compose.material3.Icon
import androidx.compose.material3.IconButton
import androidx.compose.material3.MaterialTheme
import androidx.compose.material3.OutlinedTextField
import androidx.compose.material3.Surface
import androidx.compose.material3.Text
import androidx.compose.material3.lightColorScheme
import androidx.compose.runtime.Composable
import androidx.compose.runtime.LaunchedEffect
import androidx.compose.runtime.getValue
import androidx.compose.runtime.mutableIntStateOf
import androidx.compose.runtime.remember
import androidx.compose.runtime.setValue
import androidx.compose.ui.Alignment
import androidx.compose.ui.Modifier
import androidx.compose.ui.draw.clip
import androidx.compose.ui.graphics.Brush
import androidx.compose.ui.graphics.Color
import androidx.compose.ui.graphics.asImageBitmap
import androidx.compose.ui.layout.ContentScale
import androidx.compose.ui.res.painterResource
import androidx.compose.ui.res.stringResource
import androidx.compose.ui.semantics.contentDescription
import androidx.compose.ui.semantics.semantics
import androidx.compose.ui.text.font.FontWeight
import androidx.compose.ui.text.style.TextOverflow
import androidx.compose.ui.unit.dp
import androidx.lifecycle.compose.collectAsStateWithLifecycle
import androidx.lifecycle.viewmodel.compose.viewModel
import io.github.frpdoliv3.gymbro.composeApp.data.repository.exercise.ExerciseSourceRepository
import io.github.frpdoliv3.gymbro.composeApp.data.repository.plan.PlanRepository
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseDetail
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseImage
import io.github.frpdoliv3.gymbro.composeApp.domain.model.ExerciseSummary
import io.github.frpdoliv3.gymbro.composeApp.domain.model.PlannedExercise
import io.github.frpdoliv3.gymbro.composeApp.presentation.plan.PlanViewModel
import kotlinx.coroutines.delay
import org.koin.compose.koinInject

@Composable
fun App() {
    val exerciseSourceRepository = koinInject<ExerciseSourceRepository>()
    val planRepository = koinInject<PlanRepository>()
    val viewModel = viewModel {
        PlanViewModel(
            exerciseSourceRepository = exerciseSourceRepository,
            planRepository = planRepository
        )
    }
    val state by viewModel.uiState.collectAsStateWithLifecycle()

    MaterialTheme {
        Surface {
            Box(
                modifier = Modifier
                    .fillMaxSize()
                    .background(
                        Brush.verticalGradient(
                            colors = listOf(
                                MaterialTheme.colorScheme.primaryContainer,
                                MaterialTheme.colorScheme.surface
                            )
                        )
                    )
                    .safeDrawingPadding()
            ) {
                Column(
                    modifier = Modifier
                        .fillMaxSize()
                        .padding(16.dp),
                    verticalArrangement = Arrangement.spacedBy(16.dp)
                ) {
                    val duplicateRefs = remember(state.selectedExercises) {
                        state.selectedExercises
                            .groupingBy { it.exercise.ref }
                            .eachCount()
                            .filterValues { it > 1 }
                            .keys
                    }

                    Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
                        Text(
                            text = stringResource(R.string.workout_plan_title),
                            style = MaterialTheme.typography.headlineMedium,
                            fontWeight = FontWeight.Bold
                        )
                        Text(
                            text = stringResource(R.string.workout_plan_subtitle),
                            style = MaterialTheme.typography.bodyMedium,
                            color = MaterialTheme.colorScheme.onSurfaceVariant
                        )
                    }

                    OutlinedTextField(
                        value = state.query,
                        onValueChange = viewModel::onQueryChange,
                        modifier = Modifier.fillMaxWidth(),
                        label = { Text(stringResource(R.string.search_exercises_label)) },
                        placeholder = { Text(stringResource(R.string.search_exercises_placeholder)) },
                        singleLine = true
                    )

                    Text(
                        text = stringResource(R.string.plan_exercises_title, state.selectedExercises.size),
                        style = MaterialTheme.typography.titleMedium,
                        fontWeight = FontWeight.SemiBold
                    )

                    LazyColumn(
                        modifier = Modifier.weight(1f),
                        verticalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        if (state.selectedExercises.isEmpty()) {
                            item {
                                EmptyStateCard(
                                    title = stringResource(R.string.plan_empty_title),
                                    body = stringResource(R.string.plan_empty_body)
                                )
                            }
                        }

                        items(
                            items = state.selectedExercises,
                            key = { it.id }
                        ) { plannedExercise ->
                            PlannedExerciseCard(
                                plannedExercise = plannedExercise,
                                isDuplicate = plannedExercise.exercise.ref in duplicateRefs,
                                onOpenDetails = { viewModel.openDetail(plannedExercise.exercise.ref) },
                                onMoveUp = { viewModel.moveExerciseUp(plannedExercise.id) },
                                onMoveDown = { viewModel.moveExerciseDown(plannedExercise.id) },
                                onRemove = { viewModel.removeExercise(plannedExercise.id) }
                            )
                        }
                    }

                    HorizontalDivider()

                    Text(
                        text = stringResource(R.string.exercise_library_title, state.availableExercises.size),
                        style = MaterialTheme.typography.titleMedium,
                        fontWeight = FontWeight.SemiBold
                    )

                    LazyColumn(
                        modifier = Modifier.weight(1.2f),
                        verticalArrangement = Arrangement.spacedBy(12.dp)
                    ) {
                        if (state.availableExercises.isEmpty()) {
                            item {
                                EmptyStateCard(
                                    title = stringResource(R.string.no_exercises_found_title),
                                    body = stringResource(R.string.no_exercises_found_body)
                                )
                            }
                        }

                        items(
                            items = state.availableExercises,
                            key = { it.ref.providerId + ":" + it.ref.exerciseId }
                        ) { exercise ->
                            ExerciseCatalogCard(
                                exercise = exercise,
                                onOpenDetails = { viewModel.openDetail(exercise.ref) },
                                onAdd = { viewModel.addExercise(exercise.ref) }
                            )
                        }
                    }
                }

                if (state.isDetailVisible) {
                    ExerciseDetailDialog(
                        detail = state.detail,
                        isLoading = state.isLoadingDetail,
                        onDismiss = viewModel::closeDetail
                    )
                }
            }
        }
    }
}

@Composable
private fun EmptyStateCard(title: String, body: String) {
    Card(
        colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.surfaceVariant)
    ) {
        Column(
            modifier = Modifier.padding(16.dp),
            verticalArrangement = Arrangement.spacedBy(6.dp)
        ) {
            Text(title, style = MaterialTheme.typography.titleSmall, fontWeight = FontWeight.SemiBold)
            Text(body, style = MaterialTheme.typography.bodyMedium)
        }
    }
}

@Composable
private fun PlannedExerciseCard(
    plannedExercise: PlannedExercise,
    isDuplicate: Boolean,
    onOpenDetails: () -> Unit,
    onMoveUp: () -> Unit,
    onMoveDown: () -> Unit,
    onRemove: () -> Unit
) {
    ExerciseInfoCard(
        titlePrefix = "#${plannedExercise.position + 1}",
        exercise = plannedExercise.exercise,
        isDuplicate = isDuplicate,
        onOpenDetails = onOpenDetails,
        trailingActions = {
            PlanExerciseActionRail(
                onMoveUp = onMoveUp,
                onMoveDown = onMoveDown,
                onRemove = onRemove
            )
        }
    )
}

@Composable
private fun ExerciseCatalogCard(
    exercise: ExerciseSummary,
    onOpenDetails: () -> Unit,
    onAdd: () -> Unit
) {
    ExerciseInfoCard(
        titlePrefix = stringResource(R.string.library_prefix),
        exercise = exercise,
        isDuplicate = false,
        onOpenDetails = onOpenDetails,
        bottomActions = {
            FilledTonalButton(onClick = onAdd) {
                Text(stringResource(R.string.add_to_plan))
            }
        }
    )
}

@Composable
private fun ExerciseInfoCard(
    titlePrefix: String,
    exercise: ExerciseSummary,
    isDuplicate: Boolean,
    onOpenDetails: () -> Unit,
    trailingActions: (@Composable () -> Unit)? = null,
    bottomActions: (@Composable () -> Unit)? = null
) {
    Card(
        modifier = Modifier
            .fillMaxWidth()
            .clickable(onClick = onOpenDetails),
        colors = CardDefaults.cardColors(containerColor = MaterialTheme.colorScheme.surface),
        elevation = CardDefaults.cardElevation(defaultElevation = 4.dp)
    ) {
        Column(
            modifier = Modifier.padding(16.dp),
            verticalArrangement = Arrangement.spacedBy(12.dp)
        ) {
            Column(verticalArrangement = Arrangement.spacedBy(6.dp)) {
                Row(
                    modifier = Modifier.fillMaxWidth(),
                    horizontalArrangement = Arrangement.spacedBy(8.dp),
                    verticalAlignment = Alignment.CenterVertically
                ) {
                    Text(
                        text = titlePrefix,
                        style = MaterialTheme.typography.labelLarge,
                        color = MaterialTheme.colorScheme.secondary,
                        fontWeight = FontWeight.Bold
                    )
                    if (isDuplicate) {
                        DuplicatePill()
                    }
                    Spacer(modifier = Modifier.weight(1f))
                }
                Text(
                    text = exercise.name,
                    style = MaterialTheme.typography.titleLarge,
                    fontWeight = FontWeight.SemiBold
                )
            }

            Row(
                modifier = Modifier
                    .fillMaxWidth()
                    .height(IntrinsicSize.Min),
                horizontalArrangement = Arrangement.spacedBy(16.dp),
                verticalAlignment = Alignment.Top
            ) {
                MetadataGrid(
                    modifier = Modifier.weight(1f),
                    exercise = exercise
                )

                if (trailingActions != null) {
                    Box(
                        modifier = Modifier
                            .fillMaxHeight()
                            .width(1.dp)
                            .background(MaterialTheme.colorScheme.surfaceVariant)
                    )
                    trailingActions()
                }
            }

            bottomActions?.invoke()
        }
    }
}

@Composable
private fun PlanExerciseActionRail(
    onMoveUp: () -> Unit,
    onMoveDown: () -> Unit,
    onRemove: () -> Unit
) {
    Column(
        modifier = Modifier
            .width(40.dp)
            .padding(top = 2.dp),
        verticalArrangement = Arrangement.spacedBy(4.dp),
        horizontalAlignment = Alignment.CenterHorizontally
    ) {
        CompactActionButton(
            onClick = onMoveUp,
            contentDescription = stringResource(R.string.move_exercise_up)
        ) {
            Icon(
                painter = painterResource(android.R.drawable.arrow_up_float),
                contentDescription = null
            )
        }
        CompactActionButton(
            onClick = onMoveDown,
            contentDescription = stringResource(R.string.move_exercise_down)
        ) {
            Icon(
                painter = painterResource(android.R.drawable.arrow_down_float),
                contentDescription = null
            )
        }
        CompactActionButton(
            onClick = onRemove,
            contentDescription = stringResource(R.string.remove_exercise)
        ) {
            Icon(
                painter = painterResource(android.R.drawable.ic_menu_delete),
                contentDescription = null,
                tint = MaterialTheme.colorScheme.error
            )
        }
    }
}

@Composable
private fun CompactActionButton(
    onClick: () -> Unit,
    contentDescription: String,
    content: @Composable () -> Unit
) {
    IconButton(
        onClick = onClick,
        modifier = Modifier
            .size(32.dp)
            .semantics { this.contentDescription = contentDescription }
    ) {
        Box(modifier = Modifier.size(18.dp), contentAlignment = Alignment.Center) {
            content()
        }
    }
}

@Composable
private fun DuplicatePill() {
    Box(
        modifier = Modifier.clip(RoundedCornerShape(999.dp)),
        contentAlignment = Alignment.Center
    ) {
        Text(
            text = stringResource(R.string.duplicate),
            modifier = Modifier
                .background(
                    color = Color(0xFFFFE082),
                    shape = RoundedCornerShape(999.dp)
                )
                .padding(horizontal = 10.dp, vertical = 4.dp),
            color = Color(0xFF5F4300),
            style = MaterialTheme.typography.labelMedium,
            fontWeight = FontWeight.Bold
        )
    }
}

@Composable
private fun MetadataGrid(
    modifier: Modifier = Modifier,
    exercise: ExerciseSummary
) {
    Column(
        modifier = modifier,
        verticalArrangement = Arrangement.spacedBy(10.dp)
    ) {
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.spacedBy(12.dp)
        ) {
            MetadataCell(
                modifier = Modifier.weight(1f),
                label = stringResource(R.string.force_label),
                value = exercise.force
            )
            MetadataCell(
                modifier = Modifier.weight(1f),
                label = stringResource(R.string.level_label),
                value = exercise.level
            )
        }
        Row(
            modifier = Modifier.fillMaxWidth(),
            horizontalArrangement = Arrangement.spacedBy(12.dp)
        ) {
            MetadataCell(
                modifier = Modifier.weight(1f),
                label = stringResource(R.string.mechanic_label),
                value = exercise.mechanic
            )
            MetadataCell(
                modifier = Modifier.weight(1f),
                label = stringResource(R.string.equipment_label),
                value = exercise.equipment
            )
        }
    }
}

@Composable
private fun MetadataCell(
    modifier: Modifier = Modifier,
    label: String,
    value: String?
) {
    Column(
        modifier = modifier,
        verticalArrangement = Arrangement.spacedBy(2.dp)
    ) {
        Text(
            text = label,
            style = MaterialTheme.typography.labelMedium,
            color = MaterialTheme.colorScheme.secondary,
            fontWeight = FontWeight.SemiBold,
            maxLines = 1,
            overflow = TextOverflow.Ellipsis
        )
        Text(
            text = value?.ifBlank { null } ?: stringResource(R.string.not_specified),
            style = MaterialTheme.typography.bodyMedium,
            fontWeight = FontWeight.Medium,
            maxLines = 1,
            overflow = TextOverflow.Ellipsis
        )
    }
}

@Composable
private fun ExerciseDetailDialog(
    detail: ExerciseDetail?,
    isLoading: Boolean,
    onDismiss: () -> Unit
) {
    AlertDialog(
        onDismissRequest = onDismiss,
        confirmButton = {
            Button(onClick = onDismiss) {
                Text(stringResource(R.string.close))
            }
        },
        title = {
            Text(
                text = detail?.summary?.name ?: stringResource(R.string.exercise_details),
                style = MaterialTheme.typography.titleLarge,
                fontWeight = FontWeight.SemiBold
            )
        },
        text = {
            Column(
                modifier = Modifier
                    .fillMaxWidth()
                    .heightIn(max = 520.dp)
                    .verticalScroll(rememberScrollState()),
                verticalArrangement = Arrangement.spacedBy(16.dp)
            ) {
                when {
                    isLoading -> {
                        Box(
                            modifier = Modifier
                                .fillMaxWidth()
                                .padding(vertical = 24.dp),
                            contentAlignment = Alignment.Center
                        ) {
                            CircularProgressIndicator()
                        }
                    }

                    detail == null -> {
                        Text(
                            text = stringResource(R.string.no_extra_media_or_instructions),
                            style = MaterialTheme.typography.bodyMedium
                        )
                    }

                    else -> {
                        AnimatedExerciseImage(images = detail.images)

                        if (detail.instructions.isEmpty()) {
                            Text(
                                text = stringResource(R.string.no_detailed_instructions),
                                style = MaterialTheme.typography.bodyMedium
                            )
                        } else {
                            Column(verticalArrangement = Arrangement.spacedBy(10.dp)) {
                                detail.instructions.forEachIndexed { index, instruction ->
                                    Text(
                                        text = "${index + 1}. $instruction",
                                        style = MaterialTheme.typography.bodyMedium
                                    )
                                }
                            }
                        }
                    }
                }
            }
        }
    )
}

@Composable
private fun AnimatedExerciseImage(images: List<ExerciseImage>) {
    if (images.isEmpty()) {
        EmptyStateCard(
            title = stringResource(R.string.no_images_available_title),
            body = stringResource(R.string.no_images_available_body)
        )
        return
    }

    var currentIndex by remember(images) { mutableIntStateOf(0) }

    LaunchedEffect(images) {
        currentIndex = 0
        if (images.size <= 1) return@LaunchedEffect

        while (true) {
            delay(700)
            currentIndex = (currentIndex + 1) % images.size
        }
    }

    val bitmap = remember(images, currentIndex) {
        BitmapFactory.decodeByteArray(
            images[currentIndex].bytes,
            0,
            images[currentIndex].bytes.size
        )?.asImageBitmap()
    }

    if (bitmap == null) {
        EmptyStateCard(
            title = stringResource(R.string.unsupported_image_title),
            body = stringResource(R.string.unsupported_image_body)
        )
        return
    }

    Image(
        bitmap = bitmap,
        contentDescription = null,
        modifier = Modifier
            .fillMaxWidth()
            .height(220.dp)
            .clip(RoundedCornerShape(20.dp)),
        contentScale = ContentScale.Crop
    )
}
