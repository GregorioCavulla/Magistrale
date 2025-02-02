#define DIM 20
#include <stdio.h>
#include <math.h>
#include <sys/time.h>
#include <stdlib.h>
#include <stddef.h>
#include <mpi.h>
#include <time.h>


void print_results(char *prompt, int a[DIM][DIM]);

int main(int argc, char *argv[])
{
    MPI_Init(&argc, &argv);
     
    // Get number of processes and check that 4 processes are used
    int size, rank;
    MPI_Comm_size(MPI_COMM_WORLD, &size);
    MPI_Comm_rank(MPI_COMM_WORLD, &rank);
    if(size > DIM){
        printf("il numero di processi %d è maggiore della dimensione della matrice %d.\n",size, DIM);
        MPI_Abort(MPI_COMM_WORLD, EXIT_FAILURE);
    }
    //printf("[Sono il processo %d di %d]\n",rank,size);
    
    // Define my value
    int A[DIM][DIM],B[DIM][DIM],C[DIM][DIM], i,j;
    int my_A[DIM][DIM], my_C[DIM][DIM];

     
    if (rank == 0){
        srand((unsigned int)time(NULL)); 
        for(i=0;i<DIM;i++){
            for(j=0;j<DIM;j++){
                A[i][j]=rand()%50;
                B[i][j]=rand()%50;
            }
        }
        // verifica
        /*printf("[processo %d] matrice A:\n", rank);
        for(i=0;i<DIM;i++){
            for(j=0;j<DIM;j++)
                printf("%d ",A[i][j]);
            printf("\n");
        }
        printf("[processo %d] matrice B:\n", rank);
        for(j=0;j<DIM;j++){
            for(i=0;i<DIM;i++)
                printf("%d ",B[i][j]);
            printf("\n");
        }*/
    }

    //scatter rows of matrix A to different processes     
    MPI_Scatter(A, DIM*DIM/size, MPI_INT, my_A, DIM*DIM/size, MPI_INT,0,MPI_COMM_WORLD);
    //broadcast matrix B to all processes
    MPI_Bcast(B, DIM*DIM, MPI_INT, 0, MPI_COMM_WORLD);

    //DEBUG
    if (rank == 1){
        printf("[processo %d] matrice my_A:\n", rank);
        for(i=0;i<DIM/size;i++){
            for(j=0;j<DIM;j++)
                printf("%d ",my_A[i][j]);
            printf("\n");
        }
    }
    
    for(i=0;i<DIM/size;i++){
        for(j=0;j<DIM;j++){
            my_C[i][j]=0;
            for(int k=0;k<DIM;k++){
                my_C[i][j]+=my_A[i][k]*B[k][j];
            }
        }
    }

    MPI_Gather(my_C, DIM*DIM/size, MPI_INT, C, DIM*DIM/size, MPI_INT, 0, MPI_COMM_WORLD);

    if (rank == 0){
        printf("[processo %d] matrice C:\n", rank);
        for(i=0;i<DIM;i++){
            for(j=0;j<DIM;j++)
                printf("%d ",C[i][j]);
            printf("\n");
        }
    }
    
    MPI_Finalize();
    
}
