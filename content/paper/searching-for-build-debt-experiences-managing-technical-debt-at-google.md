+++
description = "Take away from scientific papers, discussion & summary"
tags = ["paper"]
author = "J. David Morgenthaler, Misha Gridnev, Raluca Sauciuc, and Sanjay Bhansali"
name = "Technical Debt at Google"
link = "https://static.googleusercontent.com/media/research.google.com/en//pubs/archive/37755.pdf"
title = "Searching for Build Debt Experiences Managing Technical Debt at Google"
date = 2021-05-27T22:43:14+03:00
+++

## Technical debt at Google

Google engineers also make efforts to pay down that debt, whether through special Fixit
days, or via dedicated teams, variously known as janitors, cultivators, or demolition experts.

 - Automation. Use automated techniques to analyze and
(where possible) fix issues that contribute to our most
egregious technical debt. Several teams are working on
tools to make large scale changes easier. Our uniform
development style makes it feasible to do this at scale.
 - Make it easy to do the right thing. Many times technical
debt is incurred because people are not aware of it. If
we can analyze changes that developers are about to
make as part of their normal workflow (during editing,
browsing, or code review), we can prevent certain kinds
of debt. The size of the codebase and the rate at which
it changes makes this a non-trivial problem.
 - Make it hard to do the wrong thing. This is similar
to the above point, but with a greater emphasis on
introducing stricter checks on the kinds of actions
developers can do. For example, prevent people from
taking a dependency on code that is not ready for prime
time. Another example is to build stricter checks into
the compilation process and make them compile time
warnings or errors.
