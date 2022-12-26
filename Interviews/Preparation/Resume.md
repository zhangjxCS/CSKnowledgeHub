## Self Introduction

Start with fact; What I have done; What I'm doing; What I'm going to do

My name is Jingxiang, and I'm currently a master student in Data Science in Columbia University. I did my undergraduate in Tsinghua University in China. During my undergraduate study, I found my passion in data science through some courses projects. So I enrolled in many statistics and computer science courses to strengthen my technical skills. Beyond my major, I also did internships in tech companies like DiDi and Google to learn more about the industry. With solid background in coding and statistics, I am able to build efficient data pipelines and use machine learning techniques to help make business decision. And I also know how to collaborate with cross functional teams. I found I enjoyed analyzing data and make data-driven decisions. Therefore, I applied for data science intern in your company. And after internship, I hope to be a full time employee and continue my data science career.

## Professional Experience

### Google

Overall (challenge is that build an automated pipeline in a short time)

- Situation: I was an intern in Google China. Our team is to help local clothing companies do international business. Our clients put advertisements on Google search, and we helped them to do so. The biggest challenge for our customer was what to sell. In the past, our team used google trends to learn about the markets. We hoped to monitor the fashion trends on social media, too.
- Task: The challenge was that I needed to build an automated pipeline including data ETL, machine learning and visualization and put it into production in a 12-week internship. This is like what a full-stack data scientist should do.
- Action: Firstly, I met with my mentor and cross functional team and our clients to sync on our goal. And my mentor told me that there was a 3rd party web crawling tool named Brandwatch. Then I used Brandwatch's python API to extract the fashion keywords frequency in the past one year. Secondly, I built a machine learning model to predict the future keywords frequency on social media platforms. Thirdly, I established a dashboard with historical and prediction data to visualize the keywords trends and find the rising keywords. Finally, I deployed these three steps on Google Cloud compute engine. My pipeline could update weekly to help our clients find trending items in last week.
- Result: My project is a great tool for our clients to know more about the markets. For example, finding one single trending item like tiktok leggings could help us generate 200k revenue.

Hypothesis Test (use Granger Causality test to validate the hypothesis)

- Situation: I was an intern in Google China. Our team is to help local clothing companies do international business. Our clients put advertisements on Google search, and we helped them to do so. The biggest challenge for our customer is what to sell. In the past, our team used google trends to learn about the markets. However, we hope to be faster to know the market trends.
- Task: There is a hypothesis that fashion trends on social media is ahead of google search. Because a fashion trend usally begins with social media celebrities. My task was to validate this hypothesis. If validated, we could be faster to know the trends with social media data and our following work would be meaningful.
- Action: Firstly, I randomly select 30 fashion rising keywords in 2020 and see the peak in Google trends and social media trends. If social media is ahead of google search, the hypothesis was validated. However, there were two problems of this way, one was that the peak cannot represent the arrival of a fashion trend, the other was that 30 human selected keywords were not representative. After meeting with my mentor, I decided to use more scientific ways to validate hypothesis. I did literature review and found that there was a hypothesis testing method called Granger Causality Test. This method is suitable for determining whether one time series is useful in forecasting another. So it's quite suitable to our problem. After deciding to use this new way, I extracted fashion keywords on google search and social media. And then I performed Granger Causality tests. Finally, the hypothesis test showed that there were 41% fashion keywords trends on social media were ahead of google search trends. 
- Result: The result of hypothesis test meant we could use social media data source to find trending words faster for our clients. Therefore, our following work would be meaningful.

Pipelines (efficiency of the pipelines, use python decorator to solve the API unstable problem)

- Situation: I was an intern in Google China. Our team is to help local clothing companies do international business. Our clients put advertisements on Google search, and we helped them to do so. The biggest challenge for our customer was what to sell. In the past, our team used google trends to learn about the markets. We hoped to monitor the fashion trends on social media, too.
- Task: My task was to build a data pipeline to extract fashion keywords frequency on social media, and use these data to build machine learning model for prediction of keywords future frequency.
- Action: I first met with my mentor and knew that there was a 3rd party web crawling tool named Brandwatch. I could extract keywords frequency on different social media sources such as Twitter, Youtube, Facebook and Instagram. Brandwatch has a python API. And I built a data pipeline with python to extract fashion keywords frequency on Google Cloud compute engine. The pipeline excuted once a week and extracted last week keywords frequency. Then I saved the results to Bigquery database. The most annoying part is that the API is not that stable. Sometimes server errors might occur. I consulted with our team's senior engineer. And followed his suggestion, I used python decorator to build a retry module and solved this problem. And there's another challenging part that the efficiency for data ETL is low because I have to wait until the API call is over. So I tried to used advanced data structure like queue to optimize my pipeline. Finally the pipeline could maintain 10k words and update them weekly.
- Result: Keywords data was saved to Google Cloud Bigquery. And I maintained a 10 thousand keywords database. The data pipelines could update weekly. 

Model (traditional ARIMA model doesn't work well)

- Situation: I was an intern in Google China. Our team is to help local clothing companies do international business. The biggest challenge for our customer is what to sell. In the past, our team used google trends to learn about the markets. We hope to monitor the fashion trends on social media, too.
- Task: After building a data pipeline and saving data into cloud bigquery, my task was to build a machine learning model for prediction of future keywords frequency. This is a time series prediction problem.

- Action: Firstly, I did some exploratory data analysis, especially the autocorrelation and partial autocorrelation coefficient. Then I determined the best p and q for the ARIMA model. And I built a ARIMA model for prediction. However, the MAE for the model is about 0.5, which is really high with a dependent variable range from 0-1. I met with a senior data scientist in our team. And I concluded that the ARIMA model only includes the historical data, which is limited. Therefore, I did feature engineering, included many other features, such as the timestamp, for example, the month, season, week of year, etc. Besides, the keywords could also give me some information. Thus, I added keywords embedding vector trained by twitter words corpus. And I tranfered the keywords into embedding vectors. Finally, I chose XGBoost to train the model, the MAE for the model is 0.33.
- Result: Finally, the model performed well in prediction for future three months. And I also built a machine learning pipeline. Every week, the pipeline will be automatically trained and predicted and saved the prediction data to bigquery.

Dashboard (learn about different people's thoughts and needs)

- Situation: I was an intern in Google China. Our team is to help local clothing companies do international business. The biggest challenge for our customer is what to sell. We hope to build a dashboard to reflect the fashion trends to help our clients
- Task: After building the data pipelines and machine learning pipelines. I had a bigquery database with historical data and prediction. Then I chose to use Google Data Studio to do visualization. 
- Action: The dashboard is a user interface of my work. So it's important to know the clients thoughts. Before doing visualization, I met with our sales team and clients to learn about their needs. And we decieded that the dashboard should have two functions. One is to know what is the fashion trending words, the other is the historical and predictive frequency of the trending words. Thus, I desinged two parts of the dashboard. The first part is a raw table which reflects the keywords frequency and week over week and month over month growth rates. These growth rate could help our clients find rising words. The second part is a line chart for different words. Our clients are able to select words to see the historical trends and prediction of the words. Besides that, I also made some filters, such as keywords categories, social media source, etc.
- Result: The dashboard is a great tool for our clients to know the market trends, and find what to sell for stock preparation. Finding a word like tiktok leggings could generate 200 thousand revenue for the company. 

### DiDi

Overall (Challenge is to improve the efficiency of pipeline with large scale data)

- Situation:  I was a Data Engineer Intern in DiDi. DiDi is Chinese biggest online car hiring company, it's like uber in United States. Our team is in charge of route recommender system. There is one strategy for route recommendation, which is to recommend the most frequent route for a given origin and destination.
- Task: We have a database to store the frequent route for a given origin and destination. My work is to maintain and improve the data pipelines for this database.
- Action: Firstly, I met with team members and learned about the shortcomes of the system. There were two shortcomes. One was that the pipeline was not efficient with a delay of 2 days for data to be stored into database. Besides, there might be some missing data. The other is that the pipeline was only for top 50 cities. After figuring out the problems of system. I started to look out the docs for the data dependency of the system. I figured out there was one useless dataset. Without this dataset, the efficiency could be improved. So I rebuilt a new pipeline with Scala and Spark. The new pipeline was able to extract the frequent route and saved it to database in one day without data missing problem. Another task was to promote the top 50 cities to the whole country with similar data mining efficiency. I built several shell scripts to mine data of different cities in parallel. That significantly increase the efficiency by over 60%.
- Result: Finally, the pipeline was able to extract the whole country's frequent route daily without data missing problem. And it increased the recall ratio of hot route by 13%.

### Tsinghua

Overall (Challenge is to design a new causal model without reference and prior work)

- Situation: I was a research assistant in Tsinghua University. Our team focused on causal inference and industrial data analysis.
- Task: My work is to improve current state-of-the-art causal model.
- Action: Firstly, I did some literature review about state-of-the-art causal models. I found that most of the models represent a linear causal relationship. It is limited for nonlinear causal discovery. After meeting with my supervisor, I designed a structural equation model where the causal relationship described by Gaussian process. In order to model some unobserved variables, I also introduced confounders into the model. My model could find the nonlinear causal relationship between variables and confounders. This step is the most challenging one. Actually, I spent a lot of time proving my model could work and translated my model to an optimization problem. In order to solve the optimization problem, I chose Monte Carlo simulations and Expectation Maximization algorithm due to the unobservable confounders. I also built a python causal inference toolbox of my model. Finally, I evaluated my model in synthetic dataset and real-case dataset. The model showed good performance in both fitting and causal discovery.
- Result:  The toolbox is a great tool to find the causality between different time series and it's quite suitable for spatial-temporal data causal discovery like subway stations passenger flow causal discovery.

## Project Experience

### Wechat Recommender System

Overall (learn about recommender system techniques)

- Situation: I participated in a data contest this spring about short video recommender system. The contest was to predict the occurrence probability of different user actions such as like, comment. The training set was the user actions in the last two weeks and some user information and item information. 
- Task: My task was to build a stable machine learning model for prediction of probability of user actions in the future. And the criteria was weighted ROC-AUC score. 
- Action: Firstly, I did some exploratory data analysis about the frequency of different user actions and decided to do under-sampling because of lack of positive samples. Then I did feature engineering. There are three kinds of features. One is user-action interactive features. For this part, I did target encoding and computed some sliding window statistics of different actions. The second one is user information, I joined this table with previous table. The third one is video information. There is an embedding vector considering video frame, subtitles information together. And I also did label encoding for the video keywords and labels. Finally, I combined user actions, user info, item info together to predict the future user action probabilities. For the models part, I tried different models such as LightGBM, Wide & Deep and DeepFM to fit the data and make predictions. Finally, by combining the results of different models, our team made a good prediction for user action probablities.
- Result: Our model realized an AUC score of 0.67, ranking top 15% among all teams.

### Traffic Congestion Prediction

Overall (time series prediction; first time do a data competition)

- Situation: I participated in a data contest by China Computer Federation. This contest was about traffic prediction.
- Task: My task was to predict the traffic congestion levels by relevant features. The features contained link information including link length width etc. and time series traffic information.
- Action: First, I did some exploratory data analysis and found that historical traffic information had high correlation coefficient with the traffic congestion levels. Secondly, I did feature engineering. I constructed some time series statistical features such as mean, median, minimum and maximum of historical car volume and velocity. And I also considered the topological structures of the link network. I combined the upstream and downstream link information into features. Finally, the feature set were about 60 features. Then I started to build the model. I have tried many models. There were three categories. The first one was historical average model as the baseline. The second one was time series model such as ARIMA and RNN. The third one was traditional tree based model including decision tree and tree based boosting model such as LightGBM. After model building and validation, I found that the LightGBM model had the best performance on the cross validation set. Then I chose LightGBM as my final model. Then I started to tune the hyperparameters for the LightGBM model such as tree max depth, learning rate, regularization coefficient. After finding the best hyperparameters, I built the final model for prediction.  
- Result: Finally, my model result realized a weighted F1 score of 0.47 and ranked 15% among all teams in this contest.

leadership (as a group leader)

- Situation: I participated in a data contest by China Computer Federation. This contest was about traffic prediction. Our team had five people.
- Task: Our task was to predict the traffic congestion levels by relevant features. The features contained link information including link length width etc. and time series traffic information.
- Action: I took initiative to be a group leader first. And then I started to do some exploratory data analysis and find some basic information about this project. Then I set up a meeting with everyone talking about their thoughts about this project. Two of our group members had different thoughts about the project. One was trying to use time series model to solve the problem, the other was trying to use traditional machine learning model. And I listened to their ideas and decided to give both of their ideas a try. So I divided our team into three groups. The first group built history average model as a baseline. The second group tried to use traditional machine learning model like decision tree, SVM to solve the problem. The third group tried to use time series model like ARIMA and RNN to solve the problem. Finally, it showed that tree model had the best performance. So finally we used an ensemble tree model LightGBM to solve the problem and tune the hyper parameters of it.
- Result: Finally, our team performed well in final presentation. Everyone could have a chance to introduce their own model and our problem solving process.

## Leadership Experience

Class Monitor

- Situation: During my sophomore year at Tsinghua University, I was a class monitor.
- Task: My task was to organize class activities such as monthly birthday party, girls and boys festival, etc.
- Action: Girls' festival is one day before international women's day. In order to prepare for a girls' festival. I first organized a boys meeting talking about their ideas. Then I decided to organize a skating activity for all girls. Before that, we posted some advance notice on WeChat and made a video for girls. And I also booked the place at that day. On that day, I led our class to the place and start the activity. During activities, I also helped girls take photos. The activity is a great experience for us to know each other and enhance our friendship.
- Result: Finally, the activity was praised by girls, and we all were happy and immersed in that activity.

Captain of soccer team

- Situation: During my junior year at Tsinghua University, I was captain of soccer team of my school.
- Task: My task was to organize weekly training, make plans for official matches and deal with registration affairs, etc.
- Action:  

