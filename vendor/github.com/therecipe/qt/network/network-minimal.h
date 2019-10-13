// +build minimal

#pragma once

#ifndef GO_QTNETWORK_H
#define GO_QTNETWORK_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtNetwork_PackedString { char* data; long long len; };
struct QtNetwork_PackedList { void* data; long long len; };

#ifdef __cplusplus
}
#endif

#endif