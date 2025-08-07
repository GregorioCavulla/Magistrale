#!/bin/bash

#SBATCH --account=tra22_IngInfBO
#SBATCH --partition=g100_usr_prod
#SBATCH --nodes=5
#SBATCH --ntasks-per-node=48
#SBATCH -o job.out
#SBATCH -e job.err
module load autoload intelmpi
srun ./sommavet
