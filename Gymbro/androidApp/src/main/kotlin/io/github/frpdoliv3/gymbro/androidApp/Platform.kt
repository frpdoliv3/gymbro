package io.github.frpdoliv3.gymbro.androidApp

import android.os.Build

class AndroidPlatform {
    val name: String = "Android ${Build.VERSION.SDK_INT}"
}

fun getPlatform() = AndroidPlatform()
