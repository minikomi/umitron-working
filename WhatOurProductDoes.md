## Purpose of the product

Some of the problems fish farmers are facing are:

- Farmers in Japan are keeping track of everything on paper, those number on the papers don’t get digitalized, therefore the numbers don’t really get summarized and don’t provide much meaning.
- Since the fish lots get moved/split/merged a lot during their lifecycle, farmers don’t really know how much money they have spent on those fish until all of them are sold. They sell the fish at a price based on experience instead of evidence, and won’t know either they’ve made a profit until the end.

This is where we come in. We want to help the farmers input their numbers into a digital platform, calculate those numbers and provide them data in a meaningful way.

## User inputs

- Fish cages
    - Name, sizes, materials, etc
- Juveniles
    - Name, maker, initial fish size, price, count, purchased date, etc
- Fish lots
    - Name, juvenile, fish cage, start date, end date, etc
- Fish lot transactions (moves)
    - How many fish from which lot were moved to where and when
- Feeds
    - Name, type, maker, price
- Daily feeding weights
    - How much feed were given to which lot on what date
- Daily fish death
    - How many fish died in which lot on what date
- Daily shipments (sales)
    - How many fish were sold from which lot on what date
- Other stuff

## Some of our outputs

(For example, Juvenile J1 were first put into cage A1 then evenly split into cages A2 and A3, then later merged into cage C1)

- Current fish lots
    - Shows a list of ongoing fish lots
- Fish lot history
    - Shows the history of fish lot movements, i.e. J1 in C1 had a history of A1 → A2,A3, then A2,A3 → C1, etc.
- Calculate how many fish have died / were sold in each lot
    - How many fish died/sold during a specific period
    - How many fish died/sold since the beginning
- Calculate how much feed were given to each lot
    - How much feed were given to sold sold fish
    - How much feed are given to unsold fish
    - How much feed were spent on dead fish
    - How much feed were spent altogether
- Calculate how much money were spent on feed
    - For all of the above variations of feed amount, calculate how much money was spent based on the feed amount and feed price of the day.
- Calculate earnings and profits made after final sale
- Aggregate above numbers by juveniles
- Aggregate above numbers by fish species