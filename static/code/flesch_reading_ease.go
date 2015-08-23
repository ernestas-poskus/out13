package main

import (
	"fmt"

	"github.com/ernestas-poskus/kincaid"
)

func main() {

	text := [...]string{
		// George Washington, First Inaugural Address (April 30, 1789)
		"In this conflict of emotions, all I dare aver, is, that it has been my faithful study to collect my duty from a just appreciation of every circumstance, by which it might be affected. All I dare hope, is, that, if in executing this task I have been too much swayed by a grateful remembrance of former instances, or by an affectionate sensibility to this transcendent proof, of the confidence of my fellow-citizens; and have thence too little consulted my incapacity as well as disinclination for the weighty and untried cares before me; my error will be palliated by the motives which misled me, and its consequences be judged by my Country, with some share of the partiality in which they originated.",

		// Abraham Lincoln, First Inaugural Address (March 4, 1861)
		"I take the official oath to-day with no mental reservations, and with no purpose to construe the Constitution or laws by any hypercritical rules. And while I do not choose now to specify particular acts of Congress as proper to be enforced, I do suggest that it will be much safer for all, both in official and private stations, to conform to and abide by all those acts which stand unrepealed, than to violate any of them, trusting to find impunity in having them held to be unconstitutional.",

		// Barack Obama , Acceptance Speech at the Democratic National Convention (August 28, 2008)
		"It is that promise that has always set this country apart—that through hard work and sacrifice, each of us can pursue our individual dreams but still come together as one American family, to ensure that the next generation can pursue their dreams as well. That's why I stand here tonight. Because for two hundred and thirty two years, at each moment when that promise was in jeopardy, ordinary men and women—students and soldiers, farmers and teachers, nurses and janitors—found the courage to keep it alive. We meet at one of those defining moments—a moment when our nation is at war, our economy is in turmoil, and the American promise has been threatened once more.",
	}

	for _, sentence := range text {
		score := kincaid.FleschReadingEase([]byte(sentence))
		fmt.Println(fmt.Sprintf("Score %.1f length: %d, text: %s", score, len(sentence), sentence[:50]))
	}
}
