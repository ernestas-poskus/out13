+++
Description = "The Flesch/Flesch–Kincaid Readability Tests are readability tests designed to indicate how difficult a reading passage in English is to understand."
Tags = []
Categories = ["golang", "algorithms"]
date = "2015-04-04T15:13:11+03:00"
title = "Readability tests"

+++

[Flesch–Kincaid readability tests](http://en.wikipedia.org/wiki/Flesch%E2%80%93Kincaid_readability_tests).
Are readability tests designed to indicate how difficult a reading passage in English is to understand. There are two tests, the Flesch Reading Ease, and the Flesch–Kincaid Grade Level. [WIKI]

Due to own usage reasons I have decided to implement these readability tests in favourite language.
Algorithm is quite simple I will focus on Flesch Reading ease.

> Flesch Reading Ease test
>
> ![Flesh Reading Ease test](/images/reading_ease.png)

| Score | Notes |
| ----- | ----- |
| 90.0–100.0 | easily understood by an average 11-year-old student |
| 60.0–70.0 | easily understood by 13- to 15-year-old students |
| 0.0–30.0 | best understood by university graduates |

Implementation is trivial except ``total_syllables`` part. I have had number of struggles with this one.
Main caveat with syllables is that there no reliable facicility to reach consensus on exact syllable count.
Every source where I was verifying number of syllables had a difference usually +/- 1 in some cases +2 syllables per sentence.

Eventually decided to trust this: http://www.syllablecount.com/syllables/word.

Test cases here: [syllables_test.go](https://github.com/ernestas-poskus/syllables/blob/master/syllables_test.go)

tl;dr

# Flesch Reading Ease in action

Took presidential speeches from http://millercenter.org/president/speeches.

#### George Washington, First Inaugural Address (April 30, 1789)

> "In this conflict of emotions, all I dare aver, is, that it has been my faithful study to collect my duty from a just appreciation of every circumstance, by which it might be affected. All I dare hope, is, that, if in executing this task I have been too much swayed by a grateful remembrance of former instances, or by an affectionate sensibility to this transcendent proof, of the confidence of my fellow-citizens; and have thence too little consulted my incapacity as well as disinclination for the weighty and untried cares before me; my error will be palliated by the motives which misled me, and its consequences be judged by my Country, with some share of the partiality in which they originated.",

#### Abraham Lincoln, First Inaugural Address (March 4, 1861)

> "I take the official oath to-day with no mental reservations, and with no purpose to construe the Constitution or laws by any hypercritical rules. And while I do not choose now to specify particular acts of Congress as proper to be enforced, I do suggest that it will be much safer for all, both in official and private stations, to conform to and abide by all those acts which stand unrepealed, than to violate any of them, trusting to find impunity in having them held to be unconstitutional.",

#### Barack Obama, Acceptance Speech at the Democratic National Convention (August 28, 2008)

> "It is that promise that has always set this country apart—that through hard work and sacrifice, each of us can pursue our individual dreams but still come together as one American family, to ensure that the next generation can pursue their dreams as well. That's why I stand here tonight. Because for two hundred and thirty two years, at each moment when that promise was in jeopardy, ordinary men and women—students and soldiers, farmers and teachers, nurses and janitors—found the courage to keep it alive. We meet at one of those defining moments—a moment when our nation is at war, our economy is in turmoil, and the American promise has been threatened once more.",

# Results [code](/code/flesch_reading_ease.go)

| Author | year | length | Score(Flesch Reading ease) |
| ------ | ---- | ------ | -------------------------- |
| George Washington | 1789| 701 | 16.6 |
| Abraham Lincoln | 1861 | 493 | 35.3 |
| Barack Obama | 2008 | 676 | 50.3 |

It is interesting to see how complexity of presidential speeches changed as time passed.
