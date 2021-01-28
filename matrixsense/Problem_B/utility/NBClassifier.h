/**
 * @file  NBClassifier.h
 * @brief Definition to the class NBClassifier.
 */

#ifndef NBCLASSIFIER_H_
#define NBCLASSIFIER_H_

#include <vector>
#include "DocSet.h"

/**
 * class NBClassifier
 * @brief A Naive Bayes classifier for a binary-classificaton problem.
 */
class NBClassifier
{
public:
    /**
     * @brief The constructor.
     * @param leftID,rightID  Two IDs to the classes which this classifier are
     * supposed to classify.
     */
    NBClassifier(ClassID leftID, ClassID rightID);
    virtual ~NBClassifier();
    
    /**
     * @brief Train the classifier.
     * @param docSet A set of documents for training the classifier.
     * It is not necessary that docSet only contains the documents belonging
     * to the two classes that this classifier is classifying. This method
     * will select out the related documents for training by itself.
     */
    void train(const DocSet& docSet);
    
    /**
     * @brief Predict a document's class.
     * @param doc The document to be predicted.
     * @return The ID to the class which the document belongs to.
     */
    ClassID predict(const DocVector& doc) const;
    
    /**
     * @brief  Get the left class ID.
     * @return The left class ID.
     */
    ClassID getLeftID() const;
    
    /**
     * @brief  Get the right class ID.
     * @return The right class ID.
     */
    ClassID getRightID() const;
    
private:
    /** Two classes this classifier working on. */
    ClassID leftID_, rightID_;
    /** Classes' probabilities. */
    double IDProbL_, IDProbR_;
    /** Words' conditional probabilities. */
    std::vector<double> termProbL_, termProbR_;
};

#endif

