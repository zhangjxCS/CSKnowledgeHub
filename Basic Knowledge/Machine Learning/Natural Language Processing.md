# Natural Language Processing

## Text Process

### Preprocess

- Tokenization: Splitting text into individual words, phrases, symbols, or other meaningful elements called tokens.

- Lowercasing: Converting all characters in the text to lowercase.

- Removing Punctuation and Special Characters: Eliminating punctuation marks and non-alphabetic characters.

- Removing Stop Words: Filtering out common words (e.g., "and", "the", "a") that appear frequently but offer little value in understanding the essence of the text.

- Stemming: Trimming words down to their root form, often resulting in a form that is not a valid word.

- Lemmatization: Converting words into their base or dictionary form, ensuring that the reduced form is a valid word.
- Part-of-Speech Tagging: Identifying and tagging each word's part of speech (e.g., noun, verb, adjective) based on its definition and context.

### Feature Engineering

- Bag of Words (BoW): Represents text as fixed-length vectors by counting the frequency of each word appearing in the document.

- Term Frequency-Inverse Document Frequency (TF-IDF): Weighs the frequency of each word in a document against its rarity across all documents, helping to highlight words that are more interesting and distinguishing for a document.

- Word Embeddings: Maps words or phrases to vectors of real numbers in a high-dimensional space, capturing semantic relationships between words.

- N-grams: Combines adjacent words into phrases of n items (e.g., bigrams are 2-word combinations), helping to preserve some order of words.

## Statistical Model

### Hidden Markov Model

- Transition probability: $P(t_i|t_{i-1})=\frac{C(t_{i-1},t_i)+\epsilon}{\sum_{j=1}^N C(t_{i-1},t_j)+N*\epsilon}$
- Emission probability: $P(w_i|t_i)=\frac{C(t_i,w_i)+\epsilon}{\sum_{j=1}^V C(t_i,w_j)+N*\epsilon}$

- Transition matrix: transition probability between each speech tags
- Emission matrix: emission probability between speech tag and word

<img src="../../Images/008i3skNgy1gyu0c5s7gaj31c6092wfc.jpg" alt="image-20220121194512835" style="zoom:50%;" />

Viterbi Algorithm:

<img src="../../Images/008i3skNgy1gyu0c83rl1j31ci0ju76k.jpg" alt="image-20220121195316069" style="zoom:50%;" />

<img src="../../Images/008i3skNgy1gyu0cax566j31ac0is75t.jpg" alt="image-20220121195818619" style="zoom:50%;" />

### N-Gram Language Model

N-Gram probability: $P(w_N|w_1^{N-1})=\frac{w_1^{N-1}w_N}{C(w_1^{N-1})}$

Sequence Probability: $P(A,B,C,D)=P(A)P(B|A)P(C|A,B)P(D|A,B,C)$

N-Gram Language Model:

- Markov assumption: only last N words matter $P(w_N|w_1^{N-1})=P(w_N|w_{N-1})$ (Bigram)
- Start and end of sentence: start token \<s>, end token\</s>, add N-1 start tokens for N-Gram $model$

$$
p(A,B,C,D)=P(s)P(A|s)P(B|A)P(C|B)P(D|C)P(/s|D)
$$

<img src="../../Images/008i3skNgy1gyu0cen7olj319q0a0t9w.jpg" alt="image-20220121204605339" style="zoom: 33%;" />

Model Evaluation: 

- Perplexity: The smaller the perplexity, the better the model

$$
PP(W)=P(s_1,s_2,...,s_m)^{-1/m}
$$

$$
PP(W)=m\sqrt{\prod_{i=1}^m\prod_{j=1}^{|s_i|}\frac{1}{P(w_j^{(i)}|w_{j-1}^{(i)})}},for\ bigram\ model
$$

$$
PP(W)=m\sqrt{\prod_{i=1}^m\frac{1}{P(w_i|w_{i-1})}},concatenate\ all\ sentences
$$

Out of Vocabulary Words:

- Unknown words can be represented by UNK

Smoothing:

- Laplacian smoothing: $P(w_n|w_{n-1})=\frac{C(w_{n-1},w_n)+1}{C(w_{n-1})+V}$
- Add-k smoothing: $P(w_n|w_{n-1})=\frac{C(w_{n-1},w_n)+k}{C(w_{n-1})+k*V}$

- Backoff: if the higher n-gram probability is missing, then use the lower-order probability

- Interpolation: $\hat P(w_n|w_{n-2},w_{n-1})=\lambda_1P(w_n|w_{n-2},w_{n-1})+\lambda_2P(w_n|w_{n-1})+\lambda_3P(w_n)$

## Embedding

### CBOW

<img src="../../Images/1*cuOmGT7NevP9oJFJfVpRKA.png" alt="NLP 101: Word2Vec — Skip-gram and CBOW | by Ria Kulshrestha | Towards Data  Science" style="zoom:50%;" />

- **Input Layer**: The context words are represented as one-hot vectors and are input into the model. If the context size is *C* (i.e., *C* words before and after the target word), and the vocabulary size is *V*, the input layer will have *C* one-hot encoded vectors of size *V*.
- **Projection Layer**: The one-hot vectors are projected onto a shared, dense embedding layer, which reduces the dimensionality. This layer essentially averages the context word vectors.
- **Output Layer**: A softmax layer then predicts the target word. The model is trained to maximize the probability of the correct target word given the context words.

### Skip-Gram

- **Input Layer**: A single target word is represented as a one-hot vector and input into the model.
- **Projection Layer**: The one-hot vector is projected onto a dense embedding layer, similar to CBOW, but without averaging since there's only one input word.
- **Output Layer**: Multiple softmax layers predict the context words surrounding the target word. The model is trained to maximize the probability of the correct context words given a target word.

### GLOVE

1. **Co-occurrence Matrix Construction**: First, GloVe constructs a large matrix that represents the co-occurrence information of words within a corpus. Each element $X_{ij}$ of this matrix represents how often word *i* occurs in the context of word *j*, capturing the frequency of appearance of word pairs within a defined window size across the whole corpus.
2. **Learning Word Vectors**: The learning model then aims to minimize the difference between the dot product of the vector representations of two words and the logarithm of their co-occurrence probability. The objective function incorporates both these raw co-occurrence counts and the weighted function, which helps to address the disparity in word frequencies. The function is designed to give less weight to very frequent word pairs but still differentiate between rare and common pairs.

## Sequence Model

### RNN

### Siamese Network

## Transformer

### Attention

### Transformer

### BERT

### GPT

## Prompt Engineering

