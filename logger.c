#include "logger.h"
#include <stdio.h>
#include <stdlib.h>
#include <time.h>

void writeLogger(const char* tag, const char* message) {

    FILE * fp;
    fp = fopen("log.txt", "a+");

    if (!fp) {
            printf("Unable to open file!!\n");
    }

//    time_t now;
//    struct tm * timeinfo;
//    time ( &now );
//    timeinfo=localtime(&now);
    time_t t;
    time(&t);
    fprintf(fp, "%s %s - %s \n", ctime(&t), tag, message);

//    fprintf("%s [%s]: %s\n", ctime(&now), tag, message);
    fclose(fp);
}