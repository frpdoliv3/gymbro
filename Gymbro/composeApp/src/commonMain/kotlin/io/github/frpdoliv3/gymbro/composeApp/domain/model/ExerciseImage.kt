package io.github.frpdoliv3.gymbro.composeApp.domain.model

data class ExerciseImage(
    val mimeType: String,
    val bytes: ByteArray
) {
    override fun equals(other: Any?): Boolean {
        if (this === other) return true
        if (other !is ExerciseImage) return false
        return mimeType == other.mimeType && bytes.contentEquals(other.bytes)
    }

    override fun hashCode(): Int {
        var result = mimeType.hashCode()
        result = 31 * result + bytes.contentHashCode()
        return result
    }
}
