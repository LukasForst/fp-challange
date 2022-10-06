# Solution for Fingerprint Task

See [task.md](task.md) for assignment.

The task is a nice example of [Knapsack Problem](https://en.wikipedia.org/wiki/Knapsack_problem), more specifically 0-1
knapsack problem where we either include transaction in our batch or do not.
In our case, the weights are latencies and values the amount of $$$ that we can verify.

We know that the solution is optimal, because we actually run all combinations of the transactions and their respective
values and weights. However, thanks to the dynamic programming approach, we were abel to do that quickly, as we
remembered all previous computations.

Note that this can not be generalized - we could do that thanks to the relatively small capacity (`totalTime <= 1000`),
because this algorithm as quite high space complexity that depends on the maximal capacity of the knapsack.
In a case where the capacity would be very high and the increments are high as well, we should use different algorithm
to solve this. As this is NP-complete, we would probably use some approximation
algorithm. [Wiki page](https://en.wikipedia.org/wiki/Knapsack_problem) mentions some of them.

My solution is implemented in Go, as I was a bit rusty in that language, I wanted to refresh my knowledge.
It took me about ~2-3 hours to implement it and make it work. The assignment was very clear and straightforward so I
knew directly what algorithm can be used.

[main.go](main.go) contains the algorithm and [helpers.go](helpers.go) data loading. I optimized the code for better
readability.

My results are following:

| Time Limit | Amount   |
|------------|----------|
| 50ms       | 4139.43  |
| 60ms       | 4675.71  |
| 90ms       | 6972.29  |
| 1000ms     | 35471.81 |

Side note: one should definitely not use floating point arithmetics when counting with money!
I did that only because the suggested `class Transaction` & the input defines `Amount` as float value...
Normally, one should use BigDecimal or similar and count with the smallest possible currency value in integers (so
pennies instead of dollars).


