## Problem

- 5 philosophers sitting at a round table.
- 1 chopstick is placed between each adjacent pair.
- Want to eat rice from their plate, but needs two chopsticks.
- Only one philosopher can hold a chopstick at a time.
- Not enough chopsticks for everyone to eat at once.
 
**Implement the dining philosopher’s problem with the following constraints/modifications.**

- There should be 5 philosophers sharing chopsticks, with one chopstick between each adjacent pair of philosophers.
- Each philosopher should eat only 3 times (not in an infinite loop as we did in lecture).
- The philosophers pick up the chopsticks in any order, not lowest-numbered first (which we did in lecture).
- In order to eat, a philosopher must get permission from a host which executes in its own goroutine.
- The host allows no more than 2 philosophers to eat concurrently.
- Each philosopher is numbered, 1 through 5.
- When a philosopher starts eating (after it has obtained necessary locks) it prints `"starting to eat <number>"` on a line by itself, where `<number>` is the number of the philosopher.
- When a philosopher finishes eating (before it has released its locks) it prints `"finishing eating <number>"` on a line by itself, where `<number>` is the number of the philosopher.


