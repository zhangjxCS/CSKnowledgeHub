## Leadership Principle

- Customer Obsession: Leaders start with the customer and work backwards. They work vigorously to earn and keep customer trust. Although leaders pay attention to competitors, they obsess over customers. (Google machine learning time period)

- Ownership: Leaders are owners. They think long term and don’t sacrifice long-term value for short-term results. They act on behalf of the entire company, beyond just their own team. They never say “that’s not my job." (Google clients visit)

- Invent and Simplify: Leaders expect and require innovation and invention from their teams and always find ways to simplify. They are externally aware, look for new ideas from everywhere, and are not limited by “not invented here." As we do new things, we accept that we may be misunderstood for long periods of time. (Google hypothesis testing; Google data pipeline)

- Are Right, A Lot: Leaders are right a lot. They have strong judgment and good instincts. They seek diverse perspectives and work to disconfirm their beliefs. (Traffic congestion prediction)

- Learn and Be Curious: Leaders are never done learning and always seek to improve themselves. They are curious about new possibilities and act to explore them. (Short video recommender system; DiDi data pipelines)

- Hire and Develop the Best: Leaders raise the performance bar with every hire and promotion. They recognize exceptional talent, and willingly move them throughout the organization. Leaders develop leaders and take seriously their role in coaching others. We work on behalf of our people to invent mechanisms for development like Career Choice. (Traffic congestion prediction)

- Insist on the Highest Standards: Leaders have relentlessly high standards — many people may think these standards are unreasonably high. Leaders are continually raising the bar and drive their teams to deliver high quality products, services, and processes. Leaders ensure that defects do not get sent down the line and that problems are fixed so they stay fixed. (Google hypothesis testing)

- Think Big: Thinking small is a self-fulfilling prophecy. Leaders create and communicate a bold direction that inspires results. They think differently and look around corners for ways to serve customers. (Short video recommender system)

- Bias for Action: Speed matters in business. Many decisions and actions are reversible and do not need extensive study. We value calculated risk taking. (Google clients visit)

- Frugality: Accomplish more with less. Constraints breed resourcefulness, self-sufficiency, and invention. There are no extra points for growing headcount, budget size, or fixed expense. (DiDi data pipeline; Google hypothesis testing)

- Earn Trust: Leaders listen attentively, speak candidly, and treat others respectfully. They are vocally self-critical, even when doing so is awkward or embarrassing. Leaders do not believe their or their team’s body odor smells of perfume. They benchmark themselves and their teams against the best. (Traffic Congestion Prediction; Google machine learning time period)

- Dive Deep: Leaders operate at all levels, stay connected to the details, audit frequently, and are skeptical when metrics and anecdote differ. No task is beneath them. (Machine Learning Research; Short video recommender system)

- Have Backbone; Disagree and Commit: Leaders are obligated to respectfully challenge decisions when they disagree, even when doing so is uncomfortable or exhausting. Leaders have conviction and are tenacious. They do not compromise for the sake of social cohesion. Once a decision is determined, they commit wholly. (Traffic congestion prediction; Google machine learning time period)

- Deliver Results: Leaders focus on the key inputs for their business and deliver them with the right quality and in a timely fashion. Despite setbacks, they rise to the occasion and never settle. (DiDi data pipelines)

- Strive to be Earth's Best Employer: Leaders work every day to create a safer, more productive, higher performing, more diverse, and more just work environment. They lead with empathy, have fun at work, and make it easy for others to have fun. Leaders ask themselves: Are my fellow employees growing? Are they empowered? Are they ready for what's next? Leaders have a vision for and commitment to their employees' personal success, whether that be at Amazon or elsewhere. (Class monitor)

- Success and Scale Bring Broad Responsibility: We started in a garage, but we're not there anymore. We are big, we impact the world, and we are far from perfect. We must be humble and thoughtful about even the secondary effects of our actions. Our local communities, planet, and future generations need us to be better every day. We must begin each day with a determination to make better, do better, and be better for our customers, our employees, our partners, and the world at large. And we must end every day knowing we can do even more tomorrow. Leaders create more than they consume and always leave things better than how they found them. (All the internship and research)

## Behavior

- Why Amazon?

Firstly, Amazon is a top internet company. The mission of Amazon is to continually raise the bar of the customer experience by using the internet and technology. The vision of company inspires me. I'm also an Amazon customer, and I hope join Amazon to help improve the customer experienced and bring broad responsibility to the society. Secondly, the work environment attracts me a lot. Amazon strives to be the best employer and has an open, diverse and productive work environment. I think working in such a company is happy and productive. Finally, there are many talents working for Amazon, and I think working with these people could help me grow up fast.

- Why choose this position?

The job I have applied is the data engineer intern. When I was an undergraduate student in Tsinghua University, I found my passion in Data Science. And I enrolled in many programming and statistics courses to strengthen my technical skills. I also did some internships in Tech companies like DiDi and Google. However, when I were doing internship, I found the most challenging and demanding part of data science is to build data pipelines. Without a stable data pipeline, you cannot build a good machine learning model. Therefore, I hope to be a data engineer in the future. And besides that, I also want to do more engineering work, compared to data analysis, engineering is more visible and predictable. I felt sense of achievement when doing such kind of work. 

- Introduce yourself

My name is Jingxiang, and I'm currently a master student in Data Science in Columbia University. I did my undergraduate in Tsinghua University. During my undergraduate study, I found my passion in data science and data engineering through some courses projects. So I enrolled in many computer science courses to strengthen my technical skills. Beyond my major, I also did internships in tech companies like DiDi and Google to learn more about the industry. With solid background in coding, I am able to build efficient data pipelines and use machine learning techniques to help make data-driven decision. I find it is an interesting and proud thing to successfully build a large scale data pipeline to help make business impacts. Therefore, I applied for data engineer intern in Amazon. And after internship, I hope to convert to a full time employee and continue my data engineering career.

- Introduce a data pipeline you built (Most challenging project)

Invent and simplify; Insist on the highest standards

The situation was that I was an intern in Google China. Our team was to help local clothing companies do international business. We helped our clients published ads on Google search. My task was to build a data pipeline of fashion keywords frequency on social media and helped our clients find the trending items on social media. The challenge was that I need to build the whole pipeline and data warehouse on my own. In order to do this, I did my research and follow the instruction step by step. Firstly, I used a 3rd party web crawling tool named Brandwatch to query keywords frequency on different social media sources such as Twitter, Youtube, and Facebook. Secondly, I wrote a python script with Brandwatch's API to extract fashion keywords frequency. The most annoying part is that the API is not that stable and sometimes server errors might occur. And I tried a new way to build up several retry modules using python decorator to solve this problem. For data warehouse building, although it's simple, I still used star schema to design my data warehouse. There were two tables, one was the fact table recording the keywords frequency from each social media source each day, the other was the dimensional table recording the keywords categories and tags. Finally, I published and automated the entire pipeline on Google cloud compute engine. The results were that I maintained a 10k keywords data warehouse and the data update once a week automatically. That is a good warehouse for machine learning job and visualization.

Learn and be curious; Deliver results

The situation was that I was a Data Engineer Intern in DiDi. DiDi is Chinese biggest online car hiring company, it's like uber in United States. Our team is in charge of route recommender system. There is one strategy for route recommendation, which is to recommend the most frequent route for a given origin and destination. The task for me was to improve the data pipelines for frequent routes mining. And there were two shortcomings for the system, one was the efficiency and missing data, the other was the scale of the pipelines. The pipelines were only for top 50 cities before. The biggest challenge for me was that I was new to Scala and Spark. My action was that firstly I met with my mentor asking for some studying resources to learn about these tech stacks. I even spent some of my personal time for study because I hoped to start my work as soon as possible. Secondly, after knowing some basic knowledge, I started to look out the docs for the data dependency of the system. I figured out there was one useless dataset. Without this dataset, the efficiency could be improved. So I rebuilt a new pipeline with Scala and Spark without this redundant dataset. Another task was to promote the top 50 cities to the whole country. In order to improve the efficiency, I split the cities into 10 groups and built several shell scripts to mine data of different cities  sets in parallel. That significantly increase the efficiency. The result was that I built a data pipeline for frequent route mining for the whole country. The mining time decreased from 2 days to 1 day and the recall ratio of frequent route increased by 13%.

- Failure (Customer obsessed; Earn trust)

Customer obsessed; Earn trust (Negative feedback)

The situation was that I was an intern in Google China this summer. The task was to build up a dashboard reflecting the market trends to help our clients make stock preparation. In my first a few weeks, I spent most of my time in building data pipelines and dashboards. However, I ignore the needs of clients. So in my midterm evaluation, my manager gave me a negative feedback on customer obsession. After that, I felt the most important thing is to know our clients. And I schedule several meetings with our clients and sales people to know about their business. And I knew that the biggest challenge for our clients is what to sell. Therefore, I focused on finding the trending items for our clients in the remaining internship. Finally, my dashboard was a great tool for clients to know the market and decide what to sell in the future. And finding one single item like tiktok leggings could help our clients generate 200k revenue. And I helped our team earn clients' trust. The learning is that I should always keep customers in mind. If there is some negative feedback, ask for the reasons and correct the mistake immediately to minimize the negative influence.

- Most challenging project (Insist on the highest standards; Ownership; Learn and be curious)

Insist on the highest standards; Learn and be curious; Dive deep

The situation was that I was a research assistant in Tsinghua University. Our team focused on causal inference. My task was to improve current state-of-the-art causal model. The challenge was that I never heard about causal inference before. Firstly, I did some literature review about causal inference research. I found that most of the models represent a linear causal relationship. It is limited for nonlinear causal discovery. After meeting with my supervisor, I designed a structural equation model where the causal relationship described by Gaussian process. In order to model some unobserved variables, I also introduced confounders into the model. My model could find the nonlinear causal relationship between variables and confounders. This step is the most challenging one. Actually, I spent a lot of time proving my model could work and translated my model to an optimization problem. In order to solve the optimization problem, I chose Monte Carlo simulations and Expectation Maximization algorithm due to the unobservable confounders. I also built a python causal inference toolbox of my model. Finally, I evaluated my model in synthetic dataset and real-case dataset. The model showed good performance in both fitting and causal discovery. The toolbox was a great tool to find the causality between different time series and it's quite suitable for spatial-temporal data causal discovery like subway stations passenger flow causal discovery.

- Miss deadlines (Customer Obsessed; Deliver Results; Ownership; Dive deep)

Learn and be curious; Deliver results;

The situation was that I was a research assistant in Tsinghua University. And my task was to design a new model for causal inference and submit my work to a top machine learning conference. Before my research, my mentor and I set up a goal to submit my paper on ICML 2021, and the deadline was Feb, 2021. Because I was not familiar with causal inference research. Therefore, firstly I spent a lot of time doing literature review. The causal inference is quite an interesting area in machine learning research. However, I spent too much time learning about it. Then I started to design my own model, but I encountered some problems in parameter estimating method. The time was limited. In order to do it, I often asked my supervisor for advice. And that significantly speed up my research. After solving the problem with my supervisor, I didn't have enough time to finish my paper. And I didn't make it to submit my paper on ICML. Then I asked my supervisor to extend the deadline and resubmit the paper on NIPS 2021. Finally, I finished all the experiments and writing and submitted my well-defined paper to NIPS. The takeaway of this failure was that building up a new thing in research is difficult. So I should communicate more with my supervisor and peers to learn new thing. And I should keep up with the deadline and move fast. If missing a deadline, talk with the stakeholder about it to reduce negative impacts of it and finish it as soon as possible.

- Conflict (Customer Obsessed; Earn trust)

Are right, a lot; Earn trust; Have Backbone, Disagree and Commit

The situation was that I was doing a data contest by China Computer Federation. My work was to predict the traffic congestion level by relevant data. Our team had 5 people. And when we met firstly, I told them that we could use traditional tree based model for prediction. One of my teammates said we should considered time series models instead. I disagreed with him because in this project, the time steps were short, so the time series might not have the best performance. However, he insisted his opinion. So I decided to try both models and build a prototype to compare the results. Finally, it showed that my model performance was better than his. Then I earned his trust and he listened to my suggestion. Our team used tree-based model to make prediction and realize good performance. The best way to persuade others is to show them your result and prove it could work better.

Customer Obsessed; Earn trust;

The situation was that I was a data scientist intern in Google China. Our team was to help local clothing company do international business. My task was to build a machine learning model to predict the future fashion keywords frequency. And I had conflicts with a business analyst. She didn't know too much about machine learning and time series prediction. And she wanted to predict for future half year or even one year. However, there were some limits for time series prediction that the accuracy will decrease with the prediction time steps increase. In order to solve this conflict. I firstly met with sales team to know the customers needs. The stock preparation cycle is about 2 weeks to 1 month. Therefore, if I predicted for future one month trending items, then it would be helpful and enough for our clients to make stock preparation in advance. And I also tested the prediction error for future half year. And I found that the error increased but for future three months, the prediction error are acceptable. After doing these, I met with the business analyst and told her the results. She totally agreed with me and said we only predicted for future three months and showed them on dashboard.

- Take a risk, or do not have much time, to make a decision (Bias for action; Ownership)

Customer Obsessed; Bias for action;

The situation was that I was an intern in Google, and our team is to help local clothing company do international business. I helped our clients build up a dashboard reflecting the market trends. And my colleague and I visited clients to present the result. The task for me was to present the dashboard to clients and collect feedback. Our clients were satisfied with my dashboard. And they had one further more question that if it is possible to use some item pictures to predict the future sales. My colleague didn't know too much about machine learning and deep learning. Although it was not my round to answer the question, I took initiatives to answer that. Firstly, I said that it was possible to extract information from picture using some computer vision techniques like convolutional neural nets to extract features from item pictures. Then we could train machine learning models using these features to predict the future sales. However, there are many features influencing the sales, not only item features, but some features about customers and some confounders like weather, season. So the prediction might not be useful. Our clients were satisfied with my answer. So in this case, I took action and risk to answer the clients' questions and helped our team maintain good relationship with clients.

- Sacrifice short term for long term (Think big)

Think big; Dive deep

The situation was that I participated in a data contest this spring about short video recommender system. The contest was to predict the occurrence probability of different user actions such as like, comment. The training set was the user actions in the last two weeks and some user information and item information. There was a baseline model provided by the contest, and that was the Wide and Deep model. I found that if I fine tune the wide and deep model, it could realize a good performance in prediction. However, I won't learn too much about recommender system. Considering the long term gains, I decided to dive deep to learn more about the recommender system from the traditional collaborate filtering algorithm, matrix decomposition algorithm to the state-of-the-art DeepFM model. And I also summarized my learning and made a GitHub repo to take notes. After that, I started to participate in this contest. Firstly, I did some exploratory data analysis and feature engineering. Then I built three kinds of model, including LighGBM, Wide and Deep and DeepFM to make prediction. Finally, I combined the results of three models. The results showed that my model realized an AUC score of 0.67, ranking top 15% among all teams.

- Tough decision (Bias for action)

Bias for action

The situation was that in my junior year I was about to apply for the master study. And I had two choices, one was to continue studying my undergraduate major, the other was to switch my career to data science. At that time I found I like using data to make impacts. After serious consideration, I decided to follow my heart and switched my career to data science in tech company. After making my decision, my task was to improve my background about data science. And I started to prepare for it immediately. Firstly, I enrolled in many programming courses to strengthen my technical skills. Secondly, I did two internships in tech company to learn more about industry. Then I continued my graduate study in Data Science at Columbia University. The results were that I have learned a lot about data science. Now I'm fully prepared to start my data engineer and data science career. The learning is that when making a tough decision, follow your heart and then start to do it as soon as possible.

- Tell me a time when you came up with a simple solution to solve a problem

Invent and Simplify; Deliver results;

The situation was that I was an intern in Google China. Our team was to help local clothing companies do international business. We helped our clients publish ads on Google search. The biggest challenge for our customer is what to sell. In the past, our team used google trends to learn about the markets. However, we hope to be faster to know the market trends. There is a hypothesis that fashion trends on social media is ahead of google search. Because a fashion trend usally begins with social media celebrities. My task was to validate this hypothesis. If validated, we could be faster to know the trends with social media data and our following work would be meaningful. Firstly, I randomly select 30 fashion rising keywords in 2020 and see the peak in Google trends and social media trends. If social media is ahead of google search, the hypothesis was validated. However, there were two problems of this way, one was that the peak cannot represent the arrival of a fashion trend, the other was that 30 human selected keywords were not representative. I did literature review and found that there was a simple hypothesis testing method called Granger Causality Test. This method is suitable for determining whether one time series is useful in forecasting another. So it's quite suitable to our problem. And then I performed Granger Causality tests for given fashion keywords. Finally, the hypothesis test showed that there were 41% fashion keywords trends on social media were ahead of google search trends. The result of hypothesis test meant we could use social media data source to find trending words faster for our clients. Therefore, our following work would be meaningful.

- Tell me a time when you took on something significant outside your area of responsibility

Ownership; Deliver results; Bias for action

The situation was that I was a data engineer intern at DiDi. Our team was in charge of route recommender system. My task was to improve the efficiency of frequent route mining engine. One day there was an emergent task to calculate the recall ratio of frequent route in all orders. The task was important because we could make future plans based on that. However, all of my colleagues were busy doing their own tasks and did not have spare time to do this. At this point, I took initiative to help with this task. Firstly, I referred to the previous Scala code of the problem, then I started to do this need. Finally I gave the result to product manager on time and continued to do my own project. The results showed that the recall ratio increased with our frequent route mining strategy. So the next step was to improve the process.

- Your strength

I think my strength is that I'm a quick learner. My undergraduate major was Environmental Engineering. It's not that data heavy. But when I found my passion through some course projects. I took initiatives to enroll in many statistics and programming courses and did internships in tech companies to strengthen my data skills. Now I think I have strong data science expertise to work as a data scientist.

- Your weakness

I think my weakness is trying to do everything best but ignore the priority of different tasks. Google (pipeline) ... My mentor told me that sometimes we don't have to do everything best. As an intern, I only need to do one project. But as a full time employee. They usually have a lot work to do. What they do is to meet the standard and expectation but not to do everything best. This way they could do more projects and make more impacts than only do one thing best. I think in the future I should have more experience on deciding priority of different tasks.

- Why Microsoft?

Firstly, Microsoft is a great technology company. I have used many Microsoft fancy products and heard a lot about Bill Gates. Microsoft's mission is to help empower every person and organization to achieve more. And I hope to be one of it. Secondly, Microsoft has the most advanced technologies and growth mindset. I think I will learn many technical skills and grow up fast working here. Finally, Microsoft has a good working environment, not only the benefits and compensation, but also the people. It's diverse and inclusive. Everyone working here could have a chance to contribute their ideas and make a difference.

- What team you want to work with?

I'd like to work in the Office 365 team. Because I have been using office for a long time. And I have some ideas to improve the product. For example, I think OneNote still has something to improve. Firstly, we should be able to set default font for different languages. Secondly, we should be allowed to write through writing pad. I think this is a quite useful function that we may use to take notes in class. In addition, the formula function is not convenient. We shoud be able to write formula through inline latex. If I am a product data scientist, I would learn the customers' need and use A/B testing to help improve the product.

I'd like to work in the Bing Ads team. Because I have some previous experience on Ads and recommendation. In google, I helped our clients put ads on google and generate revenue.

- What's your favorite Microsoft's product/program/team

My favorite Microsoft's product is Office. I have been using them for a long time. I used word for paperwork, excel for managing data, powerpoint for presentation and one note for taking notes. One thing that attracts me most is the diverse built-in function in Office. I could use these function to make fancy presentation or dealing with complex data. The other thing that attracts me is the clarity of Office. The font, document type is enough to make a clear document. Therefore, my favorite product is Office.

- What do you think can be improved for Microsoft products?

I know many Microsoft products. For software, there are Office and Windows. For hardware, there are Surface Laptop and Xbox. For cloud service, there are azure. I have been using office for a long time. And I have some ideas to improve the product. (For example, OneNote)

- Why Data Science? 

My undergraduate specialization was Environmental Data Analysis and Modeling. Through that, I found it was interesting to analyze data to make data-driven decisions. And through my internship in tech companies. I found I could make business impacts with data science techniques. That made me feel sense of achievement. Besides technical skills, data scientist should also have domain knowledge. Compared to software engineer, I could have more opportunities to learn about the business logic and collaborate with cross functional teams. I think this point attracts me a lot.

- What's the most difficult thing in data science?

I think the most difficult and challenging part is to identify the problem and translate business problems into data problems. Not all the business problems could be solved by data science techniques. For example, in Google, our business need is to find trending fashion words. And I translate the problem into finding most frequent fashion keywords on social media platforms. Then I start to collect data and make predictions. Actually, I spent a lot of time meeting with our sales team to learn about their thoughts and need because the most important thing is to make impacts for our business.

- What's the biggest problem you have been facing in Data Science

Identify the problem and find appropriate data for it. Not all the business problems could be solved by data science way. A data scientist should know how to translate business problems into data science problems. For example, I was an intern in Google China this summer...(dashboard building) 

- What's your favorite area in machine learning?

I think my favorite area is time series prediction and recommender system. Time series prediction helps us predict the future. Recommender System helps us personalize user experience. In google, my work is about market trends time series prediction. Through that, we could know the trending items in advance and make stock preparation. In that case, Our clients' products could have more impression and click on internet. And I also did a competition this summer, it's about short video recommender system...(Experience)

- Do you have experiences in image and text data?

Yes, I have taken some deep learning courses in school, and there are some course projects about image data, I did image classification in the project. The project is to classify the hand written digit and I used convolutional neural nets to do so. The accuracy is about 80%. For text data, in one of my summer internship in Google, I analyzed keywords frequency and dealt with text data. I introduced a word embedding vector in my prediction model and it showed good feature importance in the model.

- What neural nets have you trained?

Deep Neural Nets

Convolutional Neural Nets

Recurrent Neural Nets

- What kind of data you have modeled?

Time series data, text data, image data

- Explain deep learning to a 5-year-old child

I would say deep learning is like we human. It could extract information from image like we could watch. And it could extract information from text like we could read and listen.  Then we human may use the information we get to make decisions. Deep learning could do it too. It could make prediction based on the information it get. A machine could behave like human. That is what we called deep learning and artificial intelligence. 