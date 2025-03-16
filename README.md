# Project: Share Depository with GO

##  ![Project Demo](https://raw.githubusercontent.com/thehealer15/share-depostitory/main/go%20project%20recording.mp4)
[![Watch the video](https://img.youtube.com/vi/YOUR_VIDEO_ID/0.jpg)](https://youtu.be/W2d_fk07iyE)


**Design**
--
Whenever a investor is onboarded ; we are creating schema per investor. 
so investor gets isolation and this does not remain bottleneck when application scales 
To ensure we are consistent, app used a transaction/ rollback mechanism inherited with go lang  
App is enough capable to operate several investor (different) at same time while not much good for one investor operating on same API at same time 

---

**What's done in project ?**

* Started postgres:alpine-17 image locally so local database is taken care of 
* Seeded DB with "platform" schema which is holding information about platform includes,
"companies" information which are primary truth for which companies are onboarded and "investor" table holding record-keeping of onboarded investor 
* In per investor schema, we are having "holding" table which has portfolio of investor 

**Tech used**
--

* Go lang for APIs 
* postgres for DB
--

To run the project : 
- clone repo
- have docker desktop with postgres alpine 17 image
- start go app

Test cases validated listed in file:/testcases.http file
