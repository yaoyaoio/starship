// LINT_C_FILE
//
// Code generated by c-struct-bindgen and ecc - DO NOT EDIT
// See https://github.com/eunomia-bpf/c-struct-bindgen for details.
// struct-bindgen versions: 0.1.0
// source file path: src/agent/wasm/libs/c/event.h
#pragma once

#include <assert.h>
#include <stdint.h>
#include <string.h>

struct event2 {
  float X;
  char B;
  char __pad0[3];
  double Y;
  int Z;
  char __pad1[4];
  long long A;
  short Comm[13];
  char __pad2[6];
  uint64_t Unused_ptr;
} __attribute__((packed));
static_assert(sizeof(struct event2) == 72, "Size of event2 is not 72");
