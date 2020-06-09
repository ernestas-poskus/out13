+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper"]
author = "James Kirkpatrick, Razvan Pascanu, Neil Rabinowitz, Joel Veness, Guillaume Desjardins, Andrei A. Rusu, Kieran Milan, John Quan, Tiago Ramalho, Agnieszka Grabska-Barwinska, Demis Hassabis, Claudia Clopath b , Dharshan Kumaran, and Raia Hadsell"
name = "Overcoming catastrophic forgetting in neural networks"
link = "https://arxiv.org/abs/1612.00796"
title = "Overcoming catastrophic forgetting in neural networks"
date = 2017-11-12T23:05:18+02:00
+++

## Catastrophic forgetting in neural networks

The ability to learn tasks in a sequential fashion is crucial to the development of
artificial intelligence. Neural networks are not, in general, capable of this and it
has been widely thought that catastrophic forgetting is an inevitable feature of
connectionist models.

It is possible to overcome this limitation.

Our approach remembers old tasks by selectively slowing down
learning on the weights important for those tasks.

Achieving artificial general intelligence requires that agents are able to learn and remember many
different tasks.

A multitask learning paradigm—deep learning techniques have been used to
train single agents that can successfully play multiple Atari games.

The lack of algorithms to support continual learning thus remains a key
barrier to the development of artificial general intelligence.

### Elastic weight consolidation

A deep neural network consists of multiple layers of linear projection followed by element-wise
non-linearities. Learning a task consists of adjusting the set of weights and biases θ of the linear
projections, to optimize performance.

In order to justify this choice of constraint and to define which weights are most important for a task,
it is useful to consider neural network training from a probabilistic perspective.

#### Notes

> catastrophic forgetting: knowledge of previously learnt task(s) (e.g. task A) to be abruptly lost as information relevant to the current task (e.g. task B) is incorporated.

> elastic weight consolidation: this algorithm slows down learning on certain weights based on how important they are to previously seen tasks.
