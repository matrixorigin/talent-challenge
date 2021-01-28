/**
 * @file  DocSet.cpp
 * @brief Implementation to the class DocSet.
 */

#include "DocSet.h"
#include <fstream>
#include <sstream>

using namespace std;

bool DocSet::readFromFile(const std::string& filePath, DocSet& docSet)
{
    ifstream inFile;
    inFile.open(filePath.c_str());
    if (!inFile.is_open())
        return false;
    
    string lineBuffer;
    stringstream strStream;
    
    // All class IDs.
    ClassID catID;
    getline(inFile, lineBuffer);
    size_t pos = lineBuffer.find(',');
    while (pos != string::npos)
    {
        strStream.clear();
        strStream.str(lineBuffer.substr(0, pos));
        strStream >> catID;
        docSet.addClass(catID);
        lineBuffer.erase(0, pos + 1);
        pos = lineBuffer.find(',');
    }
    if (!lineBuffer.empty())
    {
        strStream.clear();
        strStream.str(lineBuffer);
        strStream >> catID;
        docSet.addClass(catID);
    }
    
    // Document count and vector length.
    int nDoc = 0, vecLength = 0;
    getline(inFile, lineBuffer);
    pos = lineBuffer.find(',');
    strStream.clear();
    strStream.str(lineBuffer.substr(0, pos));
    strStream >> nDoc;
    strStream.clear();
    strStream.str(lineBuffer.substr(pos + 1));
    strStream >> vecLength;
    docSet.setVectorLength(vecLength);
    
    // Documents.
    DocID docID;
    DocVector docVec;
    string subStr;
    int index, value;
    size_t semiPos;
    while (nDoc > 0 && !inFile.eof())
    {
        docVec.clearDocVector();
        getline(inFile, lineBuffer);
        
        // DocID and ClassID.
        pos = lineBuffer.find(',');
        strStream.clear();
        strStream.str(lineBuffer.substr(0, pos));
        strStream >> docID;
        lineBuffer.erase(0, pos + 1);
        pos = lineBuffer.find(',');
        strStream.clear();
        strStream.str(lineBuffer.substr(0, pos));
        strStream >> catID;
        
        docVec.setDocID(docID);
        docVec.setClassID(catID);
        
        // Document vector.
        lineBuffer.erase(0, pos + 1);
        pos = lineBuffer.find(',');
        while (pos != string::npos)
        {
            subStr = lineBuffer.substr(0, pos);
            semiPos = subStr.find(':');
            strStream.clear();
            strStream.str(subStr.substr(0, semiPos));
            strStream >> index;
            strStream.clear();
            strStream.str(subStr.substr(semiPos + 1));
            strStream >> value;
            docVec.setDocVector(index, value);
            
            lineBuffer.erase(0, pos + 1);
            pos = lineBuffer.find(',');
        }
        if (!lineBuffer.empty())
        {
            subStr = lineBuffer;
            semiPos = subStr.find(':');
            strStream.clear();
            strStream.str(subStr.substr(0, semiPos));
            strStream >> index;
            strStream.clear();
            strStream.str(subStr.substr(semiPos + 1));
            strStream >> value;
            docVec.setDocVector(index, value);
        }
        
        docSet.addDoc(docVec);
        
        --nDoc;
    }
    
    return true;
}
    


bool DocSet::writeToFile(const std::string& filePath, const DocSet& docSet)
{
    ofstream outFile;
    outFile.open(filePath.c_str());
    if (!outFile.is_open())
        return false;
    
    // All class IDs.
    const set<ClassID>& allClasses = docSet.getAllClasses();
    set<ClassID>::const_iterator citor = allClasses.begin();
    if (citor != allClasses.end())
    {
        outFile << *citor;
        ++citor;
    }
    while (citor != allClasses.end())
    {
        outFile << "," << *citor;
        ++citor;
    }
    outFile << endl;
    
    // Count of documents.
    int nDoc = docSet.getDocCount();
    outFile << nDoc << endl;
    
    // Documents and their class IDs.
    for (int i = 0; i < nDoc; ++i)
    {
        const DocVector& doc = docSet.getDoc(i);
        outFile << doc.getDocID() << "," << doc.getClassID() << endl;
    }
    
    return true;
}



DocSet::DocSet() : vectorLength_(0)
{
}



DocSet::~DocSet()
{
}
    


void DocSet::addClass(ClassID catID)
{
    classIDs_.insert(catID);
}



const std::set<ClassID>& DocSet::getAllClasses() const
{
    return classIDs_;
}



void DocSet::setVectorLength(int length)
{
    vectorLength_ = length;
}



int DocSet::getVectorLength() const
{
    return vectorLength_;
}
    


void DocSet::addDoc(const DocVector& doc)
{
    docs_.push_back(doc);
}
    


int DocSet::getDocCount() const
{
    return static_cast<int>(docs_.size());
}



const DocVector& DocSet::getDoc(int docIndex) const
{
    if (docIndex < 0 || docIndex >= getDocCount())
        return nullDoc_;
    else
        return docs_[docIndex];
}



DocVector& DocSet::getDocA(int docIndex)
{
    if (docIndex < 0 || docIndex >= getDocCount())
        return nullDoc_;
    else
        return docs_[docIndex];
}

