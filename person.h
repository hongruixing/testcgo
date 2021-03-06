/*
 *  * person.h
 *   * Copyright (C) 2019 Tim Hughes
 *    *
 *     * Distributed under terms of the MIT license.
 *      */

#ifndef PERSON_H
#define PERSON_H

typedef struct APerson  {
	    const char * name;
	        const char * long_name;
} APerson ;

APerson *get_person(const char * name, const char * long_name);
struct Greetee {
    const char *name;
    int year;
};

int greet(struct Greetee *g, char *out);
#endif /* !PERSON_H */
