/* Automatically generated by kernel/time/timeconst.bc */
/* Time conversion constants for HZ == 300 */

#ifndef KERNEL_TIMECONST_H
#define KERNEL_TIMECONST_H

#include <linux/param.h>
#include <linux/types.h>

#if HZ != 300
#error "include/generated/timeconst.h has the wrong HZ value!"
#endif

#define HZ_TO_MSEC_MUL32 U64_C(0xD5555556)
#define HZ_TO_MSEC_ADJ32 U64_C(0x2AAAAAAA)
#define HZ_TO_MSEC_SHR32 30
#define MSEC_TO_HZ_MUL32 U64_C(0x9999999A)
#define MSEC_TO_HZ_ADJ32 U64_C(0x1CCCCCCCC)
#define MSEC_TO_HZ_SHR32 33
#define HZ_TO_MSEC_NUM 10
#define HZ_TO_MSEC_DEN 3
#define MSEC_TO_HZ_NUM 3
#define MSEC_TO_HZ_DEN 10

#define HZ_TO_USEC_MUL32 U64_C(0xD0555556)
#define HZ_TO_USEC_ADJ32 U64_C(0xAAAAA)
#define HZ_TO_USEC_SHR32 20
#define USEC_TO_HZ_MUL32 U64_C(0x9D495183)
#define USEC_TO_HZ_ADJ32 U64_C(0x7FFCB923A29)
#define USEC_TO_HZ_SHR32 43
#define HZ_TO_USEC_NUM 10000
#define HZ_TO_USEC_DEN 3
#define USEC_TO_HZ_NUM 3
#define USEC_TO_HZ_DEN 10000
#define HZ_TO_NSEC_NUM 10000000
#define HZ_TO_NSEC_DEN 3
#define NSEC_TO_HZ_NUM 3
#define NSEC_TO_HZ_DEN 10000000

#endif /* KERNEL_TIMECONST_H */
