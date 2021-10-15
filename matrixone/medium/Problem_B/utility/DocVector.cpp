/**
 * @file DocVector.cpp
 * @brief Implementation to the class DocVector.cpp.
 */
 
#include "DocVector.h"

DocVector::DocVector() : docID_(-1), classID_(-1)
{
}



DocVector::~DocVector()
{
}



void DocVector::setDocID(DocID docID)
{
    docID_ = docID;
}



DocID DocVector::getDocID() const
{
    return docID_;
}



void DocVector::setClassID(ClassID catID)
{
    classID_ = catID;
}



ClassID DocVector::getClassID() const
{
    return classID_;
}



void DocVector::setDocVector(int index, int value)
{
    docVector_[index] = value;
}



int DocVector::getDocVector(int index) const
{
    DocumentVector::const_iterator citor = docVector_.find(index);
    if (citor != docVector_.end())
        return citor->second;
    else
        return 0;
}



void DocVector::clearDocVector()
{
    docVector_.clear();
}



int DocVector::getDocLength() const
{
    int length = 0;
    DocumentVector::const_iterator citor = docVector_.begin();
    while (citor != docVector_.end())
    {
        length += citor->second;
        ++citor;
    }
    return length;
}

