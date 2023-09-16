#include <stdint.h>


void* multi_scalar_init_wrapper(void* points, uint64_t len);
void multi_scalar_mult_wrapper(void* p, void* ctx, void* scalars, uint64_t len);
