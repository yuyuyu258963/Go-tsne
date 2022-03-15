import numpy as np


def cal_pairwise_dist(x):
    sum_x = np.sum(np.square(x), 1)
    dist = np.add(np.add(-2 * np.dot(x, x.T), sum_x).T, sum_x) + 1
    print(1 / dist)
    return dist

if __name__ == '__main__':
  a = np.array([[1,2],[3,4],[3,4]])
  print(cal_pairwise_dist(a))
  pass


