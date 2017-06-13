# Chip
--------

Chip is a set of utilities for basic audio / music synthesis for the Go language.
Included are a few (hacked-out) waveform generators, a basic rhythm assistant,
some audio transforms and a wav encoder.

# Terminology

Because music has words for just about everything (and some of them are used in 
several different senses) I've settled on an adhoc vocabulary for describing 
audio / musical entities:

* Tone: a sound of indeterminate length
* Scale: a collection of tones
* Note: a sound of specific duration
* Chord: a collection of notes
* Mod: a transformation (eg. fade, boost, filter)
* Sample: an individual pulse in PCM
