/**
 * @file  NBClassifier.cpp
 * @brief Implementation to the class NBClassifier.
 */

#include <cmath>
#include "NBClassifier.h"

#include <iostream>
using namespace std;

NBClassifier::NBClassifier(ClassID leftID, ClassID rightID) :
    leftID_(leftID), rightID_(rightID)
{
    IDProbL_ = IDProbR_ = 0.0;
}



NBClassifier::~NBClassifier()
{
}



void NBClassifier::train(const DocSet& docSet)
{
    using namespace std;
    
    // Document's statistic information.
    int nDocL = 0, nDocR = 0;
    int docLengthSumL = 0, docLengthSumR = 0;
    int nDoc = docSet.getDocCount();
    if (nDoc == 0)
        return;
    int vectorLength = docSet.getVectorLength();
    vector<int> termFrequencyL(vectorLength, 0);
    vector<int> termFrequencyR(vectorLength, 0);
    int i = 0, j = 0;
    for (i = 0; i < nDoc; ++i)
    {
        const DocVector& doc = docSet.getDoc(i);
        ClassID catID = doc.getClassID();
        if (catID == leftID_)
        {
            ++nDocL;
            docLengthSumL += doc.getDocLength();
            for (j = 0; j < vectorLength; ++j)
                termFrequencyL[j] += doc.getDocVector(j);
        }
        else if (catID == rightID_)
        {
            ++nDocR;
            docLengthSumR += doc.getDocLength();
            for (j = 0; j < vectorLength; ++j)
                termFrequencyR[j] += doc.getDocVector(j);
        }
    }
    
    // Classes' probabilities.
    IDProbL_ = log((double)nDocL / (nDocL + nDocR));
    IDProbR_ = log((double)nDocR / (nDocL + nDocR));
    
    // Terms' conditional probabilities.
    termProbL_.clear();
    termProbR_.clear();
    termProbL_.reserve(vectorLength);
    termProbR_.reserve(vectorLength);
    int denominatorL = docLengthSumL + vectorLength;
    int denominatorR = docLengthSumR + vectorLength;
    for (i = 0; i < vectorLength; ++i)
    {
        termProbL_.push_back(log((double)(termFrequencyL[i] + 1) / denominatorL));
        termProbR_.push_back(log((double)(termFrequencyR[i] + 1) / denominatorR));
    }
    /*
    cout << "L(" << leftID_ << "):" << endl
         << "\t" << IDProbL_;
    for (i = 0; i < vectorLength; ++i)
        cout << ", " << termProbL_[i];
    cout << endl;
    cout << "R(" << rightID_ << "):" << endl
         << "\t" << IDProbR_;
    for (i = 0; i < vectorLength; ++i)
        cout << ", " << termProbR_[i];
    cout << endl;
    */
}



ClassID NBClassifier::predict(const DocVector& doc) const
{
    int termFrequency, vecLength = static_cast<int>(termProbL_.size());
    double gradeL = IDProbL_, gradeR = IDProbR_;
    for (int i = 0; i < vecLength; ++i)
    {
        termFrequency = doc.getDocVector(i);
        gradeL += (termFrequency * termProbL_[i]);
        gradeR += (termFrequency * termProbR_[i]);
    }
    //cout << "L(" << leftID_ << "):" << gradeL << " - " << "R(" << rightID_ << "):" << gradeR << endl;
    if (gradeL > gradeR)
        return leftID_;
    else
        return rightID_;
}



ClassID NBClassifier::getLeftID() const
{
    return leftID_;
}
    


ClassID NBClassifier::getRightID() const
{
    return rightID_;
}

