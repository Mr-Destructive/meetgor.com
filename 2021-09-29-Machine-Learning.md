---
article_html: '<h2>Introduction</h2>

  <p>We all hear about terms like <code>Machine Learning</code>, <code>Artificial
  Intelligence</code>, and others, but what do they actually mean and why do you need
  to care about these as a developer. This won''t be a perfect guide in terms of experience
  but surely enough to get anyone through the basics of Machine Learning.</p>

  <p>This is not the kind of article I write but, having such challenges can help
  me become a better technical writer, this is the challenge put forward in the Hashnode
  Bootcamp 4 to get out of my comfort zone. Here''s my take on what I know about Machine
  Learning till now (P.S. Half of the stuff I discovered and re-learned during writing).</p>

  <h2>What is Machine Learning?</h2>

  <p>Machine Learning is a technique in software development to predict and react
  to the inputs without being explicitly programmed or written. We can use the if-else
  condition till a point in time, after seeing real-world examples like customer service,
  driving, playing games(chess, checkers, etc), image prediction, and so on. You can''t
  write code for every single case of these applications, So that is where we see
  Artificial Intelligence.</p>

  <blockquote>

  <p>Artificial Intelligence is a process of simulating human-like behavior into computers
  /robots/ electronic systems.</p>

  </blockquote>

  <p>These are two quite similar terms(A.I., M.L.) but they have their own differences.
  Let''s look at those differences:</p>

  <ul>

  <li>

  <p><strong>Artificial Intelligence is a technology that enables computer systems
  to act and decide like humans.</strong></p>

  </li>

  <li>

  <p><strong>Machine Learning is a process of extracting data and learning from the
  past experience or outcomes from that data.</strong></p>

  </li>

  </ul>

  <p>Machine learning is actually a subset of AI. Machine learning actually is about
  training the computer system about an outcome using the data feed into it. We will
  actually look at the detailed process about it in the next few sections.</p>

  <h2>The Process of Machine Learning</h2>

  <p>The first step should be to choose an idea or a goal that you would like to make
  the system predict or output the results.</p>

  <ol>

  <li>Data Gathering</li>

  <li>Filtering Data</li>

  <li>Selecting an Algorithm</li>

  <li>Training the system</li>

  <li>Verifying and Evaluation of Training</li>

  <li>Improving and Deploying the Model</li>

  </ol>

  <p>Let''s take the example of classifying a picture as either dog or cat.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632902918426/AoHlrL13z.gif"
  alt="Images.gif" /></p>

  <h3>1. Data Gathering</h3>

  <p>We can now move on to collecting the medium of data that will be used by our
  system to get the desired outcome, it might be to predict something, classify certain
  things, take decisions, etc.</p>

  <p>In our example, as we want to classify an image into either a dog or a cat is
  classifying things from the given data set. For that, we will require images that
  depict this constraint.  We can use our personal data, public data and from other
  sources, you would like to get your hands on.</p>

  <p>These are some of the popular places to get data publicly:</p>

  <ul>

  <li><a href="https://www.kaggle.com/">Kaggle</a></li>

  <li><a href="https://datasetsearch.research.google.com/">Google Data Search</a></li>

  <li><a href="https://www.reddit.com/r/datasets/top/?sort=top&amp;t=all">Reddit Datasets</a></li>

  <li><a href="https://github.com/awesomedata/awesome-public-datasets#machinelearning">Public
  Datasets on GitHub</a></li>

  <li><a href="https://registry.opendata.aws/">AWS Registry of Open Data</a></li>

  </ul>

  <h3>2. Filtering Data</h3>

  <p>After you have collected the data from some sources, you will notice that it
  is not perfect as per your needs. And to be honest, there is no dataset that is
  perfect for your requirements, because otherwise there will be a ton of data to
  work with, it might be inefficient for humans to create and sort out from that data.
  So we may have to do it manually or take help from a data scientist.</p>

  <p>But if you are just learning, it will be helpful for you to filter and clean
  the data yourself. There will be things in the data sets missing or there will be
  unwanted things in it. This is a critical step that everyone tends to ignore but
  at the end of the day,  spend about 80% of the time unknowingly. This is quite an
  important step as it decides the efficiency of the model you will have made.</p>

  <ul>

  <li>Remove/Fill in the rows which are empty.</li>

  <li>Remove the columns which are not related to your objective.</li>

  <li>Fix certain wrong or inconsistent data.</li>

  </ul>

  <h4>Group data as Training and Testing</h4>

  <p>After the procedure has been applied, you can now separate the data set as Training
  and Testing Data. You have to create two datasets from one, the prior for training
  and the latter for testing the model system after evaluating the tests.</p>

  <p>For our example, we have to separate the images which will be relatively easier
  to distinguish in the training data, and the tough ones for the testing data as
  it will challenge the model appropriately.</p>

  <h3>3. Selecting an Algorithm</h3>

  <p>Now, this is again an important step as it will make your project''s backbone.
  This will be the algorithm that will identify, predict or decide on the outcomes
  from the data given to it.</p>

  <p>We have the following types of algorithms</p>

  <ul>

  <li>Linear Regression</li>

  <li>Logistic Regression</li>

  <li>Decision Tree</li>

  <li>Artificial Neural Network</li>

  <li>k-Nearest Neighbors (KNN)</li>

  <li>k-Means</li>

  </ul>

  <p>You can choose any one of the above or find other types which will be more or
  less based on these algorithms. This algorithm will be decided by the outcomes you
  want, for example, whether you have to predict, classify, recommend, cluster, etc.
  the outcome from the given data. Different algorithms have different complexity
  as they have a completely different approaches.</p>

  <p>You can research this more and find more about which will be suitable for your
  objectives or application.</p>

  <p>Now an important topic that is misleading, Model is the program that will work
  with the data in association with the algorithm and output the actual objective.
  Model is not the algorithm but it works with the chosen algorithm and processes
  the actual learning in machine learning.</p>

  <p>So,</p>

  <blockquote>

  <p>Model = Algorithm + Data</p>

  </blockquote>

  <p>The model will actually process the data according to the algorithm given and
  fill in the objectives may they be classifying or predicting.</p>

  <h3>4. Training the system</h3>

  <p>Training is a step that is very interesting as it involves actually testing the
  model and it''s really fun. We provide the model the <code>training data</code>
  that we segregated while filtering the data. In this process we try to minimize
  the loss by making changes to the algorithm, fix some data set or bring in some
  additional dataset as per needs and again evaluate the results. This is a loop called
  <code>model fitting</code>.</p>

  <p>This step depends on the learning into consideration, whether you want to provide
  any supervision or not.</p>

  <h3>5. Verifying and Evaluation of Training</h3>

  <p>This is a part of <code>model fitting</code> as it is the part of the loop which
  allows us to evaluate and verify the model.  We can evaluate the model based on
  its accuracy, precision, labels, etc. So based on those parameters we should be
  able to decide its complexity and performance.</p>

  <p>These are important aspects to consider in evaluating the model.</p>

  <ul>

  <li>Accuracy</li>

  <li>Precision</li>

  <li>Recall</li>

  </ul>

  <p>You can get the details of the mathematics and logic involved in evaluating the
  model with some references like:</p>

  <ul>

  <li><a href="https://www.jeremyjordan.me/evaluating-a-machine-learning-model/">Evaluating
  ML model - Jeremy Jordan</a></li>

  <li><a href="https://towardsdatascience.com/various-ways-to-evaluate-a-machine-learning-models-performance-230449055f15">Ways
  to Evaluate ML model - Towards DataScience </a></li>

  </ul>

  <h3>6. Improving and Deploying the Model</h3>

  <p>This might be the final step generally but it depends on the project, there are
  certain aspects that need to be taken care of like:</p>

  <ul>

  <li>Creating an API endpoint</li>

  <li>Analysis and Visualization integration with client-side (web/android/ios/desktop
  app)</li>

  <li>Creating a Pipeline for data input and output from the model.</li>

  </ul>

  <p>There might be other options like CI/CD, Testing, feedback, and other production
  level details that need to be taken care of, you can read more about the deployment
  of machine learning models  <a href="https://christophergs.com/machine%20learning/2019/03/17/how-to-deploy-machine-learning-models/">here</a>.</p>

  <p>You can learn about deploying an ML model for your learning and testing for free
  with the recommendations of  <a href="https://www.freecodecamp.org/news/deploy-your-machine-learning-models-for-free/">FCC</a>.</p>

  <h2>Different Types of Machine Learning</h2>

  <p>There are four basic types of Machine Learning :</p>

  <h3>1. Supervised Learning</h3>

  <p>In this type of ML, the model is given the labeled data in the training dataset
  and is evaluated. We provide both input and output to the model and hence it is
  supervised or tracked throughout the process.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632909214720/1uHALiOG-.png"
  alt="3.png" /></p>

  <p>The above image is just for reference and not directly a model, height and weight
  can be a parameter to consider and are not only the thing to be considered her.
  It''s just for making understand the concept of the learning process.</p>

  <h3>2. Unsupervised Learning</h3>

  <p>In this type of machine learning, the model is trained with unlabeled data. It
  is on the algorithm to actually see the pattern or logic in the dataset provided
  and give the output. The output will be known to the user but is not given to the
  model, hence called unsupervised learning.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632909227565/KtU7Oppkc.png"
  alt="4.png" /></p>

  <h3>3. Semi-Supervised Learning</h3>

  <p>As the name suggests, it is a combination of both Supervised and Unsupervised
  learning. The dataset is given with the label but the model is also allowed to process
  its own label(kind of) into the output. Hence having the best of both worlds. There
  might be even some labeled and some unlabeled datasets as per the requirements of
  the application.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632909276475/TPGy91_LQ.png"
  alt="5.png" /></p>

  <h3>4. Reinforcement Learning</h3>

  <p>In reinforcement learning, the model learns from feedback. It might look similar
  to supervised learning but here the feedback might not be instant and hence causing
  delay and improper decision making from the model. Though it is used in many places,
  it is like the realistic learning behavior of humans depicted to computers with
  this learning system.</p>

  <p><strong>There are other types of learning systems as well but these are used
  quite commonly and are quite versatile as well.</strong></p>

  <h2>Applications of Machine Learning</h2>

  <p>The applications of Machines Learning are all around you. Just look carefully,
  you would have even feedback on a model!</p>

  <ul>

  <li>Personal Assistants (Google Assistant/ Siri/ Alexa)</li>

  <li>Gmail Inbox filter.</li>

  <li><a href="https://towardsdatascience.com/how-youtube-recommends-videos-b6e003a5ab2f">Youtube
  Video Recommendation system</a> .</li>

  <li>Face recognition ( <a href="https://en.wikipedia.org/wiki/DeepFace">DeepFace</a>
  )</li>

  <li>Product Recommendations.</li>

  <li>Self-Driving Cars( <a href="https://www.tesla.com/autopilot">Tesla</a> )</li>

  <li>Traffic Alerts (Google Map)</li>

  <li>Text Improvement (<a href="https://www.grammarly.com/blog/how-grammarly-uses-ai/">Grammarly</a>)</li>

  </ul>

  <p>This list is quite huge and is increasing every day with new technologies and
  growing popularity.</p>

  <h2>Can GitHub Copilot take away developers'' jobs?</h2>

  <p>This just doesn''t focus on developers, it''s every human''s job on target this
  day but really? Is it a matter of concern?

  I don''t think so, because,</p>

  <blockquote>

  <p>the number of jobs lost = the number of jobs given.</p>

  </blockquote>

  <p>There will be a need for humans in some or the other way, remember a computer
  cannot is <strong>not smart</strong> like humans, surely it has improved from what
  we thought a couple of years ago. But who discovered this? HUMANS.</p>

  <p>Yes, Machine Learning is quite a powerful technique but humans will remain the
  essence in the world. It will be dependent on humans how we treat the models and
  use them to our and nature''s advantage and not use them against nature to face
  the consequences later.</p>

  <h2>Conclusion</h2>

  <p>Ok, so from this big article, we can summarize the Machine Learning concept.</p>

  <blockquote>

  <p>Machine learning is a program that has a dataset and algorithm along with the
  model for the objective, we train the model as per the requirements and objectives
  with our dataset.</p>

  </blockquote>

  <p>We were able to understand some common processes involved in Machine Learning.
  We even discussed the applications and the state of Machine Learning in today''s
  world. I hope you found this article helpful. Thank you for reading. Happy Coding
  :)</p>

  '
cover: ''
date: 2021-09-29
datetime: 2021-09-29 00:00:00+00:00
description: We all hear about terms like  This is not the kind of article I write
  but, having such challenges can help me become a better technical writer, this is
  the chal
edit_link: https://github.com/mr-destructive/meetgor.com/edit/gh-pages/blog/posts/2021-09-29-Machine-Learning.md
html: '<h2>Introduction</h2>

  <p>We all hear about terms like <code>Machine Learning</code>, <code>Artificial
  Intelligence</code>, and others, but what do they actually mean and why do you need
  to care about these as a developer. This won''t be a perfect guide in terms of experience
  but surely enough to get anyone through the basics of Machine Learning.</p>

  <p>This is not the kind of article I write but, having such challenges can help
  me become a better technical writer, this is the challenge put forward in the Hashnode
  Bootcamp 4 to get out of my comfort zone. Here''s my take on what I know about Machine
  Learning till now (P.S. Half of the stuff I discovered and re-learned during writing).</p>

  <h2>What is Machine Learning?</h2>

  <p>Machine Learning is a technique in software development to predict and react
  to the inputs without being explicitly programmed or written. We can use the if-else
  condition till a point in time, after seeing real-world examples like customer service,
  driving, playing games(chess, checkers, etc), image prediction, and so on. You can''t
  write code for every single case of these applications, So that is where we see
  Artificial Intelligence.</p>

  <blockquote>

  <p>Artificial Intelligence is a process of simulating human-like behavior into computers
  /robots/ electronic systems.</p>

  </blockquote>

  <p>These are two quite similar terms(A.I., M.L.) but they have their own differences.
  Let''s look at those differences:</p>

  <ul>

  <li>

  <p><strong>Artificial Intelligence is a technology that enables computer systems
  to act and decide like humans.</strong></p>

  </li>

  <li>

  <p><strong>Machine Learning is a process of extracting data and learning from the
  past experience or outcomes from that data.</strong></p>

  </li>

  </ul>

  <p>Machine learning is actually a subset of AI. Machine learning actually is about
  training the computer system about an outcome using the data feed into it. We will
  actually look at the detailed process about it in the next few sections.</p>

  <h2>The Process of Machine Learning</h2>

  <p>The first step should be to choose an idea or a goal that you would like to make
  the system predict or output the results.</p>

  <ol>

  <li>Data Gathering</li>

  <li>Filtering Data</li>

  <li>Selecting an Algorithm</li>

  <li>Training the system</li>

  <li>Verifying and Evaluation of Training</li>

  <li>Improving and Deploying the Model</li>

  </ol>

  <p>Let''s take the example of classifying a picture as either dog or cat.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632902918426/AoHlrL13z.gif"
  alt="Images.gif" /></p>

  <h3>1. Data Gathering</h3>

  <p>We can now move on to collecting the medium of data that will be used by our
  system to get the desired outcome, it might be to predict something, classify certain
  things, take decisions, etc.</p>

  <p>In our example, as we want to classify an image into either a dog or a cat is
  classifying things from the given data set. For that, we will require images that
  depict this constraint.  We can use our personal data, public data and from other
  sources, you would like to get your hands on.</p>

  <p>These are some of the popular places to get data publicly:</p>

  <ul>

  <li><a href="https://www.kaggle.com/">Kaggle</a></li>

  <li><a href="https://datasetsearch.research.google.com/">Google Data Search</a></li>

  <li><a href="https://www.reddit.com/r/datasets/top/?sort=top&amp;t=all">Reddit Datasets</a></li>

  <li><a href="https://github.com/awesomedata/awesome-public-datasets#machinelearning">Public
  Datasets on GitHub</a></li>

  <li><a href="https://registry.opendata.aws/">AWS Registry of Open Data</a></li>

  </ul>

  <h3>2. Filtering Data</h3>

  <p>After you have collected the data from some sources, you will notice that it
  is not perfect as per your needs. And to be honest, there is no dataset that is
  perfect for your requirements, because otherwise there will be a ton of data to
  work with, it might be inefficient for humans to create and sort out from that data.
  So we may have to do it manually or take help from a data scientist.</p>

  <p>But if you are just learning, it will be helpful for you to filter and clean
  the data yourself. There will be things in the data sets missing or there will be
  unwanted things in it. This is a critical step that everyone tends to ignore but
  at the end of the day,  spend about 80% of the time unknowingly. This is quite an
  important step as it decides the efficiency of the model you will have made.</p>

  <ul>

  <li>Remove/Fill in the rows which are empty.</li>

  <li>Remove the columns which are not related to your objective.</li>

  <li>Fix certain wrong or inconsistent data.</li>

  </ul>

  <h4>Group data as Training and Testing</h4>

  <p>After the procedure has been applied, you can now separate the data set as Training
  and Testing Data. You have to create two datasets from one, the prior for training
  and the latter for testing the model system after evaluating the tests.</p>

  <p>For our example, we have to separate the images which will be relatively easier
  to distinguish in the training data, and the tough ones for the testing data as
  it will challenge the model appropriately.</p>

  <h3>3. Selecting an Algorithm</h3>

  <p>Now, this is again an important step as it will make your project''s backbone.
  This will be the algorithm that will identify, predict or decide on the outcomes
  from the data given to it.</p>

  <p>We have the following types of algorithms</p>

  <ul>

  <li>Linear Regression</li>

  <li>Logistic Regression</li>

  <li>Decision Tree</li>

  <li>Artificial Neural Network</li>

  <li>k-Nearest Neighbors (KNN)</li>

  <li>k-Means</li>

  </ul>

  <p>You can choose any one of the above or find other types which will be more or
  less based on these algorithms. This algorithm will be decided by the outcomes you
  want, for example, whether you have to predict, classify, recommend, cluster, etc.
  the outcome from the given data. Different algorithms have different complexity
  as they have a completely different approaches.</p>

  <p>You can research this more and find more about which will be suitable for your
  objectives or application.</p>

  <p>Now an important topic that is misleading, Model is the program that will work
  with the data in association with the algorithm and output the actual objective.
  Model is not the algorithm but it works with the chosen algorithm and processes
  the actual learning in machine learning.</p>

  <p>So,</p>

  <blockquote>

  <p>Model = Algorithm + Data</p>

  </blockquote>

  <p>The model will actually process the data according to the algorithm given and
  fill in the objectives may they be classifying or predicting.</p>

  <h3>4. Training the system</h3>

  <p>Training is a step that is very interesting as it involves actually testing the
  model and it''s really fun. We provide the model the <code>training data</code>
  that we segregated while filtering the data. In this process we try to minimize
  the loss by making changes to the algorithm, fix some data set or bring in some
  additional dataset as per needs and again evaluate the results. This is a loop called
  <code>model fitting</code>.</p>

  <p>This step depends on the learning into consideration, whether you want to provide
  any supervision or not.</p>

  <h3>5. Verifying and Evaluation of Training</h3>

  <p>This is a part of <code>model fitting</code> as it is the part of the loop which
  allows us to evaluate and verify the model.  We can evaluate the model based on
  its accuracy, precision, labels, etc. So based on those parameters we should be
  able to decide its complexity and performance.</p>

  <p>These are important aspects to consider in evaluating the model.</p>

  <ul>

  <li>Accuracy</li>

  <li>Precision</li>

  <li>Recall</li>

  </ul>

  <p>You can get the details of the mathematics and logic involved in evaluating the
  model with some references like:</p>

  <ul>

  <li><a href="https://www.jeremyjordan.me/evaluating-a-machine-learning-model/">Evaluating
  ML model - Jeremy Jordan</a></li>

  <li><a href="https://towardsdatascience.com/various-ways-to-evaluate-a-machine-learning-models-performance-230449055f15">Ways
  to Evaluate ML model - Towards DataScience </a></li>

  </ul>

  <h3>6. Improving and Deploying the Model</h3>

  <p>This might be the final step generally but it depends on the project, there are
  certain aspects that need to be taken care of like:</p>

  <ul>

  <li>Creating an API endpoint</li>

  <li>Analysis and Visualization integration with client-side (web/android/ios/desktop
  app)</li>

  <li>Creating a Pipeline for data input and output from the model.</li>

  </ul>

  <p>There might be other options like CI/CD, Testing, feedback, and other production
  level details that need to be taken care of, you can read more about the deployment
  of machine learning models  <a href="https://christophergs.com/machine%20learning/2019/03/17/how-to-deploy-machine-learning-models/">here</a>.</p>

  <p>You can learn about deploying an ML model for your learning and testing for free
  with the recommendations of  <a href="https://www.freecodecamp.org/news/deploy-your-machine-learning-models-for-free/">FCC</a>.</p>

  <h2>Different Types of Machine Learning</h2>

  <p>There are four basic types of Machine Learning :</p>

  <h3>1. Supervised Learning</h3>

  <p>In this type of ML, the model is given the labeled data in the training dataset
  and is evaluated. We provide both input and output to the model and hence it is
  supervised or tracked throughout the process.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632909214720/1uHALiOG-.png"
  alt="3.png" /></p>

  <p>The above image is just for reference and not directly a model, height and weight
  can be a parameter to consider and are not only the thing to be considered her.
  It''s just for making understand the concept of the learning process.</p>

  <h3>2. Unsupervised Learning</h3>

  <p>In this type of machine learning, the model is trained with unlabeled data. It
  is on the algorithm to actually see the pattern or logic in the dataset provided
  and give the output. The output will be known to the user but is not given to the
  model, hence called unsupervised learning.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632909227565/KtU7Oppkc.png"
  alt="4.png" /></p>

  <h3>3. Semi-Supervised Learning</h3>

  <p>As the name suggests, it is a combination of both Supervised and Unsupervised
  learning. The dataset is given with the label but the model is also allowed to process
  its own label(kind of) into the output. Hence having the best of both worlds. There
  might be even some labeled and some unlabeled datasets as per the requirements of
  the application.</p>

  <p><img src="https://cdn.hashnode.com/res/hashnode/image/upload/v1632909276475/TPGy91_LQ.png"
  alt="5.png" /></p>

  <h3>4. Reinforcement Learning</h3>

  <p>In reinforcement learning, the model learns from feedback. It might look similar
  to supervised learning but here the feedback might not be instant and hence causing
  delay and improper decision making from the model. Though it is used in many places,
  it is like the realistic learning behavior of humans depicted to computers with
  this learning system.</p>

  <p><strong>There are other types of learning systems as well but these are used
  quite commonly and are quite versatile as well.</strong></p>

  <h2>Applications of Machine Learning</h2>

  <p>The applications of Machines Learning are all around you. Just look carefully,
  you would have even feedback on a model!</p>

  <ul>

  <li>Personal Assistants (Google Assistant/ Siri/ Alexa)</li>

  <li>Gmail Inbox filter.</li>

  <li><a href="https://towardsdatascience.com/how-youtube-recommends-videos-b6e003a5ab2f">Youtube
  Video Recommendation system</a> .</li>

  <li>Face recognition ( <a href="https://en.wikipedia.org/wiki/DeepFace">DeepFace</a>
  )</li>

  <li>Product Recommendations.</li>

  <li>Self-Driving Cars( <a href="https://www.tesla.com/autopilot">Tesla</a> )</li>

  <li>Traffic Alerts (Google Map)</li>

  <li>Text Improvement (<a href="https://www.grammarly.com/blog/how-grammarly-uses-ai/">Grammarly</a>)</li>

  </ul>

  <p>This list is quite huge and is increasing every day with new technologies and
  growing popularity.</p>

  <h2>Can GitHub Copilot take away developers'' jobs?</h2>

  <p>This just doesn''t focus on developers, it''s every human''s job on target this
  day but really? Is it a matter of concern?

  I don''t think so, because,</p>

  <blockquote>

  <p>the number of jobs lost = the number of jobs given.</p>

  </blockquote>

  <p>There will be a need for humans in some or the other way, remember a computer
  cannot is <strong>not smart</strong> like humans, surely it has improved from what
  we thought a couple of years ago. But who discovered this? HUMANS.</p>

  <p>Yes, Machine Learning is quite a powerful technique but humans will remain the
  essence in the world. It will be dependent on humans how we treat the models and
  use them to our and nature''s advantage and not use them against nature to face
  the consequences later.</p>

  <h2>Conclusion</h2>

  <p>Ok, so from this big article, we can summarize the Machine Learning concept.</p>

  <blockquote>

  <p>Machine learning is a program that has a dataset and algorithm along with the
  model for the objective, we train the model as per the requirements and objectives
  with our dataset.</p>

  </blockquote>

  <p>We were able to understand some common processes involved in Machine Learning.
  We even discussed the applications and the state of Machine Learning in today''s
  world. I hope you found this article helpful. Thank you for reading. Happy Coding
  :)</p>

  '
image_url: https://res.cloudinary.com/dgpxbrwoz/image/upload/v1643288219/blogmedia/uh0xyxjnpp1olfcksbza.png
long_description: We all hear about terms like  This is not the kind of article I
  write but, having such challenges can help me become a better technical writer,
  this is the challenge put forward in the Hashnode Bootcamp 4 to get out of my comfort
  zone. Here Machine L
now: 2025-05-01 18:11:33.313699
path: blog/posts/2021-09-29-Machine-Learning.md
prevnext: null
slug: ml-intro
status: published
subtitle: 'Understanding the concept and process of Machine Learning. '
tags:
- hashnode
- machine-learning
templateKey: blog-post
title: What is Machine Learning?
today: 2025-05-01
---

## Introduction

We all hear about terms like `Machine Learning`, `Artificial Intelligence`, and others, but what do they actually mean and why do you need to care about these as a developer. This won't be a perfect guide in terms of experience but surely enough to get anyone through the basics of Machine Learning.

This is not the kind of article I write but, having such challenges can help me become a better technical writer, this is the challenge put forward in the Hashnode Bootcamp 4 to get out of my comfort zone. Here's my take on what I know about Machine Learning till now (P.S. Half of the stuff I discovered and re-learned during writing).

## What is Machine Learning?

Machine Learning is a technique in software development to predict and react to the inputs without being explicitly programmed or written. We can use the if-else condition till a point in time, after seeing real-world examples like customer service, driving, playing games(chess, checkers, etc), image prediction, and so on. You can't write code for every single case of these applications, So that is where we see Artificial Intelligence.

> Artificial Intelligence is a process of simulating human-like behavior into computers /robots/ electronic systems.

These are two quite similar terms(A.I., M.L.) but they have their own differences. Let's look at those differences:

- **Artificial Intelligence is a technology that enables computer systems to act and decide like humans.**

-  **Machine Learning is a process of extracting data and learning from the past experience or outcomes from that data.**

Machine learning is actually a subset of AI. Machine learning actually is about training the computer system about an outcome using the data feed into it. We will actually look at the detailed process about it in the next few sections. 

## The Process of Machine Learning

The first step should be to choose an idea or a goal that you would like to make the system predict or output the results.

1. Data Gathering
2. Filtering Data
3. Selecting an Algorithm 
4. Training the system
5. Verifying and Evaluation of Training
6. Improving and Deploying the Model

Let's take the example of classifying a picture as either dog or cat.

![Images.gif](https://cdn.hashnode.com/res/hashnode/image/upload/v1632902918426/AoHlrL13z.gif)

### 1. Data Gathering

We can now move on to collecting the medium of data that will be used by our system to get the desired outcome, it might be to predict something, classify certain things, take decisions, etc.

In our example, as we want to classify an image into either a dog or a cat is classifying things from the given data set. For that, we will require images that depict this constraint.  We can use our personal data, public data and from other sources, you would like to get your hands on.

These are some of the popular places to get data publicly:
   - [Kaggle](https://www.kaggle.com/)
   - [Google Data Search](https://datasetsearch.research.google.com/)
   - [Reddit Datasets](https://www.reddit.com/r/datasets/top/?sort=top&t=all)
   - [Public Datasets on GitHub](https://github.com/awesomedata/awesome-public-datasets#machinelearning)
   - [AWS Registry of Open Data](https://registry.opendata.aws/)

### 2. Filtering Data

After you have collected the data from some sources, you will notice that it is not perfect as per your needs. And to be honest, there is no dataset that is perfect for your requirements, because otherwise there will be a ton of data to work with, it might be inefficient for humans to create and sort out from that data. So we may have to do it manually or take help from a data scientist. 

But if you are just learning, it will be helpful for you to filter and clean the data yourself. There will be things in the data sets missing or there will be unwanted things in it. This is a critical step that everyone tends to ignore but at the end of the day,  spend about 80% of the time unknowingly. This is quite an important step as it decides the efficiency of the model you will have made. 

- Remove/Fill in the rows which are empty.
- Remove the columns which are not related to your objective.
- Fix certain wrong or inconsistent data.

#### Group data as Training and Testing 
After the procedure has been applied, you can now separate the data set as Training and Testing Data. You have to create two datasets from one, the prior for training and the latter for testing the model system after evaluating the tests.

For our example, we have to separate the images which will be relatively easier to distinguish in the training data, and the tough ones for the testing data as it will challenge the model appropriately. 

### 3. Selecting an Algorithm 

Now, this is again an important step as it will make your project's backbone. This will be the algorithm that will identify, predict or decide on the outcomes from the data given to it.

We have the following types of algorithms 

- Linear Regression
- Logistic Regression
- Decision Tree
- Artificial Neural Network
- k-Nearest Neighbors (KNN)
- k-Means

You can choose any one of the above or find other types which will be more or less based on these algorithms. This algorithm will be decided by the outcomes you want, for example, whether you have to predict, classify, recommend, cluster, etc. the outcome from the given data. Different algorithms have different complexity as they have a completely different approaches.

You can research this more and find more about which will be suitable for your objectives or application.  

Now an important topic that is misleading, Model is the program that will work with the data in association with the algorithm and output the actual objective. Model is not the algorithm but it works with the chosen algorithm and processes the actual learning in machine learning. 

So,

> Model = Algorithm + Data

 The model will actually process the data according to the algorithm given and fill in the objectives may they be classifying or predicting. 

### 4. Training the system

Training is a step that is very interesting as it involves actually testing the model and it's really fun. We provide the model the `training data` that we segregated while filtering the data. In this process we try to minimize the loss by making changes to the algorithm, fix some data set or bring in some additional dataset as per needs and again evaluate the results. This is a loop called `model fitting`.

This step depends on the learning into consideration, whether you want to provide any supervision or not. 

### 5. Verifying and Evaluation of Training

This is a part of `model fitting` as it is the part of the loop which allows us to evaluate and verify the model.  We can evaluate the model based on its accuracy, precision, labels, etc. So based on those parameters we should be able to decide its complexity and performance.

These are important aspects to consider in evaluating the model.
- Accuracy 
- Precision  
- Recall 

You can get the details of the mathematics and logic involved in evaluating the model with some references like:

- [Evaluating ML model - Jeremy Jordan](https://www.jeremyjordan.me/evaluating-a-machine-learning-model/)
- [Ways to Evaluate ML model - Towards DataScience ](https://towardsdatascience.com/various-ways-to-evaluate-a-machine-learning-models-performance-230449055f15)

### 6. Improving and Deploying the Model

This might be the final step generally but it depends on the project, there are certain aspects that need to be taken care of like:

- Creating an API endpoint
- Analysis and Visualization integration with client-side (web/android/ios/desktop app)
- Creating a Pipeline for data input and output from the model.

There might be other options like CI/CD, Testing, feedback, and other production level details that need to be taken care of, you can read more about the deployment of machine learning models  [here](https://christophergs.com/machine%20learning/2019/03/17/how-to-deploy-machine-learning-models/). 

You can learn about deploying an ML model for your learning and testing for free with the recommendations of  [FCC](https://www.freecodecamp.org/news/deploy-your-machine-learning-models-for-free/).

## Different Types of Machine Learning

There are four basic types of Machine Learning :

### 1. Supervised Learning

In this type of ML, the model is given the labeled data in the training dataset and is evaluated. We provide both input and output to the model and hence it is supervised or tracked throughout the process. 


![3.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1632909214720/1uHALiOG-.png)

The above image is just for reference and not directly a model, height and weight can be a parameter to consider and are not only the thing to be considered her. It's just for making understand the concept of the learning process. 

### 2. Unsupervised Learning

In this type of machine learning, the model is trained with unlabeled data. It is on the algorithm to actually see the pattern or logic in the dataset provided and give the output. The output will be known to the user but is not given to the model, hence called unsupervised learning. 


![4.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1632909227565/KtU7Oppkc.png)

### 3. Semi-Supervised Learning

As the name suggests, it is a combination of both Supervised and Unsupervised learning. The dataset is given with the label but the model is also allowed to process its own label(kind of) into the output. Hence having the best of both worlds. There might be even some labeled and some unlabeled datasets as per the requirements of the application.

![5.png](https://cdn.hashnode.com/res/hashnode/image/upload/v1632909276475/TPGy91_LQ.png)


### 4. Reinforcement Learning

In reinforcement learning, the model learns from feedback. It might look similar to supervised learning but here the feedback might not be instant and hence causing delay and improper decision making from the model. Though it is used in many places, it is like the realistic learning behavior of humans depicted to computers with this learning system.  

**There are other types of learning systems as well but these are used quite commonly and are quite versatile as well.**


## Applications of Machine Learning

The applications of Machines Learning are all around you. Just look carefully, you would have even feedback on a model! 
- Personal Assistants (Google Assistant/ Siri/ Alexa)
- Gmail Inbox filter.
- [Youtube Video Recommendation system](https://towardsdatascience.com/how-youtube-recommends-videos-b6e003a5ab2f) .
- Face recognition ( [DeepFace](https://en.wikipedia.org/wiki/DeepFace) )
- Product Recommendations.
- Self-Driving Cars( [Tesla](https://www.tesla.com/autopilot) )
- Traffic Alerts (Google Map)
- Text Improvement ([Grammarly](https://www.grammarly.com/blog/how-grammarly-uses-ai/))

This list is quite huge and is increasing every day with new technologies and growing popularity. 

## Can GitHub Copilot take away developers' jobs?

This just doesn't focus on developers, it's every human's job on target this day but really? Is it a matter of concern? 
I don't think so, because,

> the number of jobs lost = the number of jobs given.

There will be a need for humans in some or the other way, remember a computer cannot is **not smart** like humans, surely it has improved from what we thought a couple of years ago. But who discovered this? HUMANS. 

Yes, Machine Learning is quite a powerful technique but humans will remain the essence in the world. It will be dependent on humans how we treat the models and use them to our and nature's advantage and not use them against nature to face the consequences later. 

## Conclusion

Ok, so from this big article, we can summarize the Machine Learning concept.

> Machine learning is a program that has a dataset and algorithm along with the model for the objective, we train the model as per the requirements and objectives with our dataset. 

We were able to understand some common processes involved in Machine Learning. We even discussed the applications and the state of Machine Learning in today's world. I hope you found this article helpful. Thank you for reading. Happy Coding :)