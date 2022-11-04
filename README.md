
# Fluidity Secret Snake Hiring Challenge

Thank you for applying for a frontend role with Fluidity Money!

In this task to test your Typescript React capabilities, you are expected
to create a webapp with functionality similar to the following over
a weekend:

![Example mock](mock.png)

You must use Typescript and React to implement your solution for this
task. You may use as many tools as you are comfortable with that don't
detract from an honest skills assessment.

*Please see [CHECKLIST.md](CHECKLIST.md) for how we score your approaches
to specific problems.*

*Please create a new repo to get started!*

**Please keep the repo private!**

## Background

You are being asked to build a mock Vickrey Auction for a collection of
snake-themed trading cards. These snake cards will appear and enter the
centre of the screen, with information on them being rendered (including
their stats and such.)

Bids will be delivered over websocket for the card, until its bidding
"slot" ends. During that bidding slot, another card can enter frame and
be bid on, and the previously seen cards will reenter the frame if their
bidding slot is chosen again. The bids made increase the "value locked"
of the snake card, which is presented on the screen.

For each of these cards, the bidding will close after a period known
ahead of time, and during that period, parties bidding are able to
"withdraw" their bids, and this will take place at random.

When a bid is chosen, the last bid is chosen as the "winner" and the
amount bid prior to that is chosen as the amount to be paid. This should
be displayed on screen, then hidden.

![Snake ends bidding period](bidding-finished.png)

### Stages (repeated)

So, there are three stages for a snake card.

The first being the stage of bidding.

At any point the snake card can enter vogue and be bid on. These bids
are entered randomly and any higher than the other highest bid are the
top highest bid. Other cards can enter the frame during this period if
the slot for this card ends.

The second is the expiration of the bidding for the snake. The expiration
means that only bids can be withdrawn, reducing the TVL.

Finally, at the end of this stage, the third stage is the selection of
the winner. The winner is the second highest bid for the snake.

### APIs

The following APIs are available:

#### /api/snakes

Get the list of currently available snakes and their respective stages.

Will return the following:

	[
		{
			"id": "snake-eater",
			"stage": 1,
			"bid": 0
		},
		{
			"id": "snake-foresight",
			"stage": 2,
			"bid": 0
		}
	]

Bid will always return 0 here and should be ignored.

#### /api/bids

Get the outstanding bids for a given snake. Takes the query param
"snake-id":

	curl -s http://localhost:8080/api/bids?snake-id=snake-eater

Would return:

	[ 192, 281, 28121, 201028, 271771 ]

#### /api/updates

A websocket update feed is provided to get updates on each snake in
their respective stages.

You could test the connection to the websocket using:

	wscat --connect http://localhost:8080/api/updates

Messages look like the following in practice:

![Websocket messages](wscat-messages.png)

Updates will look like the following:

	{
		"id": "snake-eater",
		"stage": 1,
		"bid": 18281
	}

The id field will identify the snake, with the stage being used to
indicate whether a transition to a new stage has taken place (if it
differs from the previous stage.)

New snakes will be identified by a new id and a bid of 0 at stage 1:

	{
		"id": "snake-gotham",
		"stage": 1,
		"bid": 0
	}

Bids will have different meanings depending on the current stage. Bids
in stage 1 are the addition of bids and increase the TVL.

Bids in stage 2 are exclusively the taking away of bids, reducing the
TVL.

For `snake-eater` with bids of: `[ 192, 281, 28121, 201028, 271771 ]`:


	{
		"id": "snake-eater",
		"stage": 2,
		"bid": 28121
	}

Would remove the bid `28121` and reduce the TVL by `28121`.

Snakes that enter stage 3 are in the stage of deletion. They should be
presented as having completed the bidding process and advertised in the
frontend the appropriate amount to be rewarded:

	{
		"id": "snake-eater",
		"stage": 3,
		"bid": 0
	}

(Bids are 0 in this stage since the bidding has completed for this snake.)

Good luck! Email Alex at [alex@fluidity.money](mailto:alex@fluidity.money)
if you have any issues or need any help.
