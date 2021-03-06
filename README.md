# Amatriciana: the infinitely customizable chess engine

How do you make playing chess against a computer interesting? Chess engines beat the best human players with absolutely no effort, especially with recent advances in reinforcement learning technologies, but when they try to play worse on purpose it's just not the same as playing with a human. Computers playing weakly on purpose usually just play perfect moves and throw in gigantic blunders here and there. That is not how humans play at all. Human amateur players rarely just hang pieces, sometimes miss some tactics, and often lack long term strategical play. I think that it's possible to write a computers that plays like this.

My plan for this project is to write a chess engine that is reasonably strong, but can also be configured to:
1. Occasionally miss non-obvious tactics, without making one-move blunders.
2. Have a very customizable evaluation function, to simulate specific weaknesses in strategic play.