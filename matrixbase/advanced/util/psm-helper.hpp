#ifndef _PSM_HELPER_HPP
#define _PSM_HELPER_HPP 1

#include <iostream>
#include <string>
#include <stdint.h>
#include <sys/time.h>

#if defined(_MSC_VER)

#include <stdlib.h>

#define BIG_CONSTANT(x) (x)

// Other compilers

#else	// defined(_MSC_VER)

#define BIG_CONSTANT(x) (x##LLU)

#endif // !defined(_MSC_VER)

class PSMHelper
{
public:
    static void print_current_timestamp()
    {
        struct timeval tv;
        gettimeofday(&tv, NULL);
        std::cout << (int64_t) tv.tv_sec * 1000000 + (int64_t) tv.tv_usec << std::endl;
    }

    static uint64_t MurmurHash64A(const std::string& key, uint64_t seed = BIG_CONSTANT(0))
    {
        static const uint64_t m = BIG_CONSTANT(0xc6a4a7935bd1e995);
        static const int r = 47;

        uint64_t h = seed ^ (key.length() * m);

        const uint64_t * data = (const uint64_t *)key.c_str();
        const uint64_t * end = data + (key.length()/8);

        while (data != end)
        {
            uint64_t k = *data++;

            k *= m;
            k ^= k >> r;
            k *= m;

            h ^= k;
            h *= m;
        }

        const unsigned char * data2 = (const unsigned char*)data;

        switch(key.length() & 7)
        {
            case 7: h ^= uint64_t(data2[6]) << 48;
            case 6: h ^= uint64_t(data2[5]) << 40;
            case 5: h ^= uint64_t(data2[4]) << 32;
            case 4: h ^= uint64_t(data2[3]) << 24;
            case 3: h ^= uint64_t(data2[2]) << 16;
            case 2: h ^= uint64_t(data2[1]) << 8;
            case 1: h ^= uint64_t(data2[0]);
                    h *= m;
        };

        h ^= h >> r;
        h *= m;
        h ^= h >> r;

        return h;
    }

};

#endif
