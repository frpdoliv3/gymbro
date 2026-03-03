package io.github.frpdoliv3.gymbro.composeApp.data.local.entity

import androidx.room.ColumnInfo
import androidx.room.Entity
import androidx.room.ForeignKey
import androidx.room.Index
import androidx.room.PrimaryKey

@Entity(
    tableName = "exercise_images",
    indices = [
        Index(value = ["exercise_id"])
    ],
    foreignKeys = [
        ForeignKey(
            entity = ExerciseEntity::class,
            parentColumns = arrayOf("id"),
            childColumns = arrayOf("exercise_id")
        )
    ]
)
data class ExerciseImageEntity(
    @PrimaryKey
    val id: Int,
    @ColumnInfo(name = "exercise_id")
    val exerciseId: Int,
    @ColumnInfo(name = "image_order")
    val imageOrder: Int,
    @ColumnInfo(name = "image_blob")
    val imageBlob: ByteArray,
    @ColumnInfo(name = "mime_type")
    val mimeType: String
) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (javaClass != other?.javaClass) return false

        other as ExerciseImageEntity

        if (id != other.id) return false
        if (exerciseId != other.exerciseId) return false
        if (imageOrder != other.imageOrder) return false
        if (!imageBlob.contentEquals(other.imageBlob)) return false
        if (mimeType != other.mimeType) return false

        return true
    }

    override fun hashCode(): Int {
        var result = id
        result = 31 * result + exerciseId
        result = 31 * result + imageOrder
        result = 31 * result + imageBlob.contentHashCode()
        result = 31 * result + mimeType.hashCode()
        return result
    }
}
