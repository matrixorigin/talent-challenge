/**
 * @file  DocSet.h
 * @brief Definition of the class DocSet.
 */

#ifndef DOCSET_H_
#define DOCSET_H_

#include <set>
#include <string>
#include <vector>
#include "DocVector.h"

/**
 * class DocSet
 * @brief A data structure represents a collection of documents.
 */
class DocSet
{
public:
    /**
     * @brief Read a document collection from an input file.
     * @param filePath Path to the input file.
     * @param docSet The DocSet object to hold the data stored by the input file.
     * @return Returning true means the data stored in the file has been
     * successfully read out. Otherwise, it returns false.
     */
    static bool readFromFile(const std::string& filePath, DocSet& docSet);
    
    /**
     * @brief Save a DocSet to the output file.
     * @param filePath Path to the output file.
     * @param docSet The DocSet object to be saved.
     * @return Returning true means the data has been successfully written.
     */
    static bool writeToFile(const std::string& filePath, const DocSet& docSet);
    
public:
    DocSet();
    virtual ~DocSet();
    
    /**
     * @brief Add a possible class ID.
     * @param catID Add a new possible class ID.
     */
    void addClass(ClassID catID);
    
    /**
     * @brief Get all possible class IDs.
     */
    const std::set<ClassID>& getAllClasses() const;
    
    /**
     * @brief Set the length of a document vector that represents a document
     * in this document collection.
     * @param length The length of the document vector.
     */
    void setVectorLength(int length);
    
    /**
     * @brief  Get the length of a document vector.
     * @return The length of a document vector.
     */
    int getVectorLength() const;
    
    /**
     * @brief Add a document.
     * @param doc A new document.
     */
    void addDoc(const DocVector& doc);
    
    /**
     * @brief  Get the count of documents in this collection.
     * @return The count of documents in this collection.
     */
    int getDocCount() const;
    
    /**
     * @brief Get a document in this collection.
     * @param docIndex A zero-based index to a document in this collection.
     * @return The const reference to the specified document. If the docIndex
     * is out of range, a reference to an empty document will be returned.
     */
    const DocVector& getDoc(int docIndex) const;
    
    /**
     * @brief Get a document in this collection.
     * @param docIndex A zero-based index to a document in this collection.
     * @return The non-const reference to the specified document. If the docIndex
     * is out of range, a reference to an empty document will be returned.
     */
    DocVector& getDocA(int docIndex);
    
private:
    /** All possible classes. */
    std::set<ClassID> classIDs_;
    /** All documents. */
    std::vector<DocVector> docs_;
    /** An empty document. */
    DocVector nullDoc_;
    /** Length of a document. */
    int vectorLength_;
};

#endif
