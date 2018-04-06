from mpl_toolkits.mplot3d import Axes3D
import matplotlib.pyplot as plt


nano = "µs"
milli = "ms"
minute = "m"
second = "s"

with open("new.csv", "r") as f:
	entries = list()
	for l in f.readlines():
		e = [i.strip() for i in l.split(",")]
		entries.append((int(e[0]), int(e[1]), float(e[2]), int(e[3])))

# SID = SentimentIntensityAnalyzer()
# for tweet in tweet_list:
# 	print(tweet)
# 	sent = SID.polarity_scores(tweet)
# 	positive.append(sent['pos'])
# 	negative.append(sent['neg'])
# 	neutral.append(sent['neu'])


size = [e[0] for e in entries]
cores = [e[1] for e in entries]
time = [e[2] for e in entries]
mallocs = [e[3] for e in entries]

# Let's make our graph.
fig = plt.figure()
ax = fig.add_subplot(111, projection='3d')
plt.title("Execution Time") # The name of the topic in question.

from matplotlib import cm
#			  X 		Y 		  Z
ax.scatter(size, cores, time, cmap=cm.Spectral, marker='o')

ax.set_xlabel('size')
ax.set_ylabel('cores')
ax.set_zlabel('time (s)')

plt.show()