# 如何解决DP

DP本质上是一种优化技术。动态规划本质上是将问题分解为一系列更小的问题的解决办法，解决每个子问题，然后将它们的结果存储起来，等到下一次相同的子问题出现的时候，只需要查找以前存储的这个问题的结果就可以了，不用在重复的计算。

http://blog.refdash.com/dynamic-programming-tutorial-example/

http://20bits.com/article/introduction-to-dynamic-programming

https://www.quora.com/Are-there-any-good-resources-or-tutorials-for-dynamic-programming-DP-besides-the-TopCoder-tutorial


* 如何发现一个问题是一个DP类型的问题

   这个问题的解决方案是否可以分解为一个个类似的更小问题。

* 识别问题的变量

   现在我们确认这个问题的子问题之间存在某种递归的结构，下一步就要看那些函数的参数在发生变化。
   找出变化的参数的方法是列出几个子问题的参数，然后比较他们。

* 清楚的表达递归关系

   假设你已经解决了子问题，那么如何用子问题构造解决这个大问题。

* 识别base case
   
   一个问题的base case就是它不在依赖其他的子问题。

* 递归的实现还是迭代的实现

   递归的实现会有栈溢出的问题

* 添加记忆

   将昂贵的计算结果存储起来，当相同的输入出现时可以直接查找结果。

* 确定时间复杂度