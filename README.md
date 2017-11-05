# bowl
[![Build Status](https://travis-ci.org/rosenhouse/bowl.svg?branch=master)](https://travis-ci.org/rosenhouse/bowl)

A Go implementation of the [Bowling Kata++](https://docs.google.com/a/pivotal.io/document/d/16Zont3c1qD1hcO7mlMqcaRHUdYBKxGDm88qRO59kebg/edit?usp=sharing)

## Intended Usage

1. Generate a blank score file
  ```
  bowl new > my-game.txt
  ```


2. In your favorite text editor, mark up your game file


3. Score it
  ```
  bowl score < my-game.txt
  ```

## Working features
- Generating the blank score file
- Scoring simple games
- Games with spares

## Not yet implemented
- Scoring non-simple games
  - Games with strikes
  - Games with 3 throws in the final frame
