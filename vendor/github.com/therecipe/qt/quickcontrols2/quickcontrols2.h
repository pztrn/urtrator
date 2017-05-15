// +build !minimal

#pragma once

#ifndef GO_QTQUICKCONTROLS2_H
#define GO_QTQUICKCONTROLS2_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtQuickControls2_PackedString { char* data; long long len; };
struct QtQuickControls2_PackedList { void* data; long long len; };
struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_Name();
struct QtQuickControls2_PackedString QQuickStyle_QQuickStyle_Path();
void QQuickStyle_QQuickStyle_SetFallbackStyle(char* style);
void QQuickStyle_QQuickStyle_SetStyle(char* style);

#ifdef __cplusplus
}
#endif

#endif