/**
 * @file  DocVector.h
 * @brief Defintion to the class DocVector.
 */

#ifndef DOCVECTOR_H_
#define DOCVECTOR_H_

#include <map>

/** 
 * @brief   Type definition to document ID.
 * @details A valid document ID is an non-negative interger.
 */
typedef int DocID;

/**
 * @brief   Type definition to class ID.
 * @details A valid class ID is an non-negative integer.
 */
typedef int ClassID;

/**
 * class DocVector
 * @brief A data structure representing a document.
 */
class DocVector
{
public:
    DocVector();
    virtual ~DocVector();
    
    /**
     * @brief Set document ID.
     * @param docID DocID to the document.
     */
    void setDocID(DocID docID);
    
    /**
     * @brief  Get document ID.
     * @return DocID to the document.
     */
    DocID getDocID() const;
    
    
    /**
     * @brief Set class to the document.
     * @param catID ClassID to the document.
     */
    void setClassID(ClassID catID);
    
    /**
     * @brief  Get the class to the document.
     * @return ClassID to the document.
     */
    ClassID getClassID() const;
    
    
    /**
     * @brief Set document vector.
     * @param index The zero-based index of a element in the document vector.
     * @param value The value to be set to the element.
     */
    void setDocVector(int index, int value);
    
    /**
     * @brief Get document vector.
     * @param index The zero-based index of a element in the document vector.
     * @return The value of the specific element. If the index is out of range,
     * it returns 0.
     */
    int getDocVector(int index) const;
    
    /**
     * @brief Clear the document vector.
     */
    void clearDocVector();
    
    
    /**
     * @brief   Get length of this document.
     * @details The length of the document is sum of words' occurring times
     * in this document.
     * @return  Length of the document.
     */
    int getDocLength() const;

private:
    /** Document ID. */
    DocID docID_;
    /** Class ID. */
    ClassID classID_;
    /** Document vector. */
    typedef std::map<int, int> DocumentVector;
    DocumentVector docVector_;
};

#endif

