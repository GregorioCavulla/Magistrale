#include <stdio.h>
#include <immintrin.h>

#ifdef _WIN32
#include <intrin.h>
#else
#include <x86intrin.h>
#endif

#define VECTOR_LENGTH 32 
#define SSE_DATA_LANE 16
#define DATA_SIZE 1

int main(){


    u_int64_t clock_counter_scalar_start, clock_counter_scalar_end;
    u_int64_t clock_counter_SIMD_start, clock_counter_SIMD_end;

    clock_counter_scalar_start = __rdtsc();

    signed char A[VECTOR_LENGTH] __attribute__((aligned(SSE_DATA_LANE)));  // 16 byte (128 bit) aligned
    signed char min = 127;

    __m128i *p_A = (__m128i*) A;

    __m128i R_A, T_1, T_2, RES;

    for(int i=0; i<VECTOR_LENGTH; i++){
        A[i] = VECTOR_LENGTH-i;
    }

    for(int i=0; i<VECTOR_LENGTH; i++){
        
        if(A[i]<min){
            min=A[i];
        }
    }

    clock_counter_scalar_end = __rdtsc();

    printf("trovato min: %d\n", min);
    printf("CK trascorsi: %d\n", clock_counter_scalar_end-clock_counter_scalar_start);

    R_A = _mm_load_si128 (p_A);
    T_1 = _mm_srli_si128 (R_A, 8);
    RES = _mm_cmplt_epi8 (R_A, T_1);
    T_2 = _mm_and_si128 (R_A, RES);
    T_1 = _mm_andnot_si128 (RES, T_1);
    RES = _mm_or_si128 (T_1,T_2);

    


}